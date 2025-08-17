<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div><el-button type="primary" icon="CirclePlus" @click="addMenuDialog">新增</el-button></div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" row-key="id">
          <el-table-column prop="id" label="ID" />
          <el-table-column prop="pid" label="父节点" />
          <el-table-column prop="meta.title" label="展示名称">
            <template #default="scope">
              <el-tag :effect="scope.row.pid === 0 ? 'light' : 'plain'">{{ scope.row.meta.title }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="路由名称" />
          <el-table-column prop="path" label="路由路径" />
          <el-table-column prop="meta.hidden" label="是否隐藏">
            <template #default="scope">
              <el-tag v-if="!scope.row.meta.hidden" type="success" effect="plain">显示</el-tag>
              <el-tag v-else type="warning" effect="plain">隐藏</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="排序" />
          <el-table-column prop="component" label="组件路径" min-width="180" />
          <el-table-column fixed="right" label="操作" align="center" min-width="180">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editMenuDialog(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" text icon="Delete" size="small" @click="deleteMenuAction(scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      :before-close="handleClose"
      :close-on-click-modal="false"
      :draggable="true"
    >
      <warning-bar title="新增菜单，需要在角色管理内配置权限才可使用" />
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        :inline="true"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="父节点" prop="pid" style="width: 30%">
          <el-cascader
            v-model="formData.pid"
            style="width: 100%"
            :options="menuOption"
            :props="{ checkStrictly: true, emitPath: false }"
            clearable
            filterable
          />
        </el-form-item>
        <el-form-item label="路由名称" prop="name" style="width: 30%">
          <el-input v-model="formData.name" />
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
          <el-input v-model="formData.meta.title" />
        </el-form-item>
        <el-form-item label="排序" prop="sort" style="width: 30%">
          <el-input-number v-model="formData.sort" :min="1" />
        </el-form-item>
        <el-form-item label="隐藏" prop="meta.hidden" style="width: 30%">
          <el-select v-model="formData.meta.hidden">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="meta.icon" style="width: 30%">
          <icon :meta="formData.meta" style="width: 100%" />
        </el-form-item>
        <el-form-item label="固定" prop="meta.affix" style="width: 30%">
          <el-select v-model="formData.meta.affix">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="一直显示根路由" prop="meta.alwaysShow" style="width: 30%">
          <el-select v-model="formData.meta.alwaysShow">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="KeepAlive" prop="meta.keepAlive" style="width: 30%">
          <el-select v-model="formData.meta.keepAlive">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="operateAction(formRef)">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import { ElMessage, ElMessageBox, type FormInstance, type FormRules, type CascaderOption } from "element-plus"
import { type MenuDataModel, getMenus, addMenuApi, editMenuApi, deleteMenuApi } from "@/api/authority/menu"
import WarningBar from "@/components/WarningBar/warningBar.vue"
import icon from "./icon.vue"

defineOptions({
  name: "SMenu"
})

const loading = ref<boolean>(false)
const dialogVisible = ref<boolean>(false)

const tableData = ref<MenuDataModel[]>([])
const getTableData = async () => {
  loading.value = true
  const res = await getMenus()
  if (res.code === 0) {
    tableData.value = res.data
  }
  loading.value = false
}
getTableData()

const dialogTitle = ref<string>("")

const menuOption: CascaderOption[] = []

const setOptions = () => {
  menuOption.length = 0
  menuOption.push({
    value: 0,
    label: "根目录"
  })
  setMenuOptions(tableData.value, menuOption)
}

const setMenuOptions = (menuData: any, optionsData: CascaderOption[]) => {
  for (const item of menuData) {
    if (item.name === "ErrorPage") {
      continue
    }
    if (item.children && item.children.length) {
      const option = {
        label: item.meta.title,
        value: item.id,
        children: []
      }
      setMenuOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        label: item.meta.title,
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
let oKind: operationKind

const addMenuDialog = () => {
  dialogTitle.value = "新增菜单"
  setOptions()
  oKind = operationKind.Add
  dialogVisible.value = true
}

let activeRowId: number
const editMenuDialog = (row: MenuDataModel) => {
  dialogTitle.value = "编辑菜单"
  setOptions()
  oKind = operationKind.Edit
  activeRowId = row.id
  formData.pid = row.pid
  if (row.name) {
    formData.name = row.name
  }
  formData.path = row.path
  formData.sort = row.sort
  if (row.component) {
    formData.component = row.component
  }
  if (row.redirect) {
    formData.redirect = row.redirect
  }
  if (row.meta?.title) {
    formData.meta.title = row.meta?.title
  }
  if (row.meta?.svgIcon) {
    formData.meta.icon = row.meta?.svgIcon
  }
  formData.meta.hidden = !!row.meta?.hidden
  formData.meta.affix = Boolean(row.meta?.affix)
  formData.meta.keepAlive = Boolean(row.meta?.keepAlive)
  formData.meta.alwaysShow = Boolean(row.meta?.alwaysShow)
  dialogVisible.value = true
}

const deleteMenuAction = (row: MenuDataModel) => {
  ElMessageBox.confirm("此操作将永久删除所有角色下该菜单, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteMenuApi({ id: row.id }).then((res) => {
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      })
    })
    .catch(() => {})
}

// 表单
const formRef = ref<FormInstance>()

const formRules: FormRules = reactive({
  pid: [{ required: true, trigger: "blur", message: "请选择父节点" }],
  path: [{ required: true, trigger: "blur", message: "请填写路由路径" }],
  component: [{ required: true, trigger: "blur", message: "请填写前端组件路径" }]
})

const initForm = () => {
  formData.pid = 0
  formData.name = ""
  formData.path = ""
  formData.component = ""
  formData.redirect = ""
  formData.meta.title = ""
  formData.meta.icon = ""
  formData.meta.hidden = false
  formData.meta.affix = false
  formData.meta.keepAlive = false
  formData.meta.alwaysShow = false
}

const closeDialog = () => {
  dialogVisible.value = false
  initForm()
}

const handleClose = (done: Function) => {
  initForm()
  done()
}

const formData = reactive({
  name: "",
  path: "",
  component: "",
  redirect: "",
  pid: 0,
  sort: 0,
  meta: {
    title: "",
    icon: "",
    hidden: false,
    affix: false,
    keepAlive: false,
    alwaysShow: true
  }
})

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      const tempMenu = {
        pid: formData.pid,
        name: formData.name,
        path: formData.path,
        component: formData.component,
        sort: formData.sort,
        redirect: formData.redirect || undefined,
        meta: {
          title: formData.meta.title,
          icon: formData.meta.icon || undefined,
          hidden: formData.meta.hidden || undefined,
          affix: formData.meta.affix || undefined,
          keepAlive: formData.meta.keepAlive || undefined,
          alwaysShow: formData.meta.alwaysShow || undefined
        }
      }
      if (oKind === "Add") {
        const res = await addMenuApi(tempMenu)
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
        }
      } else if (oKind === "Edit") {
        const res = await editMenuApi({ id: activeRowId, ...tempMenu })
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
