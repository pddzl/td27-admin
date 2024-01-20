import { request } from "@/utils/service"

export function joinInBlacklistApi() {
  return request<ApiResponseData<null>>({
    url: "/jwt/joinInBlacklist",
    method: "post",
    data: {}
  })
}
