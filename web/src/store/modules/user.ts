import { ref, watch } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { usePermissionStore } from "./permission"
import { removeToken, setToken } from "@/utils/cache/cookies"
import router, { resetRouter } from "@/router"
import { getUserInfoApi } from "@/api/system/user"
import { type ILoginData, loginApi } from "@/api/system/base"
import { type RouteRecordRaw } from "vue-router"

export const useUserStore = defineStore("user", () => {
  const token = ref<string>(window.localStorage.getItem("token") || "")
  const roles = ref<string[]>([])
  const username = ref<string>("")

  /** 设置角色数组 */
  const setRoles = (value: string[]) => {
    roles.value = value
  }
  /** 登录 */
  const login = (loginData: ILoginData) => {
    return new Promise((resolve, reject) => {
      loginApi({
        username: loginData.username,
        password: loginData.password,
        captcha: loginData.captcha,
        captchaId: loginData.captchaId
      })
        .then((res: any) => {
          // setToken(res.data.token)
          token.value = res.data.token
          resolve(true)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }
  /** 获取用户详情 */
  const getInfo = () => {
    return new Promise((resolve, reject) => {
      getUserInfoApi()
        .then((res: any) => {
          roles.value = res.data.roles
          username.value = res.data.username
          resolve(res)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }
  /** 切换角色 */
  const changeRoles = async (role: string) => {
    const newToken = "token-" + role
    token.value = newToken
    setToken(newToken)
    await getInfo()
    const permissionStore = usePermissionStore()
    permissionStore.setRoutes(roles.value)
    resetRouter()
    permissionStore.dynamicRoutes.forEach((item: RouteRecordRaw) => {
      router.addRoute(item)
    })
  }
  /** 登出 */
  const logout = () => {
    removeToken()
    token.value = ""
    roles.value = []
    resetRouter()
  }
  /** 重置 Token */
  const resetToken = () => {
    removeToken()
    token.value = ""
    roles.value = []
  }

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value)
    }
  )

  return { token, roles, username, setRoles, login, getInfo, changeRoles, logout, resetToken }
})

/** 在 setup 外使用 */
export function useUserStoreHook() {
  return useUserStore(store)
}
