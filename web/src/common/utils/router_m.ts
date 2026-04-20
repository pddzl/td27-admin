/**
 * 后端路由转换为RouteRecordRaw格式
 * 路由component转换为vite import
 * 递归处理子路由
 */
import type { RouteRecordRaw } from "vue-router"
import type { MenuData } from "@/api/sysManagement/menu"

const modules = import.meta.glob("@/pages/**/*.vue")

export function dynamicImport(component: string) {
  if (component === "Layout") {
    return () => import("@/layouts/index.vue")
  }

  const mod = modules[`/src/pages/${component}`]
  if (!mod) {
    console.warn(`组件 ${component} 不存在！`)
    return () => import("@/pages/error/404.vue") // 可以默认返回一个404组件
  }
  return mod
}

export function formatRouter(menuList: MenuData[], formatMenu: RouteRecordRaw[]) {
  for (const menu of menuList) {
    const fMenu: RouteRecordRaw = {
      name: menu.menu_name,
      path: menu.path,
      redirect: menu.redirect,
      component: dynamicImport(menu.component),
      meta: {
        svgIcon: menu.icon || undefined,
        affix: menu.affix,
        alwaysShow: menu.alwaysShow,
        hidden: menu.hidden,
        keepAlive: menu.keepAlive,
        title: menu.title
      },
      children: []
    }
    formatMenu.push(fMenu)
    // 递归处理子路由
    if (Array.isArray(menu.children) && menu.children.length > 0) {
      formatRouter(menu.children, fMenu.children!)
    } else {
      delete (fMenu as any).children
    }
  }
}
