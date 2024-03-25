import { request } from "@/utils/service"

export const METHOD = {
  ClearTable: "clearTable",
  Shell: "shell"
}

export interface TableInfo {
  tableName: string
  compareField: string
  interval: string
}

interface addCronData {
  name: string
  method: string
  expression: string
  strategy?: string
  open?: boolean
  extraParams?: {
    tableInfo: TableInfo[]
    command: string
  }
  comment?: string
}

interface cronData {
  name: string
  method: string
  expression: string
  strategy: string
  open: boolean
  extraParams: {
    tableInfo: TableInfo[]
    command: string
  }
  entryId: number
  comment: string
}

export interface cronDataModel extends cronData, Td27Model {}

// 数据结构 - List
export type cronListData = ListData<cronDataModel[]>

// 分页获取cron
export function getCronListApi(data?: PageInfo) {
  return request<ApiResponseData<cronListData>>({
    url: "/cron/getCronList",
    method: "post",
    data
  })
}

// 添加cron
export function addCronApi(data: addCronData) {
  return request<ApiResponseData<cronDataModel>>({
    url: "/cron/addCron",
    method: "post",
    data
  })
}

// 切换cron
export function switchCronApi(data: { open: boolean } & CId) {
  return request<ApiResponseData<{ entryId: number }>>({
    url: "/cron/switchOpen",
    method: "post",
    data
  })
}

// 删除
export function deleteCronApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/cron/deleteCron",
    method: "post",
    data
  })
}

// 批量删除
export function deleteCronByIds(data: CIds) {
  return request<ApiResponseData<null>>({
    url: "/cron/deleteCronByIds",
    method: "post",
    data
  })
}

// 编辑
export function editCronApi(data: addCronData & CId) {
  return request<ApiResponseData<cronDataModel>>({
    url: "/cron/editCron",
    method: "post",
    data
  })
}
