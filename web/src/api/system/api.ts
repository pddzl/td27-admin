import { request } from "@/utils/service"

export interface ApiDataBase {
  path: string
  api_group: string
  method: string
  description: string
}

export interface ApiData extends ApiDataBase {
  ID: number
}

export interface ApiDataPageInfo {
  list: ApiData[]
  total: number
  page: number
  pageSize: number
}

interface reqApis extends PageInfo {
  path?: string
  api_group?: string
  method?: string
  description?: string
  orderKey?: string
  desc?: boolean
}

// 获取所有api 分页
export function getApisApi(data: reqApis) {
  return request<ApiResponseData<ApiDataPageInfo>>({
    url: "/api/getApis",
    method: "post",
    data
  })
}

interface children {
  key: string
  api_group: string
  path: string
  method: string
  description: string
}

export interface ApiTreeData {
  api_group: string
  children: children[]
}

interface ApiTreeAll {
  list: ApiTreeData[]
  checkedKey: string[]
}

// 获取所有api 不分页
export function getElTreeApisApi(data: reqId) {
  return request<ApiResponseData<ApiTreeAll>>({
    url: "/api/getElTreeApis",
    method: "post",
    data
  })
}

// 添加api
export function addApiApi(data: ApiDataBase) {
  return request<ApiResponseData<ApiData>>({
    url: "/api/addApi",
    method: "post",
    data
  })
}

// 删除api
export function deleteApiApi(data: reqId) {
  return request<ApiResponseData<null>>({
    url: "/api/deleteApi",
    method: "post",
    data
  })
}

interface reqEdit extends ApiDataBase {
  id: number
}

// 编辑api
export function editApiApi(data: reqEdit) {
  return request<ApiResponseData<null>>({
    url: "/api/editApi",
    method: "post",
    data
  })
}
