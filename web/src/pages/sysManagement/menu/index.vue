<script lang="ts" setup>
import type { CascaderOption, FormInstance, FormRules } from "element-plus"
import type { IconName } from "./icon.vue"
import type { MenuDataModel } from "@/api/sysManagement/menu"
import WarningBar from "@@/components/WarningBar/warningBar.vue"
import { reactive, ref } from "vue"
import { menuCreateApi, menuDeleteApi, menuListApi, menuUpdateApi } from "@/api/sysManagement/menu"
import icon from "./icon.vue"

defineOptions({
  name: "SMenu"
})

const loading = ref<boolean>(false)
const dialogVisible = ref<boolean>(false)

const tableData = ref<MenuDataModel[]>([])
async function getTableData() {
  loading.value = true
  const res = await menuListApi()
  if (res.code === 0) {
    tableData.value = res.data
    setOptions()
  }
  loading.value = false
}
getTableData()

const dialogTitle = ref<string>("")

const menuOption: CascaderOption[] = []

function setOptions() {
  menuOption.length = 0
  menuOption.push({
    value: 0,
    label: "根目录"
  })
  setMenuOptions(tableData.value, menuOption)
}

function setMenuOptions(menuData: MenuDataModel[], optionsData: CascaderOption[]) {
  for (const item of menuData) {
    if (item.menu_name === "ErrorPage") {
      continue
    }
    if (item.children && item.children.length) {
      const option = {
        label: item.title,
        value: item.id,
        children: []
      }
      setMenuOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        label: item.title,
        value: item.id
      }
      optionsData.push(option)
    }
  }
}

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}
const oKind = ref<operationKind>()

function addMenuDialog() {
  dialogTitle.value = "新增菜单"
  oKind.value = operationKind.Add
  dialogVisible.value = true
}

let activeRowId: number
function editMenuDialog(row: MenuDataModel) {
  dialogTitle.value = "编辑菜单"
  oKind.value = operationKind.Edit
  activeRowId = row.id
  formData.parentId = row.parentId
  if (row.menu_name) {
    formData.menu_name = row.menu_name
  }
  formData.path = row.path
  formData.sort = row.sort
  if (row.component) {
    formData.component = row.component
  }
  if (row.redirect) {
    formData.redirect = row.redirect
  }
  if (row.title) {
    formData.title = row.title
  }
  if (row.icon) {
    formData.icon = row.icon
  }
  formData.hidden = !!row.hidden
  formData.affix = Boolean(row.affix)
  formData.keepAlive = Boolean(row.keepAlive)
  formData.alwaysShow = Boolean(row.alwaysShow)
  dialogVisible.value = true
}

function deleteMenuAction(row: MenuDataModel) {
  ElMessageBox.confirm("此操作将永久删除所有角色下该菜单, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      menuDeleteApi({ id: row.id }).then((res) => {
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      })
    })
    .catch(() => { })
}

// 表单
const formRef = ref<FormInstance>()

const formRules: FormRules = reactive({
  parentId: [{ required: true, trigger: "blur", message: "请选择父节点" }],
  path: [{ required: true, trigger: "blur", message: "请填写路由路径" }],
  component: [{ required: true, trigger: "blur", message: "请填写前端组件路径" }]
})

function initForm() {
  formData.menu_name = ""
  formData.path = ""
  formData.component = ""
  formData.redirect = ""
  formData.parentId = 0
  formData.sort = 0
  formData.title = ""
  formData.icon = ""
  formData.hidden = false
  formData.affix = false
  formData.keepAlive = false
  formData.alwaysShow = true
}

function closeDialog() {
  dialogVisible.value = false
  initForm()
}

function handleClose(done: () => void) {
  initForm()
  done()
}

const formData = reactive<{
  menu_name: string
  path: string
  component: string
  redirect: string
  parentId: number
  sort: number
  title: string
  icon: IconName | ""
  hidden: boolean
  affix: boolean
  keepAlive: boolean
  alwaysShow: boolean
}>({
  menu_name: "",
  path: "",
  component: "",
  redirect: "",
  parentId: 0,
  sort: 0,
  title: "",
  icon: "",
  hidden: false,
  affix: false,
  keepAlive: false,
  alwaysShow: true
})

