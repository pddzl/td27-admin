import { request } from "@/http/axios_n"

interface ApiData {
  path: string
  apiGroup: string
  method: string
  description: string
}

export interface ApiDataModel extends ApiData, Td27Model {}

// 数据结构 - List
export type ApiListData = ListData<ApiDataModel[]>

interface reqApis extends PageInfo {
  path?: string
  apiGroup?: string
  method?: string
  description?: string
  orderKey?: string
  desc?: boolean
}

// 获取所有api 分页
export function apiListApi(data: reqApis) {
  return request<ApiResponseData<ApiListData>>({
    url: "/api/list",
    method: "post",
    data
  })
}

interface ApiChild {
  id: number // 权限ID
  key: string // path,method 组合
  apiGroup: string
  path: string
  method: string
  description: string
}

export interface ApiTreeData {
  apiGroup: string
  children: ApiChild[]
}

interface ApiTreeAll {
  list: ApiTreeData[]
  checkedKey: string[] // 选中的 key 列表 (path,method)
  checkedIds: number[] // 选中的权限ID列表
}

// 获取所有api 不分页
export function apiGetElTreeApi(data: CId) {
  return request<ApiResponseData<ApiTreeAll>>({
    url: "/api/elTree",
    method: "post",
    data
  })
}

// 添加api
export function apiCreateApi(data: ApiData) {
  return request<ApiResponseData<ApiDataModel>>({
    url: "/api/create",
    method: "post",
    data
  })
}

// 删除api
export function apiDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/api/delete",
    method: "post",
    data
  })
}

// 批量删除api
export function apiDeleteByIdsApi(data: CIds) {
  return request<ApiResponseData<null>>({
    url: "/api/deleteByIds",
    method: "post",
    data
  })
}

// 编辑api
export function apiUpdateApi(data: ApiData & CId) {
  return request<ApiResponseData<null>>({
    url: "/api/update",
    method: "post",
    data
  })
}
