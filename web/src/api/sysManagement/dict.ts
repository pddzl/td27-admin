import { request } from "@/http/axios_n"

interface dictData {
  chName: string
  enName: string
}

export interface dictDataModel extends dictData, Td27Model {}

export function dictListApi() {
  return request<ApiResponseData<dictDataModel[]>>({
    url: "/dict/list",
    method: "get"
  })
}

export function dictCreateApi(data: dictData) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/create",
    method: "post",
    data
  })
}

export function dictDeleteApi(data: CId) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/delete",
    method: "post",
    data
  })
}

export function dictUpdateApi(data: dictData & CId) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/update",
    method: "post",
    data
  })
}
