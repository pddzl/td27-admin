import { request } from "@/utils/service"

export interface MenusData {
  id: number
  pid: number
  name: string
  path: string
  redirect?: string
  component: string
  meta: {
    hidden?: boolean
    title: string
    elIcon?: string
    svgIcon?: string
    affix?: boolean
    keepAlive?: boolean
  }
  children?: MenusData[]
}

type MenusResponseData = ApiResponseData<MenusData[]>

// 获取动态路由
export function getMenus() {
  return request<MenusResponseData>({
    url: "/menu/getMenus",
    method: "get"
  })
}

interface reqMenu {
  pid: number
  name?: string
  path: string
  redirect?: string
  component: string
  meta: {
    hidden?: boolean
    title?: string
    icon?: string
    affix?: boolean
    keepAlive?: boolean
  }
}

export function addMenuApi(data: reqMenu) {
  return request<ApiResponseData<null>>({
    url: "menu/addMenu",
    method: "post",
    data
  })
}

interface editReq extends reqMenu {
  id: number
}

export function editMenuApi(data: editReq) {
  return request<ApiResponseData<null>>({
    url: "menu/editMenu",
    method: "post",
    data
  })
}

export function deleteMenuApi(data: reqId) {
  return request<ApiResponseData<null>>({
    url: "menu/deleteMenu",
    method: "post",
    data
  })
}

interface allMenus {
  list: MenusData[]
  menuIds: number[]
}

export function getElTreeMenusApi(data: reqId) {
  return request<ApiResponseData<allMenus>>({
    url: "menu/getElTreeMenus",
    method: "post",
    data
  })
}
