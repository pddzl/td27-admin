import { request } from "@/http/axios_n"

interface dictData {
  chName: string
  enName: string
}

export interface dictDataModel extends dictData, Td27Model {}

// List
// export type dictListData = ListData<dictDataModel[]>

export function getDictApi() {
  return request<ApiResponseData<dictDataModel[]>>({
    url: "/dict/getDict",
    method: "get"
  })
}

export function addDictApi(data: dictData) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/addDict",
    method: "post",
    data
  })
}

export function delDictApi(data: CId) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/delDict",
    method: "post",
    data
  })
}

export function editDictApi(data: dictData & CId) {
  return request<ApiResponseData<dictDataModel>>({
    url: "/dict/editDict",
    method: "post",
    data
  })
}
