/** 所有 api 接口的响应数据都应该准守该格式 */
interface ApiResponseData<T> {
  code: number
  data: T
  msg: string
}

/** 请求IDs */
interface CIds {
  ids: number[]
}

/** 分页 */
interface PageInfo {
  page: number
  pageSize: number
}

/** get list */
interface ApiListData<T> {
  list: T
  total: number
  page: number
  pageSize: number
}

interface Td27Model {
  id: number // 主键ID
  createdAt: string // 创建时间
  updatedAt: string // 更新时间
  deletedAt: string // 删除时间
}

/** get list */
interface ListData<T> {
  list: T
  total: number
  page: number
  pageSize: number
}

/** 请求ID */
interface CId {
  id: number
}

interface CIds {
  ids: number[]
}
