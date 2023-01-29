import { request } from "@/utils/service"

/** 获取用户详情 */
export function getUserInfoApi() {
  return request({
    url: "/user/getUserInfo",
    method: "post",
    data: {}
  })
}
