const SYSTEM_NAME = "td27-admin"

/** 缓存数据时用到的 Key */
class CacheKey {
  static readonly SIDEBAR_STATUS = `${SYSTEM_NAME}-sidebar-status-key`
  static readonly ACTIVE_THEME_NAME = `${SYSTEM_NAME}-active-theme-name-key`
}

export default CacheKey
