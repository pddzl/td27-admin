import { type RouteRecordRaw, createRouter, createWebHashHistory, createWebHistory } from "vue-router"

const Layout = () => import("@/layout/index.vue")

/** 常驻路由 */
export const constantRoutes: RouteRecordRaw[] = [
  {
    path: "/redirect",
    component: Layout,
    meta: {
      hidden: true
    },
    children: [
      {
        path: "/redirect/:path(.*)",
        component: () => import("@/views/redirect/index.vue")
      }
    ]
  },
  {
    path: "/403",
    component: () => import("@/views/error-page/403.vue"),
    meta: {
      hidden: true
    }
  },
  {
    path: "/404",
    component: () => import("@/views/error-page/404.vue"),
    meta: {
      hidden: true
    },
    alias: "/:pathMatch(.*)*"
  },
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    meta: {
      hidden: true
    }
  }
  // {
  //   path: "/",
  //   component: Layout,
  //   redirect: "/dashboard",
  //   children: [
  //     {
  //       path: "dashboard",
  //       component: () => import("@/views/dashboard/index.vue"),
  //       name: "Dashboard",
  //       meta: {
  //         title: "首页",
  //         svgIcon: "dashboard",
  //         affix: true
  //       }
  //     }
  //   ]
  // },
  // {
  //   path: "/profile",
  //   component: Layout,
  //   meta: { hidden: true },
  //   redirect: "/profile/index",
  //   children: [
  //     {
  //       path: "index",
  //       component: () => import("@/views/profile/index.vue"),
  //       name: "Profile",
  //       meta: {
  //         title: "个人中心"
  //       }
  //     }
  //   ]
  // }
]

const router = createRouter({
  history:
    import.meta.env.VITE_ROUTER_HISTORY === "hash"
      ? createWebHashHistory(import.meta.env.VITE_PUBLIC_PATH)
      : createWebHistory(import.meta.env.VITE_PUBLIC_PATH),
  routes: constantRoutes
})

export default router
