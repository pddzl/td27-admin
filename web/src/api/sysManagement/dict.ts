import { request } from "@/http/axios_n"

interface dictData {
  cn_name: string
  en_name: string
}

export interface DictModel extends dictData, Td27Model {}

interface listReq extends Partial<PageInfo> {
  cn_name?: string
  en_name?: string
}

export function dictListApi(data: listReq) {
  return request<ApiResponseData<ListData<DictModel[]>>>({
    url: "/dict/list",
    method: "post",
    data
  })
}

export function dictCreateApi(data: dictData) {
  return request<ApiResponseData<DictModel>>({
    url: "/dict/create",
    method: "post",
    data
  })
}

export function dictDeleteApi(data: CId) {
  return request<ApiResponseData<DictModel>>({
    url: "/dict/delete",
    method: "post",
    data
  })
}

export function dictUpdateApi(data: dictData & CId) {
  return request<ApiResponseData<null>>({
    url: "/dict/update",
    method: "post",
    data
  })
}
