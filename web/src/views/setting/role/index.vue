<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="CirclePlus" @click="addDialog">新增</el-button>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column prop="ID" label="ID" />
          <el-table-column prop="roleName" label="名称" align="center" />
          <el-table-column fixed="right" label="操作" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Setting" size="small">设置权限</el-button>
              <el-button type="primary" text icon="Edit" size="small">编辑</el-button>
              <el-button type="danger" text icon="Delete" size="small" :disabled="scope.row.roleName === 'root'"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <el-dialog v-model="dialogVisible" title="新增角色" :before-close="handleClose" width="30%">
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        style="display: flex; justify-content: center; align-items: center; margin-top: 20px"
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
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import { type FormInstance, type FormRules, ElMessage } from "element-plus"
import { type roleData, type reqRole, getRoles, addRole } from "@/api/system/role"

const loading = ref<boolean>(false)
const tableData = ref<roleData[]>([])

const getTableData = async () => {
  loading.value = true
  const res = await getRoles()
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
  roleName: ""
})
const formRules: FormRules = reactive({
  roleName: [{ required: true, trigger: "blur", message: "请填写角色名称" }]
})
const addDialog = () => {
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
      const tempRole: reqRole = {
        roleName: formData.roleName
      }
      const res = await addRole(tempRole)
      if (res.code === 0) {
        ElMessage({ type: "success", message: res.msg })
        const tempData: roleData = {
          ID: res.data.ID,
          roleName: res.data.roleName,
          menus: []
        }
        tableData.value.push(tempData)
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
