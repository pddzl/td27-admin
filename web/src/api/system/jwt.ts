import { request } from "@/utils/service"

export function joinInBlacklistApi() {
  return request<IApiResponseData<null>>({
    url: "/jwt/joinInBlacklist",
    method: "post",
    data: {}
  })
}
