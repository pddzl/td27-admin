/** 动态路由配置 */
interface RouteSettings {
  /**
   * 是否开启三级及其以上路由缓存功能？
   * 1. 开启后会进行路由降级（把三级及其以上的路由转化为二级路由）
   * 2. 由于都会转成二级路由，所以二级及其以上路由有内嵌子路由将会失效
   */
  thirdLevelRouteCache: boolean
}

const routeSettings: RouteSettings = {
  thirdLevelRouteCache: false
}

export default routeSettings
