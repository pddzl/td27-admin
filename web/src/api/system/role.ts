import { request } from "@/utils/service"
import { type MenusData } from "./menu"

export interface roleData {
  ID: number
  roleName: string
  menus: MenusData[]
}

type RoleResponseData = ApiResponseData<roleData[]>

/** 获取用户详情 */
export function getRolesApi() {
  return request<RoleResponseData>({
    url: "/role/getRoles",
    method: "post",
    data: {}
  })
}

export interface reqRole {
  roleName: string
}

export function addRoleApi(data: reqRole) {
  return request<ApiResponseData<roleData>>({
    url: "/role/addRole",
    method: "post",
    data: data
  })
}

export function deleteRoleApi(data: reqId) {
  return request<ApiResponseData<null>>({
    url: "/role/deleteRole",
    method: "post",
    data
  })
}

interface reqEditRole {
  id: number
  roleName: string
}

export function editRoleApi(data: reqEditRole) {
  return request<ApiResponseData<null>>({
    url: "/role/editRole",
    method: "post",
    data
  })
}

interface reqEditRE {
  roleId: number
  ids: number[]
}

export function editRoleMenuApi(data: reqEditRE) {
  return request<ApiResponseData<null>>({
    url: "/role/editRoleMenu",
    method: "post",
    data
  })
}
