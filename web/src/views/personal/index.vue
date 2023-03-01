<template>
  <div class="app-container">
    <el-card>
      <el-tabs>
        <el-tab-pane label="基本信息">
          <div class="form-container">
            <el-form label-width="120px" :model="userInfoForm" :rules="userInfoRule" ref="userInfoFormRef">
              <el-form-item label="创建时间">
                <el-input style="width: 400px" v-model="userInfoForm.createdAt" disabled />
              </el-form-item>
              <el-form-item label="用户名" prop="username" required>
                <el-input style="width: 400px" v-model="userInfoForm.username" />
              </el-form-item>
              <el-form-item label="角色">
                <el-input style="width: 400px" v-model="userInfoForm.role" disabled />
              </el-form-item>
              <el-form-item label="手机号码" prop="phone">
                <el-input style="width: 400px" v-model="userInfoForm.phone" />
              </el-form-item>
              <el-form-item label="邮箱" prop="email">
                <el-input style="width: 400px" v-model="userInfoForm.email" />
              </el-form-item>
              <el-form-item style="margin-top: 40px">
                <el-button type="primary" style="margin-right: 20px" @click="handleEditUser(userInfoFormRef)"
                  >更新</el-button
                >
                <el-button type="primary" plain @click="toDefault">关闭</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
        <el-tab-pane label="修改密码">
          <div class="form-container">
            <el-form label-width="120px">
              <el-form-item label="旧密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
              <el-form-item label="新密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
              <el-form-item label="确认密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
              <el-form-item style="margin-top: 40px">
                <el-button type="primary" style="margin-right: 20px">确定</el-button>
                <el-button type="primary" plain @click="toDefault">关闭</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"
import { useRouter } from "vue-router"
import { type FormInstance, type FormRules, ElMessage } from "element-plus"
import { formatDateTime } from "@/utils/index"
import { useUserStore } from "@/store/modules/user"
import { editUserApi } from "@/api/system/user"
import { useValidatePhone, useValidateEmail } from "@/hooks/useValidate"

const userStore = useUserStore()
const router = useRouter()

const toDefault = () => {
  router.push("/")
}

const userInfoFormRef = ref<FormInstance>()
const userInfoForm = reactive({
  id: 0,
  createdAt: "",
  username: "",
  phone: "",
  email: "",
  role: "",
  roleId: 0
})

const userInfoRule: FormRules = reactive({
  username: [{ required: true, trigger: "blur", message: "请填写用户名" }],
  phone: [{ validator: useValidatePhone, trigger: "blur" }],
  email: [{ validator: useValidateEmail, trigger: "blur" }]
})

const getCache = () => {
  userInfoForm.id = userStore.userInfo.id
  userInfoForm.createdAt = formatDateTime(userStore.userInfo.createdAt)
  userInfoForm.username = userStore.userInfo.username
  userInfoForm.phone = userStore.userInfo.phone
  userInfoForm.email = userStore.userInfo.email
  userInfoForm.role = userStore.userInfo.role
  userInfoForm.roleId = userStore.userInfo.roleId
}
getCache()

const handleEditUser = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      editUserApi({
        id: userInfoForm.id,
        username: userInfoForm.username,
        phone: userInfoForm.phone,
        email: userInfoForm.email,
        active: true,
        roleId: userInfoForm.roleId
      })
        .then((res) => {
          if (res.code === 0) {
            ElMessage({ type: "success", message: res.msg })
          }
          // 刷新userinfo store
          userStore.getInfo()
        })
        .catch(() => {})
    }
  })
}
</script>

<style lang="scss" scoped>
.form-container {
  padding: 20px;
}
</style>
