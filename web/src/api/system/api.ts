import { request } from "@/utils/service"

export interface ApiData {
  id: number
  path: string
  apiGroup: string
  method: string
  description: string
}

export interface ApiDataPageInfo {
  list: ApiData[]
  total: number
  page: number
  pageSize: number
}

type ApiResponseData = IApiResponseData<ApiDataPageInfo>

interface reqApis extends PageInfo {
  path?: string
  apiGroup?: string
  method?: string
  description?: string
}

// 获取所有api
export function getApis(data: reqApis) {
  return request<ApiResponseData>({
    url: "/api/getApis",
    method: "post",
    data
  })
}
