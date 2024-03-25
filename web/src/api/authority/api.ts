import { request } from "@/utils/service"

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
export function getApisApi(data: reqApis) {
  return request<ApiResponseData<ApiListData>>({
    url: "/api/getApis",
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
export function getElTreeApisApi(data: CId) {
  return request<ApiResponseData<ApiTreeAll>>({
    url: "/api/getElTreeApis",
    method: "post",
    data
  })
}

// 添加api
export function addApiApi(data: ApiData) {
  return request<ApiResponseData<ApiDataModel>>({
    url: "/api/addApi",
    method: "post",
    data
  })
}

// 删除api
export function deleteApiApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/api/deleteApi",
    method: "post",
    data
  })
}

// 批量删除api
export const deleteApiByIdApi = (data: CIds) => {
  return request<ApiResponseData<null>>({
    url: "/api/deleteApiById",
    method: "post",
    data
  })
}

// 编辑api
export function editApiApi(data: ApiData & CId) {
  return request<ApiResponseData<null>>({
    url: "/api/editApi",
    method: "post",
    data
  })
}
