import { getConfigLayout } from "@/utils/cache/local-storage"

/** 项目配置类型 */
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
  /** 是否显示页脚 Footer */
  showFooter: boolean
  /** 是否显示切换主题按钮 */
  showThemeSwitch: boolean
  /** 是否显示全屏按钮 */
  showScreenfull: boolean
  /** 是否显示搜索按钮 */
  showSearchMenu: boolean
  /** 是否缓存标签栏 */
  cacheTagsView: boolean
  /** 开启系统水印 */
  showWatermark: boolean
}

/** 默认配置 */
const defaultSettings: LayoutSettings = {
  layoutMode: "left",
  showSettings: true,
  showTagsView: true,
  fixedHeader: true,
  showFooter: true,
  showLogo: true,
  showThemeSwitch: true,
  showScreenfull: true,
  showSearchMenu: true,
  cacheTagsView: false,
  showWatermark: true,
}

/** 项目配置 */
export const layoutSettings: LayoutSettings = { ...defaultSettings, ...getConfigLayout() }
