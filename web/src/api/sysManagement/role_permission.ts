import { request } from "@/http/axios_n"

// 更新角色权限
export function updateRolePermissionApi(data: { role_id: number, permission_ids: number[], domain: string }) {
  return request<ApiResponseData<null>>({
    url: "/role_permission/update",
    method: "post",
    data
  })
}
