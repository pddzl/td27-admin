<template>
  <div class="login-container">
    <div class="login-card">
      <div class="title">
        <img src="@/assets/layout/logo-text-2.png" />
      </div>
      <div class="content">
        <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" @keyup.enter="handleLogin">
          <el-form-item prop="username">
            <el-input
              v-model.trim="loginForm.username"
              placeholder="用户名"
              type="text"
              tabindex="1"
              prefix-icon="User"
              size="large"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model.trim="loginForm.password"
              placeholder="密码"
              type="password"
              tabindex="2"
              prefix-icon="Lock"
              size="large"
              show-password
            />
          </el-form-item>
          <el-form-item prop="captcha">
            <el-input
              v-model.trim="loginForm.captcha"
              placeholder="验证码"
              type="text"
              tabindex="3"
              prefix-icon="Key"
              maxlength="6"
              size="large"
            >
              <template #append>
                <el-image :src="codeUrl" @click="createCode" draggable="false">
                  <template #placeholder>
                    <el-icon><Picture /></el-icon>
                  </template>
                  <template #error>
                    <el-icon><Loading /></el-icon>
                  </template>
                </el-image>
              </template>
            </el-input>
          </el-form-item>
          <el-button :loading="loading" type="primary" size="large" @click.prevent="handleLogin"> 登 录 </el-button>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"
import { useUserStore } from "@/store/modules/user"
import { type FormInstance, FormRules } from "element-plus"
import { captcha, type ILoginRequestData } from "@/api/system/base"

const router = useRouter()
const loginFormRef = ref<FormInstance | null>(null)

/** 登录按钮 Loading */
const loading = ref(false)
/** 验证码图片 URL */
const codeUrl = ref("")
/** 登录表单数据 */
const loginForm: ILoginRequestData = reactive({
  username: "",
  password: "",
  captcha: "",
  captchaId: ""
})
/** 登录表单校验规则 */
const loginFormRules: FormRules = {
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { min: 6, max: 16, message: "长度在 6 到 20 个字符", trigger: "blur" }
  ],
  captcha: [{ required: true, message: "请输入验证码", trigger: "blur" }]
}

/** 登录逻辑 */
const handleLogin = () => {
  loginFormRef.value?.validate(async (valid) => {
    if (valid) {
      loading.value = true
      const res = await useUserStore().login({
        username: loginForm.username,
        password: loginForm.password,
        captcha: loginForm.captcha,
        captchaId: loginForm.captchaId
      })
      if (res) {
        router.push({ path: "/" })
      } else {
        createCode()
      }
      loading.value = false
    }
  })
}
/** 创建验证码 */
const createCode = () => {
  // 先清空验证码的输入
  loginForm.captcha = ""
  // 获取验证码
  captcha().then((res) => {
    codeUrl.value = res.data.picPath
    loginForm.captchaId = res.data.captchaId
  })
}

/** 初始化验证码 */
createCode()
</script>

<style lang="scss" scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  min-height: 100%;
  .login-card {
    width: 480px;
    border-radius: 20px;
    box-shadow: 0 0 10px #dcdfe6;
    background-color: #fff;
    overflow: hidden;
    .title {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 150px;
      img {
        height: 100%;
      }
    }
    .content {
      padding: 20px 50px 50px 50px;
      :deep(.el-input-group__append) {
        padding: 0;
        overflow: hidden;
        .el-image {
          width: 100px;
          height: 40px;
          border-left: 0px;
          user-select: none;
          cursor: pointer;
          text-align: center;
        }
      }
      .el-button {
        width: 100%;
        margin-top: 10px;
      }
    }
  }
}
</style>
