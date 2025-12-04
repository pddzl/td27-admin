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
export function listApi(data: reqApis) {
  return request<ApiResponseData<ApiListData>>({
    url: "/api/list",
    method: "post",
    data
  })
}

interface children {
  key: string
  apiGroup: string
  path: string
  method: string
  description: string
}

export interface ApiTreeData {
  apiGroup: string
  children: children[]
}

interface ApiTreeAll {
  list: ApiTreeData[]
  checkedKey: string[]
}

// 获取所有api 不分页
export function getElTreeApi(data: CId) {
  return request<ApiResponseData<ApiTreeAll>>({
    url: "/api/getElTree",
    method: "post",
    data
  })
}

// 添加api
export function createApi(data: ApiData) {
  return request<ApiResponseData<ApiDataModel>>({
    url: "/api/create",
    method: "post",
    data
  })
}

// 删除api
export function deleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/api/delete",
    method: "post",
    data
  })
}

// 批量删除api
export function deleteByIdsApi(data: CIds) {
  return request<ApiResponseData<null>>({
    url: "/api/deleteByIds",
    method: "post",
    data
  })
}

// 编辑api
export function updateApi(data: ApiData & CId) {
  return request<ApiResponseData<null>>({
    url: "/api/update",
    method: "post",
    data
  })
}
