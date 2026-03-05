import { request } from "@/http/axios_n"

interface reqFiles extends PageInfo {
  name?: string
  orderKey?: string
  desc?: boolean
}

interface fileData {
  fileName: string
  fullPath: string
  mime: string
}

export interface fileDataModel extends fileData, Td27Model {}

// List
export type fileListData = ListData<fileDataModel[]>

// 分页获取文件信息
export function fileListApi(data: reqFiles) {
  return request<ApiResponseData<fileListData>>({
    url: "/file/list",
    method: "post",
    data
  })
}

// 下载文件
export function fileDownloadApi(params: { name: string }) {
  return request({
    url: "/file/download",
    method: "get",
    params,
    responseType: "blob"
  })
}

// 删除文件
export function fileDeleteApi(params: { name: string }) {
  return request<ApiResponseData<null>>({
    url: "/file/delete",
    method: "get",
    params
  })
}
