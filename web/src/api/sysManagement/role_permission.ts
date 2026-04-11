import { request } from "@/http/axios_n"

// 更新角色权限
export function rebuildRolePermissionApi(data: { role_id: number, domain_ids: number[], domain: string }) {
  return request<ApiResponseData<null>>({
    url: "/role_permission/rebuild",
    method: "post",
    data
  })
}
