import { request } from "@/http/axios_n"

export interface CasbinInfo {
  path: string
  method: string
}

interface reqCasbin {
  roleId: number
  casbinInfos: CasbinInfo[]
}

export function casbinUpdateApi(data: reqCasbin) {
  return request<ApiResponseData<null>>({
    url: "/casbin/update",
    method: "post",
    data
  })
}
