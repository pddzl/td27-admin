/** 所有 api 接口的响应数据都应该准守该格式 */
interface ApiResponseData<T> {
  code: number
  data: T
  msg: string
}

/** 请求ID */
interface reqId {
  id: number
}

/** 分页 */
interface PageInfo {
  page: number
  pageSize: number
}
