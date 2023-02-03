import { ref, watch } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { resetRouter } from "@/router"
import { getUserInfoApi } from "@/api/system/user"
import { type ILoginRequestData, loginApi } from "@/api/system/base"

export const useUserStore = defineStore("user", () => {
  const token = ref<string>(window.localStorage.getItem("token") || "")
  const roles = ref<string[]>([])
  const username = ref<string>("")

  /** 设置角色数组 */
  const setRoles = (value: string[]) => {
    roles.value = value
  }
  /** 登录 */
  const login = (loginData: ILoginRequestData) => {
    return new Promise((resolve, reject) => {
      loginApi({
        username: loginData.username,
        password: loginData.password,
        captcha: loginData.captcha,
        captchaId: loginData.captchaId
      })
        .then((res) => {
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
  /** 登出 */
  const logout = () => {
    token.value = ""
    roles.value = []
    resetRouter()
  }
  /** 重置 Token */
  const resetToken = () => {
    token.value = ""
    roles.value = []
  }

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value)
    }
  )

  return { token, roles, username, setRoles, login, getInfo, logout, resetToken }
})

/** 在 setup 外使用 */
export function useUserStoreHook() {
  return useUserStore(store)
}
