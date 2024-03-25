import { request } from "@/utils/service"
import { type MenuData } from "./menu"

interface roleData {
  roleName: string
  menus?: MenuData[]
}

export interface roleDataModel extends roleData, Td27Model {}

// List
// export type roleListData = ListData<roleDataModel[]>

/** 获取用户详情 */
export function getRolesApi() {
  return request<ApiResponseData<roleDataModel[]>>({
    url: "/role/getRoles",
    method: "post",
    data: {}
  })
}

export function addRoleApi(data: roleData) {
  return request<ApiResponseData<roleDataModel>>({
    url: "/role/addRole",
    method: "post",
    data: data
  })
}

export function deleteRoleApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/role/deleteRole",
    method: "post",
    data
  })
}

export function editRoleApi(data: roleData & CId) {
  return request<ApiResponseData<null>>({
    url: "/role/editRole",
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
