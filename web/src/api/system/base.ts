import { request } from "@/utils/service"

export interface ILoginRequestData {
  /** admin 或 editor */
  username: string
  /** 密码 */
  password: string
  /** 验证码 */
  captcha: string
  captchaId: string
}

type LoginCodeResponseData = IApiResponseData<{ picPath: string; captchaId: string }>
type LoginResponseData = IApiResponseData<{ token: string }>

// 获取验证码
export const captcha = () => {
  return request<LoginCodeResponseData>({
    url: "/base/captcha",
    method: "post"
  })
}

/** 登录并返回 Token */
export function loginApi(data: ILoginRequestData) {
  return request<LoginResponseData>({
    url: "/base/login",
    method: "post",
    data
  })
}
