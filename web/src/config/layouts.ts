import { getConfigLayout } from "@/utils/cache/local-storage"

/** 项目配置 */
export interface LayoutSettings {
  /** 是否显示 Settings Panel */
  showSettings: boolean
  /** 布局模式 */
  layoutMode: "left" | "top" | "left-top"
  /** 是否显示标签栏 */
  showTagsView: boolean
  /** 是否显示 Logo */
  showLogo: boolean
  /** 是否固定 Header */
  fixedHeader: boolean
  /** 是否显示切换主题按钮 */
  showThemeSwitch: boolean
  /** 是否显示全屏按钮 */
  showScreenfull: boolean
  /** 是否缓存标签栏 */
  cacheTagsView: boolean
}

export const layoutSettings: LayoutSettings = getConfigLayout() ?? {
  layoutMode: "left",
  showSettings: true,
  showTagsView: true,
  fixedHeader: true,
  showLogo: true,
  showThemeSwitch: true,
  showScreenfull: true,
  cacheTagsView: false
}

export default layoutSettings
