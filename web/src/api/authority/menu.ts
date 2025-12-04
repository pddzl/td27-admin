import type * as ElementPlusIconsVue from "@element-plus/icons-vue"
import type { SvgName } from "~virtual/svg-component"
import { request } from "@/http/axios_n"

type ElementPlusIconsName = keyof typeof ElementPlusIconsVue

export interface MenuData {
  pid: number
  name: string
  path: string
  redirect?: string
  component: string
  sort: number
  meta: {
    hidden?: boolean
    title?: string
    elIcon?: ElementPlusIconsName
    svgIcon?: SvgName
    affix?: boolean
    keepAlive?: boolean
    alwaysShow?: boolean
  }
  children?: MenuData[]
}

export interface MenuDataModel extends MenuData, Td27Model {}

// 获取动态路由
export function listMenuApi() {
  return request<ApiResponseData<MenuDataModel[]>>({
    url: "/menu/list",
    method: "get"
  })
}

export function createMenuApi(data: MenuData) {
  return request<ApiResponseData<null>>({
    url: "menu/create",
    method: "post",
    data
  })
}

export function updateMenuApi(data: MenuData & CId) {
  return request<ApiResponseData<null>>({
    url: "menu/update",
    method: "post",
    data
  })
}

export function deleteMenuApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "menu/delete",
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
