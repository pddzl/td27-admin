const SYSTEM_NAME = 'td27-admin'

/** 缓存数据时用到的 Key */
class CacheKey {
  static TOKEN = `${SYSTEM_NAME}-token-key`
  static SIDEBAR_STATUS = `${SYSTEM_NAME}-sidebar-status-key`
  static ACTIVE_THEME_NAME = `${SYSTEM_NAME}-active-theme-name-key`
}

export default CacheKey
