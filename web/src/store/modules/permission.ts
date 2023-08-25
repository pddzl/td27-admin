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

/** 在 setup 外使用 */
export function usePermissionStoreHook() {
  return usePermissionStore(store)
}
