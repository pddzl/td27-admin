import type { MenuData } from "./menu"
import { request } from "@/http/axios_n"

interface roleData {
  roleName: string
  menus?: MenuData[]
}

export interface roleDataModel extends roleData, Td27Model {}

export type roleListData = ListData<roleDataModel[]>

/** 获取用户详情 */
export function roleListApi(data: PageInfo) {
  return request<ApiResponseData<roleListData>>({
    url: "/role/list",
    method: "post",
    data
  })
}

export function roleCreateApi(data: roleData) {
  return request<ApiResponseData<roleDataModel>>({
    url: "/role/create",
    method: "post",
    data
  })
}

export function roleDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/role/delete",
    method: "post",
    data
  })
}

export function roleUpdateApi(data: roleData & CId) {
  return request<ApiResponseData<null>>({
    url: "/role/update",
    method: "post",
    data
  })
}

export function updateRoleMenuApi(data: { roleId: number } & CIds) {
  return request<ApiResponseData<null>>({
    url: "/role/updateRoleMenu",
    method: "post",
    data
  })
}
