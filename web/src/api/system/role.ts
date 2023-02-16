import { request } from "@/utils/service"
import { type MenusData } from "./menu"

export interface roleData {
  ID: number
  roleName: string
  menus: MenusData[]
}

type RoleResponseData = IApiResponseData<roleData[]>

/** 获取用户详情 */
export function getRoles() {
  return request<RoleResponseData>({
    url: "/role/getRoles",
    method: "post",
    data: {}
  })
}
