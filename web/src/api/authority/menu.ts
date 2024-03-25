import { request } from "@/utils/service"

export interface MenuData {
  pid: number
  name: string
  path: string
  redirect?: string
  component: string
  sort: number
  meta: {
    hidden?: boolean
    title: string
    elIcon?: string
    svgIcon?: string
    affix?: boolean
    keepAlive?: boolean
    alwaysShow?: boolean
  }
  children?: MenuData[]
}

export interface MenuDataModel extends MenuData, Td27Model {}

// List
// export type MenuListData = ListData<MenuDataModel[]>

// 获取动态路由
export function getMenus() {
  return request<ApiResponseData<MenuDataModel[]>>({
    url: "/menu/getMenus",
    method: "get"
  })
}

export function addMenuApi(data: MenuData) {
  return request<ApiResponseData<null>>({
    url: "menu/addMenu",
    method: "post",
    data
  })
}

export function editMenuApi(data: MenuData & CId) {
  return request<ApiResponseData<null>>({
    url: "menu/editMenu",
    method: "post",
    data
  })
}

export function deleteMenuApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "menu/deleteMenu",
    method: "post",
    data
  })
}

interface allMenus {
  list: MenuData[]
  menuIds: number[]
}

export function getElTreeMenusApi(data: CId) {
  return request<ApiResponseData<allMenus>>({
    url: "menu/getElTreeMenus",
    method: "post",
    data
  })
}
