import { request } from "@/utils/service"

type UserInfoResponseData = IApiResponseData<{ username: string; roles: string[] }>

/** 获取用户详情 */
export function getUserInfoApi() {
  return request<UserInfoResponseData>({
    url: "/user/getUserInfo",
    method: "post",
    data: {}
  })
}
