import { request } from "@/utils/service"

interface reqFiles extends PageInfo {
  name?: string
  orderKey?: string
  desc?: boolean
}

export interface FileData {
  ID: number
  fileName: string
  fullPath: string
  mime: string
  CreatedAt: string
}

export interface FileDataPageInfo {
  list: FileData[]
  total: number
  page: number
  pageSize: number
}

// 分页获取文件信息
export function getFileListApi(data: reqFiles) {
  return request<ApiResponseData<FileDataPageInfo>>({
    url: "/file/getFileList",
    method: "post",
    data
  })
}

// 下载文件
export const downloadApi = (params: { name: string }) => {
  return request({
    url: "/file/download",
    method: "get",
    params,
    responseType: "blob"
  })
}

// 删除文件
export const deleteApi = (params: { name: string }) => {
  return request<ApiResponseData<null>>({
    url: "/file/delete",
    method: "get",
    params
  })
}
