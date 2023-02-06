import { request } from "@/utils/service"

// 获取动态路由
export function getMenus() {
  return request({
    url: "/menu/getMenus",
    method: "get"
  })
}