function operateAction(formEl: FormInstance | undefined) {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      const tempMenu = {
        parentId: formData.parentId,
        menu_name: formData.menu_name,
        path: formData.path,
        component: formData.component,
        sort: formData.sort,
        redirect: formData.redirect,
        title: formData.title,
        icon: formData.icon,
        hidden: formData.hidden,
        affix: formData.affix,
        keepAlive: formData.keepAlive,
        alwaysShow: formData.alwaysShow
      }
      if (oKind.value === "Add") {
        const res = await menuCreateApi(tempMenu)
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
        }
      } else if (oKind.value === "Edit") {
        const res = await menuUpdateApi({ id: activeRowId, ...tempMenu })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      }
      dialogVisible.value = false
      initForm()
    }
  })
}
</script>

<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="addMenuDialog">
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
        <el-table :data="tableData" row-key="id">
          <el-table-column prop="id" label="ID" />
          <el-table-column prop="parentId" label="父节点" />
          <el-table-column prop="title" label="展示名称">
            <template #default="scope">
              <el-tag :effect="scope.row.parentId === 0 ? 'light' : 'plain'">
                {{ scope.row.title }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="路由名称" />
          <el-table-column prop="path" label="路由路径" />
          <el-table-column prop="hidden" label="是否隐藏">
            <template #default="scope">
              <el-tag v-if="!scope.row.hidden" type="success" effect="plain">
                显示
              </el-tag>
              <el-tag v-else type="warning" effect="plain">
                隐藏
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="排序" />
          <el-table-column prop="component" label="组件路径" min-width="180" />
          <el-table-column fixed="right" label="操作" align="center" min-width="180">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editMenuDialog(scope.row)">
                编辑
              </el-button>
              <el-button type="danger" text icon="Delete" size="small" @click="deleteMenuAction(scope.row)">
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <el-dialog
      v-model="dialogVisible" :title="dialogTitle" :before-close="handleClose" :close-on-click-modal="false"
      :draggable="true"
    >
      <WarningBar title="新增菜单，需要在角色管理内配置权限才可使用" />
      <el-form
        ref="formRef" :model="formData" :rules="formRules" :inline="true" label-position="top"
        label-width="85px"
      >
        <el-form-item label="父节点" prop="parentId" style="width: 30%">
          <el-cascader
            v-model="formData.parentId" style="width: 100%" :options="menuOption"
            :props="{ checkStrictly: true, emitPath: false }" clearable filterable
          />
        </el-form-item>
        <el-form-item label="路由名称" prop="name" style="width: 30%">
          <el-input v-model="formData.menu_name" />
        </el-form-item>
        <el-form-item label="路由路径" prop="path" style="width: 30%">
          <el-input v-model="formData.path" />
        </el-form-item>
        <el-form-item label="前端组件路径" prop="component" style="width: 30%">
          <el-input v-model="formData.component" />
        </el-form-item>
        <el-form-item label="重定向" prop="redirect" style="width: 30%">
          <el-input v-model="formData.redirect" />
        </el-form-item>
        <el-form-item label="展示名称" prop="meta.title" style="width: 30%">
          <el-input v-model="formData.title" />
        </el-form-item>
        <el-form-item label="排序" prop="sort" style="width: 30%">
          <el-input-number v-model="formData.sort" :min="1" />
        </el-form-item>
        <el-form-item label="隐藏" prop="hidden" style="width: 30%">
          <el-select v-model="formData.hidden">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="icon" style="width: 30%">
          <icon v-model="formData.icon" style="width: 100%" />
        </el-form-item>
        <el-form-item label="固定" prop="affix" style="width: 30%">
          <el-select v-model="formData.affix">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="一直显示根路由" prop="alwaysShow" style="width: 30%">
          <el-select v-model="formData.alwaysShow">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="KeepAlive" prop="keepAlive" style="width: 30%">
          <el-select v-model="formData.keepAlive">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
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
  </div>
</template>
