<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"
import type { roleDataModel } from "@/api/sysManagement/role"
import { usePagination } from "@@/composables/usePagination_n"
import { reactive, ref } from "vue"
import { 
  roleCreateApi, 
  roleDeleteApi, 
  roleListApi, 
  roleTreeApi,
  roleUpdateApi,
  setRoleInheritanceApi 
} from "@/api/sysManagement/role"
import Apis from "./components/apis.vue"
import Menus from "./components/menus.vue"

defineOptions({
  name: "Role"
})

const loading = ref<boolean>(false)
const { paginationData, changeCurrentPage, changePageSize } = usePagination()
const tableData = ref<roleDataModel[]>([])
let activeRow: roleDataModel

// 分页
function handleSizeChange(value: number) {
  changePageSize(value)
  getTableData()
}

function handleCurrentChange(value: number) {
  changeCurrentPage(value)
  getTableData()
}

async function getTableData() {
  loading.value = true
  const res = await roleListApi({ page: paginationData.currentPage, pageSize: paginationData.pageSize })
  if (res.code === 0) {
    tableData.value = res.data.list
    paginationData.total = res.data.total
  }
  loading.value = false
}
getTableData()

function initForm() {
  formData.roleName = ""
  formData.parentId = undefined
}

const dialogVisible = ref<boolean>(false)
function handleClose(done: () => void) {
  initForm()
  done()
}

const formRef = ref<FormInstance>()
const formData = reactive({
  id: 0,
  roleName: "",
  parentId: undefined as number | undefined  // 父角色ID
})
const formRules: FormRules = reactive({
  roleName: [{ required: true, trigger: "blur", message: "请填写角色名称" }]
})

const kind = ref("")
const title = ref("")
function addDialog() {
  kind.value = "Add"
  title.value = "新增角色"
  dialogVisible.value = true
}

function editDialog(row: roleDataModel) {
  kind.value = "Edit"
  title.value = "编辑角色"
  activeRow = row
  formData.roleName = row.roleName
  formData.parentId = row.parentId
  dialogVisible.value = true
}

function closeDialog() {
  initForm()
  dialogVisible.value = false
}

function operateAction(formEl: FormInstance | undefined) {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      if (kind.value === "Add") {
        const res = await roleCreateApi({ 
          roleName: formData.roleName,
          parentId: formData.parentId 
        })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          tableData.value.push(res.data)
        }
      } else if (kind.value === "Edit") {
        const res = await roleUpdateApi({ 
          id: activeRow.id, 
          roleName: formData.roleName,
          parentId: formData.parentId 
        })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          const index = tableData.value.indexOf(activeRow)
          tableData.value[index].roleName = formData.roleName
          tableData.value[index].parentId = formData.parentId
        }
      }
      closeDialog()
    }
  })
}

function deleteRoleAction(row: roleDataModel) {
  ElMessageBox.confirm("此操作将永久删除该角色, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      const index = tableData.value.indexOf(row)
      roleDeleteApi({ id: row.id }).then((res) => {
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          tableData.value.splice(index, 1)
        }
      })
    })
    .catch(() => {})
}

// 角色设置
const drawer = ref(false)
let activeId: number
function openDrawer(row: roleDataModel) {
  activeId = row.id
  drawer.value = true
}

// 角色继承设置
const inheritDialogVisible = ref(false)
const inheritForm = reactive({
  childRoleId: 0,
  parentRoleId: undefined as number | undefined
})

function openInheritDialog(row: roleDataModel) {
  inheritForm.childRoleId = row.id
  inheritForm.parentRoleId = row.parentId
  inheritDialogVisible.value = true
}

async function saveInheritance() {
  if (!inheritForm.parentRoleId) {
    ElMessage.warning("请选择父角色")
    return
  }
  const res = await setRoleInheritanceApi({
    childRoleId: inheritForm.childRoleId,
    parentRoleId: inheritForm.parentRoleId
  })
  if (res.code === 0) {
    ElMessage.success("设置成功")
    inheritDialogVisible.value = false
    getTableData()
  }
}

// 获取父角色名称
function getParentName(parentId?: number) {
  if (!parentId) return "-"
  const parent = tableData.value.find(r => r.id === parentId)
  return parent ? parent.roleName : "-"
}
</script>

<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="addDialog">
            新增
          </el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" row-key="id" default-expand-all>
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="roleName" label="角色名称" />
          <el-table-column label="父角色" width="150">
            <template #default="scope">
              <el-tag v-if="scope.row.parentId" type="info">
                {{ getParentName(scope.row.parentId) }}
              </el-tag>
              <span v-else>-</span>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" align="center" width="350">
            <template #default="scope">
              <el-button type="primary" text icon="Setting" size="small" @click="openDrawer(scope.row)">
                设置权限
              </el-button>
              <el-button type="primary" text icon="Link" size="small" @click="openInheritDialog(scope.row)">
                继承设置
              </el-button>
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">
                编辑
              </el-button>
              <el-button
                type="danger"
                text
                icon="Delete"
                size="small"
                @click="deleteRoleAction(scope.row)"
                :disabled="scope.row.roleName === 'root'"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pager-wrapper">
        <el-pagination
          background
          :layout="paginationData.layout"
          :page-sizes="paginationData.pageSizes"
          :total="paginationData.total"
          :page-size="paginationData.pageSize"
          :current-page="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="title" :before-close="handleClose" width="35%">
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
        label-position="left"
        style="width: 95%; margin-top: 15px"
      >
        <el-form-item label="角色名称" prop="roleName">
          <el-input v-model="formData.roleName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="父角色">
          <el-select v-model="formData.parentId" clearable placeholder="请选择父角色（可选）" style="width: 100%">
            <el-option 
              v-for="role in tableData.filter(r => r.id !== formData.id)" 
              :key="role.id" 
              :label="role.roleName" 
              :value="role.id" 
            />
          </el-select>
          <div class="el-form-item__tip">继承父角色的所有权限</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">
            取消
          </el-button>
          <el-button type="primary" @click="operateAction(formRef)">
            确认
          </el-button>
        </div>
      </template>
    </el-dialog>
    <el-drawer v-if="drawer" v-model="drawer" :with-header="false" size="35%" title="角色配置">
      <el-tabs type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus :id="activeId" />
        </el-tab-pane>
        <el-tab-pane label="角色接口">
          <Apis :id="activeId" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
    <el-dialog v-model="inheritDialogVisible" title="设置角色继承" width="30%">
      <el-form label-width="100px">
        <el-form-item label="父角色">
          <el-select v-model="inheritForm.parentRoleId" clearable placeholder="请选择父角色" style="width: 100%">
            <el-option 
              v-for="role in tableData.filter(r => r.id !== inheritForm.childRoleId)" 
              :key="role.id" 
              :label="role.roleName" 
              :value="role.id" 
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="inheritDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveInheritance">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
