import { request } from "@/http/axios_n"

interface dictDetailData {
  label: string
  value: string
  sort: number
  dictId: number
  parentId?: number
  children?: dictDetailDataModel[]
}

export interface dictDetailDataModel extends dictDetailData, Td27Model {}

// List
export type dictDetailListData = ListData<dictDetailDataModel[]>

interface reqDictDetail extends PageInfo {
  dictId: number
}

export function getDictDetailApi(data: reqDictDetail) {
  return request<ApiResponseData<dictDetailListData>>({
    url: "/dictDetail/getDictDetail",
    method: "post",
    data
  })
}

export function addDictDetailApi(data: dictDetailData) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/addDictDetail",
    method: "post",
    data
  })
}

export function delDictDetailApi(data: CId) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/delDictDetail",
    method: "post",
    data
  })
}

export function editDictDetailApi(data: dictDetailData & CId) {
  return request<ApiResponseData<dictDetailDataModel>>({
    url: "/dictDetail/editDictDetail",
    method: "post",
    data
  })
}
