/** 统一处理 Cookie */

import CacheKey from "@/constants/cacheKey"
import Cookies from "js-cookie"

export const setToken = (token: string) => {
  Cookies.set(CacheKey.TOKEN, token)
}
