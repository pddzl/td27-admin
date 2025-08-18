import type { RouteRecordRaw } from "vue-router"
import type { MenuData } from "@/api/authority/menu"
import { getMenus } from "@/api/authority/menu"
import { formatRouter } from "@/common/utils/router_m"
import { pinia } from "@/pinia"
import { constantRoutes } from "@/router/index_n"

export const usePermissionStore = defineStore("permission", () => {
  const routes = ref<RouteRecordRaw[]>([])
  const dynamicRoutes = ref<RouteRecordRaw[]>([])
  const asyncRouterList = ref<MenuData[]>([])

  const setRoutes = async () => {
    // 获取动态路由
    const asyncRouterRes = await getMenus()
    if (asyncRouterRes.code === 0) {
      asyncRouterList.value = asyncRouterRes.data
    }

    // 格式化后端路由
    formatRouter(asyncRouterList.value, dynamicRoutes.value)

    // 404路由放最后
    // 添加404 ErrorPage
    dynamicRoutes.value.push({
      path: "/:pathMatch(.*)*",
      redirect: "/404",
      name: "ErrorPage",
      meta: {
        hidden: true
      }
    })

    // 合并静态路由，动态路由
    routes.value = constantRoutes.concat(dynamicRoutes.value)
  }

  const resetDynamicRouter = () => {
    routes.value = []
    dynamicRoutes.value = []
    asyncRouterList.value = []
  }

  return { routes, dynamicRoutes, asyncRouterList, setRoutes, resetDynamicRouter }
})

/**
 * @description 在 SPA 应用中可用于在 pinia 实例被激活前使用 store
 * @description 在 SSR 应用中可用于在 setup 外使用 store
 */
export function usePermissionStoreOutside() {
  return usePermissionStore(pinia)
}
