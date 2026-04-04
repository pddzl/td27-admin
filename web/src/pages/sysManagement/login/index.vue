<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"
import type { LoginRequestData } from "@/api/sysManagement/login"
import { reactive, ref, onMounted } from "vue"
import { useRouter } from "vue-router"
import { captchaApi } from "@/api/sysManagement/login"
import { useUserStore } from "@/pinia/stores/user_n"
import { useTheme, type ThemeName } from "@@/composables/useTheme"
import { Sunny, Moon } from "@element-plus/icons-vue"

const router = useRouter()
const loginFormRef = ref<FormInstance | null>(null)

// Theme
const { activeThemeName, setTheme, initTheme } = useTheme()

// Initialize theme on mount
onMounted(() => {
  initTheme()
})

// Toggle between light and dark theme
function toggleTheme() {
  const newTheme: ThemeName = activeThemeName.value === "dark" ? "normal" : "dark"
  const event = new MouseEvent("click", {
    clientX: window.innerWidth / 2,
    clientY: window.innerHeight / 2
  })
  setTheme(event, newTheme)
}

/** 登录按钮 Loading */
const loading = ref(false)
/** 验证码图片 URL */
const codeUrl = ref("")
/** 登录表单数据 */
const loginFormData: LoginRequestData = reactive({
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
function handleLogin() {
  loginFormRef.value?.validate(async (valid) => {
    if (valid) {
      loading.value = true
      await useUserStore()
        .login({
          username: loginFormData.username,
          password: loginFormData.password,
          captcha: loginFormData.captcha,
          captchaId: loginFormData.captchaId
        })
        .then(() => {
          router.push({ path: "/" })
        })
        .catch(() => {
          createCode()
        })
        .finally(() => {
          loading.value = false
        })
    }
  })
}

/** 创建验证码 */
function createCode() {
  // 先清空验证码的输入
  loginFormData.captcha = ""
  // 获取验证码
  captchaApi().then((res) => {
    codeUrl.value = res.data.picPath
    loginFormData.captchaId = res.data.captchaId
  })
}

/** 初始化验证码 */
createCode()
</script>

<template>
  <div class="login-container bg">
    <!-- Theme Switch Button -->
    <div class="theme-switch-wrapper">
      <el-tooltip :content="activeThemeName === 'dark' ? '切换亮色模式' : '切换暗黑模式'" placement="left">
        <el-button
          circle
          :icon="activeThemeName === 'dark' ? Sunny : Moon"
          @click="toggleTheme"
          class="theme-switch-btn"
        />
      </el-tooltip>
    </div>
    
    <div class="login-card">
      <p class="p1">
        TD27 ADMIN
      </p>
      <p class="p2">
        Enjoy yourself with Golang and Vue
      </p>
      <el-form ref="loginFormRef" :model="loginFormData" :rules="loginFormRules" @keyup.enter="handleLogin">
        <el-form-item prop="username">
          <el-input
            v-model.trim="loginFormData.username" placeholder="用户名" type="text" tabindex="1" prefix-icon="User"
            size="large"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model.trim="loginFormData.password" placeholder="密码" type="password" tabindex="2"
            prefix-icon="Lock" size="large" show-password
          />
        </el-form-item>
        <el-form-item prop="captcha">
          <el-input
            v-model.trim="loginFormData.captcha" placeholder="验证码" type="text" tabindex="3" prefix-icon="Key"
            maxlength="6" size="large"
          >
            <template #append>
              <el-image :src="codeUrl" @click="createCode" draggable="false">
                <template #placeholder>
                  <el-icon>
                    <Picture />
                  </el-icon>
                </template>
                <template #error>
                  <el-icon>
                    <Loading />
                  </el-icon>
                </template>
              </el-image>
            </template>
          </el-input>
        </el-form-item>
        <el-button :loading="loading" type="primary" size="large" @click.prevent="handleLogin">
          登 录
        </el-button>
      </el-form>
    </div>
    <div class="footer">
      <span>Copyright © 2023 Pddzl | </span>
      <a href="https://github.com/pddzl/td27-admin" target="_blank">
        <img src="@@/assets/images/github.png" alt="github">
      </a>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.login-container {
  position: relative;
  width: 100%;
  min-height: 100%;

  .login-card {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -60%);
    width: 480px;
    max-width: 90%;
    overflow: hidden;

    padding: 20px 50px 50px 50px;

    .p1 {
      text-align: center;
      font-size: 33px;
      color: rgba(0, 0, 0, 0.796);
    }

    .p2 {
      margin-top: -20px;
      margin-bottom: 40px;
      text-align: center;
      font-size: 14px;
      color: rgba(0, 0, 0, 0.45);
    }

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

  .footer {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    bottom: 15%;

    span {
      font-size: 14px;
      color: rgba(0, 0, 0, 0.45);
    }

    img {
      width: 20px;
      height: 20px;
      transform: translateY(25%);
    }
  }
}

.bg {
  background: #f0f2f5 url(@@/assets/images/layouts/background.svg) no-repeat 50%;
  background-size: 100%;
}

// Dark mode styles
:global(.dark) .login-container,
:global(.dark-blue) .login-container {
  &.bg {
    background: #141414;
  }
  
  .login-card {
    background: #1f1f1f;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
    
    .p1 {
      color: rgba(255, 255, 255, 0.85);
    }
    
    .p2 {
      color: rgba(255, 255, 255, 0.45);
    }
  }
  
  .footer span {
    color: rgba(255, 255, 255, 0.45);
  }
}

// Theme switch button styles
.theme-switch-wrapper {
  position: absolute;
  top: 24px;
  right: 24px;
  z-index: 10;
}

.theme-switch-btn {
  width: 44px;
  height: 44px;
  font-size: 20px;
  transition: all 0.3s;
  
  &:hover {
    transform: scale(1.1);
  }
}
</style>
