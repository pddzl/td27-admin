import { ref } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { type RouteRecordRaw } from "vue-router"
import { constantRoutes } from "@/router"
import { formatRouter } from "@/utils/router"
import { type MenusData, getMenus } from "@/api/system/menu"

export const usePermissionStore = defineStore("permission", () => {
  const routes = ref<RouteRecordRaw[]>([])
  const dynamicRoutes = ref<RouteRecordRaw[]>([])
  const asyncRouterList = ref<MenusData[]>([])

  const setRoutes = async () => {
    const asyncRouterRes = await getMenus()
    asyncRouterList.value = asyncRouterRes.data

    formatRouter(asyncRouterList.value, dynamicRoutes.value)

    // 添加404 ErrorPage
    dynamicRoutes.value.push({
      path: "/:pathMatch(.*)*",
      redirect: "/404",
      name: "ErrorPage",
      meta: {
        hidden: true
      }
    })

    routes.value = constantRoutes.concat(dynamicRoutes.value)
  }

  const resetDynamicRouter = () => {
    routes.value = []
    dynamicRoutes.value = []
    asyncRouterList.value = []
  }

  return { routes, dynamicRoutes, asyncRouterList, setRoutes, resetDynamicRouter }
})

/** 在 setup 外使用 */
export function usePermissionStoreHook() {
  return usePermissionStore(store)
}
