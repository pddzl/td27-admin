import { request } from "@/http/axios_n"

interface dictDetailData {
  label: string
  value: string
  sort: number
  dictId: number | null
  parentId?: number
  children?: dictDetailDataModel[]
  description: string
}

export interface dictDetailDataModel extends dictDetailData, Td27Model {}

// List
export type dictDetailListData = ListData<dictDetailDataModel[]>
export type dictDetailFlatData = dictDetailDataModel[]

interface reqDictDetail extends PageInfo {
  dictId: number
}

export function dictDetailListApi(data: reqDictDetail) {
  return request<ApiResponseData<dictDetailListData>>({
    url: "/dictDetail/list",
    method: "post",
    data
  })
}

export function dictDetailFlatApi(data: { dictId: number }) {
  return request<ApiResponseData<dictDetailFlatData>>({
    url: "/dictDetail/flat",
    method: "post",
    data
  })
}

export function dictDetailCreateApi(data: dictDetailData) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/create",
    method: "post",
    data
  })
}

export function dictDetailDeleteApi(data: CId) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/delete",
    method: "post",
    data
  })
}

export function dictDetailUpdateApi(data: dictDetailData & CId) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/update",
    method: "post",
    data
  })
}
