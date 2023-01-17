import { type RouteRecordRaw, createRouter, createWebHashHistory } from 'vue-router'

const Layout = () => import("@/layout/index.vue")

// 常驻路由
export const constantRoutes: RouteRecordRaw[] = []

const router = createRouter({
  history: createWebHashHistory(),
  routes: constantRoutes
})

export default router
