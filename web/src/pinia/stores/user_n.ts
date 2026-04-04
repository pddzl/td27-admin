import type { LoginRequestData } from "@/api/sysManagement/login"
import { loginApi, logoutApi } from "@/api/sysManagement/login"
import { getUserInfoApi, type RoleInfo } from "@/api/sysManagement/user"
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

  // 多角色支持
  const userInfo = reactive({
    id: 0,
    createdAt: "",
    username: "",
    phone: "",
    email: "",
    role: "",        // 主角色名称（兼容旧版）
    roleId: 0,       // 主角色ID（兼容旧版）
    roles: [] as RoleInfo[]  // 多角色列表
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
      
      // 多角色支持
      userInfo.roles = res.data.roles || []
      if (userInfo.roles.length > 0) {
        userInfo.role = userInfo.roles[0].roleName
        userInfo.roleId = userInfo.roles[0].id
      }
    }
  }

  /** 获取所有角色ID */
  const getAllRoleIds = () => {
    return userInfo.roles.map(r => r.id)
  }

  /** 检查是否有指定角色 */
  const hasRole = (roleId: number) => {
    return userInfo.roles.some(r => r.id === roleId)
  }

  /** 登出 */
  const logout = async () => {
    // 调用后端登出接口，删除缓存中的token
    try {
      await logoutApi()
    } catch (error) {
      console.error("Logout API failed:", error)
    }
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
    userInfo.phone = ""
    userInfo.email = ""
    userInfo.role = ""
    userInfo.roleId = 0
    userInfo.roles = []
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

  return { 
    token, 
    username, 
    getInfo, 
    userInfo, 
    login, 
    logout, 
    resetToken,
    getAllRoleIds,
    hasRole
  }
})

/**
 * @description 在 SPA 应用中可用于在 pinia 实例被激活前使用 store
 * @description 在 SSR 应用中可用于在 pinia 实例被激活前使用 store
 */
export function useUserStoreOutside() {
  return useUserStore(pinia)
}
