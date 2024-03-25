<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="addDialog">新增</el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column prop="id" label="ID" />
          <el-table-column prop="roleName" label="名称" />
          <el-table-column fixed="right" label="操作" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Setting" size="small" @click="openDrawer(scope.row)"
                >设置权限</el-button
              >
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
              <el-button
                type="danger"
                text
                icon="Delete"
                size="small"
                @click="deleteRoleAction(scope.row)"
                :disabled="scope.row.roleName === 'root'"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="title" :before-close="handleClose" width="30%">
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
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="operateAction(formRef)">确认</el-button>
        </div>
      </template>
    </el-dialog>
    <el-drawer v-if="drawer" v-model="drawer" :with-header="false" size="35%" title="角色配置">
      <el-tabs type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :id="activeId" />
        </el-tab-pane>
        <el-tab-pane label="角色接口">
          <Apis ref="apis" :id="activeId" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox } from "element-plus"
import { type roleDataModel, getRolesApi, addRoleApi, deleteRoleApi, editRoleApi } from "@/api/authority/role"
import Menus from "./components/menus.vue"
import Apis from "./components/apis.vue"

defineOptions({
  name: "Role"
})

const loading = ref<boolean>(false)
const tableData = ref<roleDataModel[]>([])
let activeRow: roleDataModel

const getTableData = async () => {
  loading.value = true
  const res = await getRolesApi()
  if (res.code === 0) {
    tableData.value = res.data
  }
  loading.value = false
}
getTableData()

const initForm = () => {
  formData.roleName = ""
}

const dialogVisible = ref<boolean>(false)
const handleClose = (done: Function) => {
  initForm()
  done()
}

const formRef = ref<FormInstance>()
const formData = reactive({
  id: 0,
  roleName: ""
})
const formRules: FormRules = reactive({
  roleName: [{ required: true, trigger: "blur", message: "请填写角色名称" }]
})

const kind = ref("")
const title = ref("")
const addDialog = () => {
  kind.value = "Add"
  title.value = "新增角色"
  dialogVisible.value = true
}

const editDialog = (row: roleDataModel) => {
  kind.value = "Edit"
  title.value = "编辑角色"
  activeRow = row
  formData.roleName = row.roleName
  dialogVisible.value = true
}

const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      if (kind.value === "Add") {
        const res = await addRoleApi({ roleName: formData.roleName })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          // const tempData = {
          //   id: res.data.id,
          //   roleName: res.data.roleName,
          //   menus: []
          // }
          tableData.value.push(res.data)
        }
      } else if (kind.value === "Edit") {
        const res = await editRoleApi({ id: activeRow.id, roleName: formData.roleName })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          const index = tableData.value.indexOf(activeRow)
          tableData.value[index].roleName = formData.roleName
        }
      }
      closeDialog()
    }
  })
}

const deleteRoleAction = (row: roleDataModel) => {
  ElMessageBox.confirm("此操作将永久删除该角色, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      const index = tableData.value.indexOf(row)
      deleteRoleApi({ id: row.id }).then((res) => {
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
const openDrawer = (row: roleDataModel) => {
  activeId = row.id
  drawer.value = true
}
</script>
