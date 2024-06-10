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
          <el-table-column prop="username" label="用户名" />
          <el-table-column prop="phone" label="手机号" />
          <el-table-column prop="email" label="邮箱" />
          <el-table-column prop="roleName" label="角色">
            <template #default="scope">
              <el-tag type="success">{{ scope.row.roleName }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="active" label="状态">
            <template #default="scope">
              <el-switch
                v-model="scope.row.active"
                inline-prompt
                :active-value="true"
                :inactive-value="false"
                active-text="启用"
                inactive-text="禁用"
                @change="switchAction(scope.row.id, scope.row.active)"
                :disabled="scope.row.username === 'admin' && scope.row.role === 'root'"
              />
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" align="center" min-width="200px">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
              <el-button type="primary" text icon="Key" size="small" @click="modifyDialog(scope.row)"
                >修改密码</el-button
              >
              <el-button
                type="danger"
                text
                icon="Delete"
                size="small"
                @click="deleteUserAction(scope.row)"
                :disabled="scope.row.username === 'admin'"
                >删除</el-button
              >
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
          :currentPage="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
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
        <el-form-item label="用户名" prop="username">
          <el-input v-model="formData.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-if="kind === 'Add'">
          <el-input v-model="formData.password" autocomplete="off" type="password" show-password />
        </el-form-item>
        <el-form-item label="手机号码" prop="phone">
          <el-input v-model="formData.phone" autocomplete="off" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态" prop="active">
          <el-switch v-model="formData.active" active-text="启用" inactive-text="禁用" />
        </el-form-item>
        <el-form-item label="角色" prop="roleId" required>
          <el-select v-model="formData.roleId">
            <el-option v-for="role in roleOptions" :key="role.ID" :label="role.roleName" :value="role.ID" />
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
    <el-dialog v-model="mpDialogVisible" title="修改密码" :before-close="mpHandleClose" width="25%">
      <el-form
        ref="mpFormRef"
        :model="mpFormData"
        :rules="mpFormRules"
        label-width="100px"
        label-position="left"
        style="width: 95%; margin-top: 15px"
      >
        <el-form-item label="旧密码" prop="oldPassword">
          <el-input v-model="mpFormData.oldPassword" autocomplete="off" type="password" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="mpFormData.newPassword" autocomplete="off" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="rePassword">
          <el-input v-model="mpFormData.rePassword" autocomplete="off" type="password" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="mpCloseDialog">取消</el-button>
          <el-button type="primary" @click="mpOperateAction(mpFormRef)">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox } from "element-plus"
import {
  type userDataModel,
  getUsersApi,
  deleteUserApi,
  addUserApi,
  editUserApi,
  modifyPassApi,
  SwitchActiveApi
} from "@/api/authority/user"
import { getRolesApi } from "@/api/authority/role"
import { usePagination } from "@/hooks/usePagination"
import { useValidatePhone, useValidateEmail } from "@/hooks/useValidate"

defineOptions({
  name: "User"
})

const loading = ref<boolean>(false)
const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const tableData = ref<userDataModel[]>([])
let activeRow: userDataModel

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getUsersApi({ page: paginationData.currentPage, pageSize: paginationData.pageSize })
    if (res.code === 0) {
      tableData.value = res.data.list
      paginationData.total = res.data.total
    }
  } catch (error) {
    //
  }
  loading.value = false
}
getTableData()

// 修改密码对话框
const mpDialogVisible = ref<boolean>(false)

const mpFormRef = ref<FormInstance>()

const mpFormData = reactive({
  oldPassword: "",
  newPassword: "",
  rePassword: ""
})

const mpInitForm = () => {
  mpFormData.oldPassword = ""
  mpFormData.newPassword = ""
  mpFormData.rePassword = ""
}

const mpHandleClose = (done: Function) => {
  mpInitForm()
  done()
}

const equalToPassword = (rule: any, value: any, callback: any) => {
  if (mpFormData.newPassword !== value) {
    callback(new Error("两次输入的密码不一致"))
  } else {
    callback()
  }
}

