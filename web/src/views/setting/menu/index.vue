<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="CirclePlus" @click="addMenuDialog">新增</el-button>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" row-key="id">
          <el-table-column prop="id" label="ID" align="center" />
          <el-table-column prop="pid" label="父节点" align="center" />
          <el-table-column prop="meta.title" label="展示名称" align="center" />
          <el-table-column prop="name" label="路由名称" align="center" />
          <el-table-column prop="path" label="路由路径" width="130" align="center" />
          <el-table-column prop="meta.hidden" label="是否隐藏" align="center">
            <template #default="scope">
              <el-tag v-if="!scope.row.meta.hidden" type="success" effect="plain">显示</el-tag>
              <el-tag v-else type="warning" effect="plain">隐藏</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="component" label="组件路径" min-width="180" align="center" />
          <el-table-column fixed="right" label="操作" width="180" align="center">
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" :before-close="handleClose">
      <warning-bar title="新增菜单，需要在角色管理内配置权限才可使用" />
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        :inline="true"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="父节点" prop="pid" style="width: 30%" required>
          <el-cascader
            v-model="formData.pid"
            style="width: 100%"
            :options="menuOption"
            :props="{ checkStrictly: true, emitPath: false }"
            :show-all-levels="false"
            clearable
          />
        </el-form-item>
        <el-form-item label="路由名称" prop="name" style="width: 30%">
          <el-input v-model="formData.name" />
        </el-form-item>
        <el-form-item label="路由路径" prop="path" required style="width: 30%">
          <el-input v-model="formData.path" />
        </el-form-item>
        <el-form-item label="前端组件" prop="component" style="width: 30%">
          <el-input v-model="formData.component" />
        </el-form-item>
        <el-form-item label="重定向" prop="redirect" style="width: 30%">
          <el-input v-model="formData.redirect" />
        </el-form-item>
        <el-form-item label="展示名称" prop="meta.title" style="width: 30%">
          <el-input v-model="formData.meta.title" />
        </el-form-item>
        <el-form-item label="是否隐藏" prop="meta.hidden" style="width: 30%">
          <el-select v-model="formData.meta.hidden">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="meta.icon" style="width: 30%">
          <icon :meta="formData.meta" style="width: 100%" />
        </el-form-item>
        <el-form-item label="是否固定" prop="meta.affix" style="width: 30%">
          <el-select v-model="formData.meta.affix">
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
import { type MenusData, type reqMenu, getMenus, addMenuApi, editMenuApi, deleteMenuApi } from "@/api/system/menu"
import WarningBar from "@/components/warningBar/warningBar.vue"
import icon from "./icon.vue"

const loading = ref<boolean>(false)
const dialogVisible = ref<boolean>(false)

const tableData = ref<MenusData[]>([])
const getTableData = async () => {
  loading.value = false
  const res = await getMenus()
  if (res.code === 0) {
    tableData.value = res.data
  }
  loading.value = true
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

let kind: string

const addMenuDialog = () => {
  dialogTitle.value = "新增菜单"
  setOptions()
  kind = "Add"
  dialogVisible.value = true
}

const editMenuDialog = (row: MenusData) => {
  dialogTitle.value = "编辑菜单"
  setOptions()
  kind = "Edit"
  formData.id = row.id
  formData.pid = row.pid
  if (row.name) {
    formData.name = row.name
  }
  formData.path = row.path
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
  dialogVisible.value = true
}

const deleteMenuAction = (row: MenusData) => {
  ElMessageBox.confirm("此操作将永久删除所有角色下该菜单, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(() => {
    const index = tableData.value.indexOf(row)
    deleteMenuApi({ id: row.id }).then((res) => {
      if (res.code === 0) {
        ElMessage({ type: "success", message: res.msg })
        tableData.value.splice(index, 1)
      }
    })
  })
}

// 表单
const formRef = ref<FormInstance>()

const formRules: FormRules = reactive({
  pid: [{ required: true, trigger: "change", message: "请选择父节点" }],
  path: [{ required: true, trigger: "blur", message: "请填写路由路径" }]
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
}

const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}

const handleClose = (done: Function) => {
  initForm()
  done()
}

const formData = reactive<reqMenu>({
  id: 0,
  name: "",
  path: "",
  component: "",
  redirect: "",
  pid: 0,
  meta: {
    title: "",
    icon: "",
    hidden: false,
    affix: false
  }
})

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      if (kind === "Add") {
        const res = await addMenuApi(formData)
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      } else if (kind === "Edit") {
        const res = await editMenuApi(formData)
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      }
      initForm()
      dialogVisible.value = false
    }
  })
}
</script>

<style lang="scss" scoped>
.toolbar-wrapper {
  margin-bottom: 20px;
}

.table-wrapper {
  margin-bottom: 20px;
}
</style>
