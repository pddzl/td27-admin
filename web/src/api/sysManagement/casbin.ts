import { request } from "@/http/axios_n"

// 旧的Casbin格式（保留兼容）
export interface CasbinInfo {
  path: string
  method: string
}

// 新的统一权限格式
interface UpdateRoleAPIReq {
  roleId: number
  apiPermissionIds: number[]  // API权限ID列表
}

// 更新角色API权限（新格式）
export function updateRoleAPIPermissionsApi(data: UpdateRoleAPIReq) {
  return request<ApiResponseData<null>>({
    url: "/casbin/update",
    method: "post",
    data
  })
}

// 旧的API（保留兼容）
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
