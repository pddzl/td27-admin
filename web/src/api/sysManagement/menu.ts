// import type * as ElementPlusIconsVue from "@element-plus/icons-vue"
import type { SvgName } from "~virtual/svg-component"
import { request } from "@/http/axios_n"

// type ElementPlusIconsName = keyof typeof ElementPlusIconsVue

export interface MenuData {
  menu_name: string
  icon?: SvgName | ""
  path: string
  component: string
  redirect: string
  parentId: number
  sort: number
  hidden: boolean
  keepAlive: boolean
  title: string
  affix: boolean
  alwaysShow: boolean
  children?: MenuDataModel[]
}

export interface MenuDataModel extends MenuData, Td27Model {}

// 获取动态路由
export function menuListApi() {
  return request<ApiResponseData<MenuDataModel[]>>({
    url: "/menu/list",
    method: "get"
  })
}

export function menuCreateApi(data: MenuData) {
  return request<ApiResponseData<null>>({
    url: "/menu/create",
    method: "post",
    data
  })
}

export function menuUpdateApi(data: MenuData & CId) {
  return request<ApiResponseData<null>>({
    url: "/menu/update",
    method: "post",
    data
  })
}

export function menuDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/menu/delete",
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
    url: "/menu/elTree",
    method: "post",
    data
  })
}
