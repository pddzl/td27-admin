import { request } from "@/utils/service"

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
export function getFileListApi(data: reqFiles) {
  return request<ApiResponseData<fileListData>>({
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
