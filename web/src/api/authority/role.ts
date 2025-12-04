import type { MenuData } from "./menu"
import { request } from "@/http/axios_n"

interface roleData {
  roleName: string
  menus?: MenuData[]
}

export interface roleDataModel extends roleData, Td27Model {}

/** 获取用户详情 */
export function listRoleApi() {
  return request<ApiResponseData<roleDataModel[]>>({
    url: "/role/list",
    method: "post",
    data: {}
  })
}

export function createRoleApi(data: roleData) {
  return request<ApiResponseData<roleDataModel>>({
    url: "/role/create",
    method: "post",
    data
  })
}

export function deleteRoleApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/role/delete",
    method: "post",
    data
  })
}

export function updateRoleApi(data: roleData & CId) {
  return request<ApiResponseData<null>>({
    url: "/role/update",
    method: "post",
    data
  })
}

export function editRoleMenuApi(data: { roleId: number } & CIds) {
  return request<ApiResponseData<null>>({
    url: "/role/editRoleMenu",
    method: "post",
    data
  })
}
