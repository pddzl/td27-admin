import { request } from "@/utils/service"

export interface ApiDataBase {
  path: string
  apiGroup: string
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
  apiGroup?: string
  method?: string
  description?: string
}

// 获取所有api 分页
export function getApisApi(data: reqApis) {
  return request<IApiResponseData<ApiDataPageInfo>>({
    url: "/api/getApis",
    method: "post",
    data
  })
}

// 获取所有api 不分页
export function getApisTreeApi() {
  return request<IApiResponseData<ApiData[]>>({
    url: "/api/getApisTree",
    method: "get"
  })
}

// 添加api
export function addApiApi(data: ApiDataBase) {
  return request<IApiResponseData<ApiData>>({
    url: "/api/addApi",
    method: "post",
    data
  })
}

// 删除api
export function deleteApiApi(data: reqId) {
  return request<IApiResponseData<null>>({
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
  return request<IApiResponseData<null>>({
    url: "/api/editApi",
    method: "post",
    data
  })
}
