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
          <el-table-column prop="meta.title" label="是否隐藏" align="center">
            <template #default="scope">
              <el-tag v-if="scope.row.meta.title" type="success" effect="plain">显示</el-tag>
              <el-tag v-else type="warning" effect="plain">隐藏</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="component" label="组件路径" min-width="180" align="center" />
          <el-table-column fixed="right" label="操作" width="180" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editMenuAction()">编辑</el-button>
              <el-button type="danger" text icon="Delete" size="small" @click="deleteMenuAction()">删除</el-button>
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
        <el-form-item label="展示名称" prop="title" style="width: 30%">
          <el-input v-model="formData.title" />
        </el-form-item>
        <el-form-item label="是否隐藏" prop="hidden" style="width: 30%">
          <el-select v-model="formData.hidden">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="icon" style="width: 30%">
          <el-select v-model="formData.icon" />
        </el-form-item>
        <el-form-item label="是否固定" prop="affix" style="width: 30%">
          <el-select v-model="formData.affix">
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
import { ElMessage, type FormInstance, type FormRules, type CascaderOption } from "element-plus"
import { usePermissionStoreHook } from "@/store/modules/permission"
import { type MenusData, type addMenuData, addMenuApi } from "@/api/system/menu"
import WarningBar from "@/components/warningBar/warningBar.vue"

const loading = ref<boolean>(false)
const dialogVisible = ref<boolean>(false)
const tableData = ref<MenusData[]>([])
const permissionStore = usePermissionStoreHook()
tableData.value = permissionStore.asyncRouterList

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
    if (item.name === "ErrorPage" || item.path === "/") {
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

const addMenuDialog = () => {
  dialogTitle.value = "新增菜单"
  setOptions()
  dialogVisible.value = true
}

const editMenuAction = () => {}
const deleteMenuAction = () => {}

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
  formData.title = ""
  formData.icon = ""
  formData.hidden = false
  formData.affix = false
}

const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}

const handleClose = (done: Function) => {
  initForm()
  done()
}

const formData = reactive<addMenuData>({
  name: "",
  path: "",
  component: "",
  redirect: "",
  pid: 0,
  title: "",
  icon: "",
  hidden: false,
  affix: false
})

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      const res = await addMenuApi(formData)
      if (res.code === 0) {
        ElMessage({ type: "success", message: res.msg })
      }
      initForm()
      dialogVisible.value = false
    }
  })
}
</script>

<style lang="scss">
.toolbar-wrapper {
  margin-bottom: 20px;
}

.table-wrapper {
  margin-bottom: 20px;
}
</style>
