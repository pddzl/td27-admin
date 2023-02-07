import { ref } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { type RouteRecordRaw } from "vue-router"
import { constantRoutes } from "@/router"
import { dynamicImport } from "@/utils/asyncRouter"
import { type MenusData, getMenus } from "@/api/system/menu"

export const usePermissionStore = defineStore("permission", () => {
  const routes = ref<RouteRecordRaw[]>([])
  const dynamicRoutes = ref<RouteRecordRaw[]>([])
  /** 左侧栏菜单管理的数据 */
  interface MenusDataFormat extends MenusData {
    children: MenusData[]
  }
  const menusDataFormatList = ref<MenusDataFormat[]>([])
  const asyncRouterList = ref<MenusData[]>([])

  const setRoutes = async () => {
    const asyncRouterRes = await getMenus()
    asyncRouterList.value = asyncRouterRes.data
    // const asyncRouterList: any[] = [
    //   {
    //     id: 1,
    //     pid: 0,
    //     path: "/table",
    //     name: "Table",
    //     component: "layout/index.vue",
    //     redirect: "/table/host",
    //     meta: { title: "表格", elIcon: "Grid" }
    //   },
    //   {
    //     id: 2,
    //     pid: 1,
    //     path: "host",
    //     name: "Host",
    //     component: "table/host/index.vue",
    //     redirect: "",
    //     meta: { title: "主机" }
    //   },
    //   {
    //     id: 3,
    //     pid: 1,
    //     path: "container",
    //     name: "Container",
    //     component: "table/container/index.vue",
    //     redirect: "",
    //     meta: { title: "容器" }
    //   }
    // ]

    // 初始化路由信息对象
    const menusMap: any = {}
    const menusMapRaw: any = {}
    asyncRouterList.value.map((v) => {
      const { path, name, component, redirect, meta } = v

      // 重新构建路由对象
      const item = {
        name,
        path,
        component: () => dynamicImport(component),
        redirect,
        meta
      }

      // 判断是否为根节点
      if (v.pid === 0) {
        menusMap[v.id] = item
        menusMapRaw[v.id] = v
      } else {
        !menusMap[v.pid].children && (menusMap[v.pid].children = [])
        menusMap[v.pid].children.push(item)

        !menusMapRaw[v.pid].children && (menusMapRaw[v.pid].children = [])
        menusMapRaw[v.pid].children.push(v)
      }
    })

    dynamicRoutes.value = Object.values(menusMap)
    menusDataFormatList.value = Object.values(menusMapRaw)

    routes.value = constantRoutes.concat(dynamicRoutes.value)
  }

  return { routes, dynamicRoutes, menusDataFormatList, setRoutes }
})

/** 在 setup 外使用 */
export function usePermissionStoreHook() {
  return usePermissionStore(store)
}
