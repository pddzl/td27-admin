import { ref } from "vue"
import { defineStore } from "pinia"
import { type RouteRecordRaw } from "vue-router"
import store from "@/store"
import { constantRoutes } from "@/router"
import { getMenuList } from "@/api/system/menu"

// type MenusMap = {
//   [key: number]: RouteRecordRaw
// }

export const useRouterStore = defineStore("router", () => {
  // return { asyncRouter, SetAsyncRouter }

  const routes = ref<RouteRecordRaw[]>([])
  const dynamicRoutes = ref<RouteRecordRaw[]>([])

  const setRoutes = async () => {
    const asyncRouterRes = await getMenuList()
    const asyncRouterList: any[] = asyncRouterRes.data.list
    // 初始化路由信息对象
    // const menusMap: MenusMap = {}
    const menusMap: any = {}
    asyncRouterList.map((v) => {
      const { path, name, component, redirect, meta } = v

      // 重新构建路由对象
      const item: RouteRecordRaw = {
        path,
        name,
        component: () => import(`@/views/${component}`),
        redirect,
        meta
      }

      // 判断是否为根节点
      if (v.pid === 0) {
        menusMap[v.id] = item
      } else {
        !menusMap[v.pid].children && (menusMap[v.pid].children = [])
        menusMap[v.pid].children.push(item)
      }
    })

    dynamicRoutes.value = Object.values(menusMap)

    routes.value = constantRoutes.concat(dynamicRoutes.value)
  }

  return { routes, dynamicRoutes, setRoutes }
})

/** 在 setup 外使用 */
export function useRouterStoreHook() {
  return useRouterStore(store)
}
