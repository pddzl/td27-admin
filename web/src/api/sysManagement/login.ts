import { request } from "@/http/axios_n"

export interface LoginRequestData {
  /** admin 或 editor */
  username: string
  /** 密码 */
  password: string
  /** 验证码 */
  captcha: string
  captchaId: string
}

type LoginCodeResponseData = ApiResponseData<{ picPath: string, captchaId: string }>
type LoginResponseData = ApiResponseData<{ token: string }>

// 获取验证码
export function captchaApi() {
  return request<LoginCodeResponseData>({
    url: "/captcha",
    method: "post"
  })
}

/** 登录并返回 Token */
export function loginApi(data: LoginRequestData) {
  return request<LoginResponseData>({
    url: "/login",
    method: "post",
    data
  })
}

// 登出
export function logoutApi() {
  return request<ApiResponseData<null>>({
    url: "/logout",
    method: "post",
    data: {}
  })
}
