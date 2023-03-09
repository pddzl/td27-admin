import { request } from "@/utils/service"

export interface CasbinInfo {
  path: string
  method: string
}

interface reqCasbin {
  roleId: number
  casbinInfos: CasbinInfo[]
}

export function editCasbinApi(data: reqCasbin) {
  return request<IApiResponseData<null>>({
    url: "/casbin/editCasbin",
    method: "post",
    data
  })
}
