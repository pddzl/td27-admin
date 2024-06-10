import { reactive, ref, watch } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { useTagsViewStore } from "./tags-view"
import { useSettingsStore } from "./settings"
import { getUserInfoApi } from "@/api/authority/user"
import { type LoginRequestData, loginApi } from "@/api/base/logReg"
import { usePermissionStoreHook } from "@/store/modules/permission"

export const useUserStore = defineStore("user", () => {
  const token = ref<string>(window.localStorage.getItem("token") || "")
  const username = ref<string>("")
  const userInfo = reactive({
    id: 0,
    createdAt: "",
    username: "",
    phone: "",
    email: "",
    role: "",
    roleId: 0
  })
  const tagsViewStore = useTagsViewStore()
  const permissionStore = usePermissionStoreHook()
  const settingsStore = useSettingsStore()

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
    _resetTagsView()
  }

  /** 重置 Token */
  const resetToken = () => {
    token.value = ""
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

  /** 重置 visited views 和 cached views */
  const _resetTagsView = () => {
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

  return { token, username, userInfo, login, getInfo, logout, resetToken }
})

/** 在 setup 外使用 */
export function useUserStoreHook() {
  return useUserStore(store)
}
