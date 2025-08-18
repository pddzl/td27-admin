import { request } from "@/http/axios_n"

export interface CasbinInfo {
  path: string
  method: string
}

interface reqCasbin {
  roleId: number
  casbinInfos: CasbinInfo[]
}

export function editCasbinApi(data: reqCasbin) {
  return request<ApiResponseData<null>>({
    url: "/casbin/editCasbin",
    method: "post",
    data
  })
}