const mpFormRules: FormRules = reactive({
  oldPassword: [{ required: true, trigger: "blur", message: "旧密码不能为空" }],
  newPassword: [
    { required: true, trigger: "blur", message: "新密码不能为空" },
    { min: 6, max: 20, message: "长度在 6 到 20 个字符", trigger: "blur" }
  ],
  rePassword: [
    { required: true, trigger: "blur", message: "确认密码不能为空" },
    { required: true, validator: equalToPassword, trigger: "blur" }
  ]
})

const mpCloseDialog = () => {
  mpFormRef.value?.resetFields()
  mpInitForm()
  mpDialogVisible.value = false
}

const modifyDialog = (row: userDataModel) => {
  activeRow = row
  mpDialogVisible.value = true
}

const mpOperateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      await modifyPassApi({
        id: activeRow.id,
        oldPassword: mpFormData.oldPassword,
        newPassword: mpFormData.newPassword
      })
        .then((res) => {
          if (res.code === 0) {
            ElMessage({ type: "success", message: res.msg })
            mpCloseDialog()
          }
        })
        .catch(() => {})
    }
  })
}

// 添加、编辑用户对话框

const initForm = () => {
  formRef.value?.resetFields()
  formData.username = ""
  formData.password = ""
  formData.phone = ""
  formData.email = ""
  formData.active = false
  formData.roleId = ""
}

const dialogVisible = ref<boolean>(false)
const handleClose = (done: Function) => {
  initForm()
  done()
}

const formRef = ref<FormInstance>()
const formData = reactive({
  username: "",
  password: "",
  phone: "",
  email: "",
  active: false,
  roleId: ""
})
const formRules: FormRules = reactive({
  username: [{ required: true, trigger: "blur", message: "请填写用户名" }],
  password: [{ required: true, trigger: "blur", message: "请填写密码" }],
  phone: [{ validator: useValidatePhone, trigger: "blur" }],
  email: [{ validator: useValidateEmail, trigger: "blur" }],
  roleId: [{ required: true, trigger: "change", message: "请选择角色" }]
})
const kind = ref("")
const title = ref("")
const addDialog = () => {
  kind.value = "Add"
  title.value = "新增用户"
  dialogVisible.value = true
}

const closeDialog = () => {
  formRef.value?.resetFields
  initForm()
  dialogVisible.value = false
}

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      if (kind.value === "Add") {
        const res = await addUserApi({
          username: formData.username,
          password: formData.password,
          phone: formData.phone,
          email: formData.email,
          active: formData.active,
          roleId: Number(formData.roleId)
        })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          getTableData()
        }
      } else if (kind.value === "Edit") {
        const res = await editUserApi({
          id: activeRow.id,
          username: formData.username,
          phone: formData.phone,
          email: formData.email,
          active: formData.active,
          roleId: Number(formData.roleId)
        })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          // 替换数据
          const index = tableData.value.indexOf(activeRow)
          tableData.value.splice(index, 1, res.data)
        }
      }
      closeDialog()
    }
  })
}

const deleteUserAction = (row: userDataModel) => {
  ElMessageBox.confirm("此操作将永久删除该用户, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      const index = tableData.value.indexOf(row)
      deleteUserApi({ id: row.id }).then((res) => {
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          tableData.value.splice(index, 1)
        }
      })
    })
    .catch(() => {})
}

// 分页
const handleSizeChange = (value: number) => {
  changePageSize(value)
  getTableData()
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  getTableData()
}

interface option {
  ID: string
  roleName: string
}
const roleOptions: option[] = []
const getRoleOption = async () => {
  const res = await getRolesApi()
  if (res.code === 0) {
    res.data.forEach((element) => {
      roleOptions.push({ ID: String(element.id), roleName: element.roleName })
    })
  }
}
getRoleOption()

const editDialog = (row: userDataModel) => {
  activeRow = row
  formData.username = row.username
  formData.phone = row.phone
  formData.email = row.email
  formData.active = row.active
  formData.roleId = String(row.roleId)
  kind.value = "Edit"
  title.value = "编辑用户"
  dialogVisible.value = true
}

// 切换用户状态
const switchAction = (id: number, active: boolean) => {
  SwitchActiveApi({ id: id, active: active })
    .then((res) => {
      if (res.code === 0) {
        if (active) {
          ElMessage({ type: "success", message: "启用成功" })
        } else {
          ElMessage({ type: "success", message: "禁用成功" })
        }
      }
    })
    .catch(() => {})
}
</script>
