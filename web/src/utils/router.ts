/**
 * 后端路由转换为RouteRecordRaw格式
 * 路由component转换为vite import
 * 递归处理子路由
 */
import { type RouteRecordRaw } from "vue-router"
import { type MenusData } from "@/api/system/menu"

const modules = import.meta.glob("../views/**/*.vue", { eager: true })

export function dynamicImport(component: string) {
  return new Promise((resolve) => {
    if (component === "Layout") {
      resolve(import("@/layouts/index.vue"))
    }
    resolve(modules[`../views/${component}`])
  })
}

export function formatRouter(menuList: MenusData[], formatMenu: RouteRecordRaw[]) {
  for (const menu of menuList) {
    const fMenu: RouteRecordRaw = {
      name: menu.name,
      path: menu.path,
      redirect: menu.redirect,
      component: () => dynamicImport(menu.component),
      meta: menu.meta,
      children: []
    }
    formatMenu.push(fMenu)
    // 递归处理子路由
    if (Array.isArray(menu.children) && menu.children.length > 0) {
      formatRouter(menu.children, fMenu.children)
    } else {
      // @ts-ignore
      fMenu.children = null
    }
  }
}
