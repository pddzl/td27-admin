import { request } from "@/utils/service"

// 获取动态路由
export function getMenuList() {
  return request({
    url: "/menu/getMenuList",
    method: "get"
  })
}
