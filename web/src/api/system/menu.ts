import { request } from "@/utils/service"

export interface MenusData {
  id: number
  pid: number
  name: string
  path: string
  redirect: string
  component: string
  meta: {
    hidden: boolean
    title: string
    elIcon: string
    svgIcon: string
    affix: string
  }
}

/** 左侧栏菜单管理的数据 */
export interface MenusDataFormat extends MenusData {
  children: MenusData[]
}

type MenusResponseData = IApiResponseData<MenusData[]>

// 获取动态路由
export function getMenus() {
  return request<MenusResponseData>({
    url: "/menu/getMenus",
    method: "get"
  })
}
