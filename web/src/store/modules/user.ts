import { reactive, ref, watch } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { useTagsViewStore } from "./tags-view"
import { getUserInfoApi } from "@/api/system/user"
import { type LoginRequestData, loginApi } from "@/api/system/base"
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

  /** 登录 */
  const login = async (loginData: LoginRequestData): Promise<boolean> => {
    try {
      const res = await loginApi({
        username: loginData.username,
        password: loginData.password,
        captcha: loginData.captcha,
        captchaId: loginData.captchaId
      })
      if (res.code === 0) {
        token.value = res.data.token
        return true
      }
    } catch (error) {
      // console.log(error)
      return false
    }
    return false
  }

  /** 获取用户详情 */
  const getInfo = () => {
    return new Promise((resolve, reject) => {
      getUserInfoApi()
        .then((res) => {
          username.value = res.data.username
          userInfo.id = res.data.ID
          userInfo.createdAt = res.data.createdAt
          userInfo.username = res.data.username
          userInfo.phone = res.data.phone
          userInfo.email = res.data.email
          userInfo.role = res.data.role
          userInfo.roleId = res.data.roleId
          resolve(res)
        })
        .catch((error) => {
          reject(error)
        })
    })
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
    tagsViewStore.delAllVisitedViews()
    tagsViewStore.delAllCachedViews()
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
