import type { LoginRequestData } from "@/api/base/login"
import { getUserInfoApi } from "@/api/authority/user"
import { loginApi } from "@/api/base/login"
import { pinia } from "@/pinia"
import { usePermissionStoreOutside } from "@/pinia/stores/permission_n"
import { useSettingsStore } from "./settings"
import { useTagsViewStore } from "./tags-view"

export const useUserStore = defineStore("user", () => {
  const username = ref<string>("")

  const tagsViewStore = useTagsViewStore()

  const settingsStore = useSettingsStore()

  // mine start
  const token = ref<string>(window.localStorage.getItem("token") || "")

  const userInfo = reactive({
    id: 0,
    createdAt: "",
    username: "",
    phone: "",
    email: "",
    role: "",
    roleId: 0
  })

  const permissionStore = usePermissionStoreOutside()

  /** 登录 */
  const login = async (loginData: LoginRequestData) => {
    const res = await loginApi({
      username: loginData.username,
      password: loginData.password,
      captcha: loginData.captcha,
      captchaId: loginData.captchaId
    })
    if (res.code === 0) {
      token.value = res.data.token
    }
  }

  /** 获取用户详情 */
  const getInfo = async () => {
    const res = await getUserInfoApi()
    if (res.code === 0) {
      username.value = res.data.username
      userInfo.id = res.data.id
      userInfo.createdAt = res.data.createdAt
      userInfo.username = res.data.username
      userInfo.phone = res.data.phone
      userInfo.email = res.data.email
      userInfo.role = res.data.roleName
      userInfo.roleId = res.data.roleId
    }
  }

  /** 登出 */
  const logout = () => {
    resetUserInfo()
    token.value = ""
    permissionStore.resetDynamicRouter()
    resetTagsView()
  }

  /** 重置 Token */
  const resetToken = () => {
    token.value = ""
    window.localStorage.removeItem("token")
  }

  const resetUserInfo = () => {
    userInfo.username = ""
    userInfo.id = 0
    userInfo.createdAt = ""
    userInfo.username = ""
    userInfo.phone = ""
    userInfo.email = ""
    userInfo.role = ""
    userInfo.roleId = 0
  }
  // mine end

  // 重置 Visited Views 和 Cached Views
  const resetTagsView = () => {
    if (!settingsStore.cacheTagsView) {
      tagsViewStore.delAllVisitedViews()
      tagsViewStore.delAllCachedViews()
    }
  }

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value)
    }
  )

  return { token, username, getInfo, userInfo, login, logout, resetToken }
})

/**
 * @description 在 SPA 应用中可用于在 pinia 实例被激活前使用 store
 * @description 在 SSR 应用中可用于在 setup 外使用 store
 */
export function useUserStoreOutside() {
  return useUserStore(pinia)
}
