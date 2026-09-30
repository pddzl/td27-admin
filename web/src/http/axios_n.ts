import type { AxiosInstance, AxiosRequestConfig } from "axios"
import axios from "axios"
import { get, merge } from "lodash-es"
import { useUserStore } from "@/pinia/stores/user_n"

/** 退出登录并跳转到登录页（携带 redirect 以便登录后回跳） */
function logout() {
  useUserStore().logout()
  const redirect = encodeURIComponent(`${location.pathname}${location.search}`)
  location.href = `/login?redirect=${redirect}`
}

/** In-flight request tracking for cancellation and GET dedupe */
const pendingControllers = new Set<AbortController>()
const inflightGets = new Map<string, Promise<unknown>>()

/** Cancel all in-flight requests (called on route change so stale responses never land in the new page) */
export function cancelPendingRequests() {
  pendingControllers.forEach(controller => controller.abort())
  pendingControllers.clear()
  inflightGets.clear()
}

/** Throttle identical error toasts within 2s to avoid message flooding */
let lastErrorMessage = ""
let lastErrorTime = 0
function toastError(message: string) {
  const now = Date.now()
  if (message === lastErrorMessage && now - lastErrorTime < 2000) return
  lastErrorMessage = message
  lastErrorTime = now
  ElMessage.error(message)
}

/** 创建请求实例 */
function createInstance() {
  // 创建一个 axios 实例命名为 instance
  const instance = axios.create()
  // 请求拦截器
  instance.interceptors.request.use(
    // 发送之前
    config => config,
    // 发送失败
    error => Promise.reject(error)
  )
  // 响应拦截器（可根据具体业务作出相应的调整）
  instance.interceptors.response.use(
    (response) => {
      // apiData 是 api 返回的数据
      const apiData = response.data
      // 二进制数据则直接返回
      const responseType = response.config.responseType
      if (responseType === "blob" || responseType === "arraybuffer") return apiData
      // 这个 code 是和后端约定的业务 code
      const code = apiData.code
      // 如果没有 code, 代表这不是项目后端开发的 api
      if (code === undefined) {
        toastError("非本系统的接口")
        return Promise.reject(new Error("非本系统的接口"))
      }
      switch (code) {
        case 0:
          // 本系统采用 code === 0 来表示没有业务错误
          return apiData

        default:
          // 不是正确的 code
          if (apiData.data && apiData.data.reload) {
            useUserStore().logout()
          }

          toastError(apiData.msg || "Error")
          return Promise.reject(apiData.msg || "Error")
      }
    },
    (error) => {
      // Cancelled requests (e.g. on route change) fail silently
      if (axios.isCancel(error)) return new Promise(() => {})
      // status 是 HTTP 状态码
      const status = get(error, "response.status")
      const message = get(error, "response.data.message")
      switch (status) {
        case 400:
          error.message = "请求错误"
          break
        case 401:
          // Token 过期时
          error.message = message || "未授权"
          logout()
          break
        case 403:
          error.message = message || "拒绝访问"
          break
        case 404:
          error.message = "请求地址出错"
          break
        case 408:
          error.message = "请求超时"
          break
        case 500:
          error.message = "服务器内部错误"
          break
        case 501:
          error.message = "服务未实现"
          break
        case 502:
          error.message = "网关错误"
          break
        case 503:
          error.message = "服务不可用"
          break
        case 504:
          error.message = "网关超时"
          break
        case 505:
          error.message = "HTTP 版本不受支持"
          break
      }
      toastError(error.message)
      return Promise.reject(error)
    }
  )
  return instance
}

/** 创建请求方法 */
function createRequest(instance: AxiosInstance) {
  return <T>(config: AxiosRequestConfig): Promise<T> => {
    // 默认配置
    const defaultConfig: AxiosRequestConfig = {
      // 接口地址
      baseURL: import.meta.env.VITE_BASE_URL,
      // 请求头
      headers: {
        // 携带 Token
        "x-token": useUserStore().token,
        "Content-Type": "application/json"
      },
      // 请求体
      data: {},
      // 请求超时
      timeout: 5000,
      // 跨域请求时是否携带 Cookies
      withCredentials: false
    }
    // 将默认配置 defaultConfig 和传入的自定义配置 config 进行合并成为 mergeConfig
    const mergeConfig = merge(defaultConfig, config)

    const method = (mergeConfig.method ?? "get").toLowerCase()
    const key = `${method}:${mergeConfig.url}:${JSON.stringify(mergeConfig.params)}:${JSON.stringify(mergeConfig.data)}`

    // Share identical in-flight GET requests
    if (method === "get") {
      const pending = inflightGets.get(key)
      if (pending) return pending as Promise<T>
    }

    // Attach an AbortController so route changes can cancel stale requests
    const controller = new AbortController()
    if (!mergeConfig.signal) mergeConfig.signal = controller.signal
    pendingControllers.add(controller)

    const promise = instance(mergeConfig).finally(() => {
      pendingControllers.delete(controller)
      if (method === "get") inflightGets.delete(key)
    }) as Promise<T>
    if (method === "get") inflightGets.set(key, promise)
    return promise
  }
}

/** 用于请求的实例 */
const instance = createInstance()

/** 用于请求的方法 */
export const request = createRequest(instance)
