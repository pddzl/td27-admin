import { request } from "@/utils/service"

// 获取验证码
export const captcha = () => {
  return request({
    url: "/base/captcha",
    method: "post"
  })
}

export interface ILoginData {
  /** admin 或 editor */
  username: string
  /** 密码 */
  password: string
  /** 验证码 */
  code: string
}

/** 登录并返回 Token */
export function loginApi(data: ILoginData) {
  return request({
    url: "/base/login",
    method: "post",
    data
  })
}
