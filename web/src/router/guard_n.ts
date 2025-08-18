import type { Router } from "vue-router"
import { setRouteChange } from "@@/composables/useRouteListener"
import { useTitle } from "@@/composables/useTitle"
import NProgress from "nprogress"
import { usePermissionStore } from "@/pinia/stores/permission_n"
import { useUserStore } from "@/pinia/stores/user_n"
import { isWhiteList } from "@/router/whitelist"

NProgress.configure({ showSpinner: false })

const { setTitle } = useTitle()

const LOGIN_PATH = "/login"

export function registerNavigationGuard(router: Router) {
  // 全局前置守卫
  router.beforeEach(async (to, _from) => {
    NProgress.start()
    const userStore = useUserStore()
    const permissionStore = usePermissionStore()
    // 如果没有登录
    if (!userStore.token) {
      // 如果在免登录的白名单中，则直接进入
      if (isWhiteList(to)) return true
      // 其他没有访问权限的页面将被重定向到登录页面
      return LOGIN_PATH
    }
    // 如果已经登录，并准备进入 Login 页面，则重定向到主页
    if (to.path === LOGIN_PATH) return "/"
    // 如果用户已经获得其权限角色
    if (userStore.userInfo.username !== "") return true
    // 否则要重新获取权限角色
    try {
      // 获取用户信息
      await userStore.getInfo()
      // 获取路由
      await permissionStore.setRoutes()
      // 将'有访问权限的动态路由' 添加到 Router 中
      permissionStore.dynamicRoutes.forEach((route) => {
        router.addRoute(route)
      })
      // 设置 replace: true, 因此导航将不会留下历史记录
      return { ...to, replace: true }
    } catch (error) {
      // 过程中发生任何错误，都直接重置 Token，并重定向到登录页面
      userStore.resetToken()
      ElMessage.error((error as Error).message || "路由守卫发生错误")
      return LOGIN_PATH
    }
  })

  // 全局后置钩子
  router.afterEach((to) => {
    setRouteChange(to)
    setTitle(to.meta.title)
    NProgress.done()
  })
}
