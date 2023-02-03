import { ref } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { type RouteRecordRaw } from "vue-router"
// import { constantRoutes, asyncRoutes } from "@/router"
import { constantRoutes } from "@/router"
import { getMenuList } from "@/api/system/menu"

// const hasPermission = (roles: string[], route: RouteRecordRaw) => {
//   if (route.meta && route.meta.roles) {
//     return roles.some((role) => {
//       if (route.meta?.roles !== undefined) {
//         return route.meta.roles.includes(role)
//       } else {
//         return false
//       }
//     })
//   } else {
//     return true
//   }
// }

// const filterAsyncRoutes = (routes: RouteRecordRaw[], roles: string[]) => {
//   const res: RouteRecordRaw[] = []
//   routes.forEach((route) => {
//     const r = { ...route }
//     if (hasPermission(roles, r)) {
//       if (r.children) {
//         r.children = filterAsyncRoutes(r.children, roles)
//       }
//       res.push(r)
//     }
//   })
//   return res
// }

const modules = import.meta.glob("../../views/**/*.vue", { eager: true })

export const usePermissionStore = defineStore("permission", () => {
  const routes = ref<RouteRecordRaw[]>([])
  const dynamicRoutes = ref<RouteRecordRaw[]>([])

  // const setRoutes = (roles: string[]) => {
  //   let accessedRoutes
  //   if (roles.includes("admin")) {
  //     accessedRoutes = asyncRoutes
  //   } else {
  //     accessedRoutes = filterAsyncRoutes(asyncRoutes, roles)
  //   }
  //   routes.value = constantRoutes.concat(accessedRoutes)
  //   dynamicRoutes.value = accessedRoutes
  // }

  const setRoutes = async () => {
    // const asyncRouterRes = await getMenuList()
    // const asyncRouterList: any[] = asyncRouterRes.data.list
    const asyncRouterList: any[] = [
      {
        id: 1,
        pid: 0,
        path: "/table",
        name: "Table",
        component: "layout/index.vue",
        redirect: "/table/host",
        meta: { title: "表格", elIcon: "Grid" }
      },
      {
        id: 2,
        pid: 1,
        path: "host",
        name: "Host",
        component: "table/host/index.vue",
        redirect: "",
        meta: { title: "主机" }
      },
      {
        id: 3,
        pid: 1,
        path: "container",
        name: "Container",
        component: "table/container/index.vue",
        redirect: "",
        meta: { title: "容器" }
      }
    ]
    // 初始化路由信息对象
    // const menusMap: MenusMap = {}
    const menusMap: any = {}
    asyncRouterList.map((v) => {
      const { path, name, component, redirect, meta } = v

      // 重新构建路由对象
      const item: RouteRecordRaw = {
        path,
        name,
        // component: () => import(`@/views/${component}`),
        component: () => modules[`../../views/${component}`],
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
export function usePermissionStoreHook() {
  return usePermissionStore(store)
}
