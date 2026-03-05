import { request } from "@/http/axios_n"

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
export function cronListApi(data?: PageInfo) {
  return request<ApiResponseData<cronListData>>({
    url: "/cron/list",
    method: "post",
    data
  })
}

// 添加cron
export function cronCreateApi(data: addCronData) {
  return request<ApiResponseData<cronDataModel>>({
    url: "/cron/create",
    method: "post",
    data
  })
}

// 切换cron
export function cronSwitchOpenApi(data: { open: boolean } & CId) {
  return request<ApiResponseData<{ entryId: number }>>({
    url: "/cron/switchOpen",
    method: "post",
    data
  })
}

// 删除
export function cronDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/cron/delete",
    method: "post",
    data
  })
}

// 批量删除
export function cronDeleteByIds(data: CIds) {
  return request<ApiResponseData<null>>({
    url: "/cron/deleteByIds",
    method: "post",
    data
  })
}

// 编辑
export function cronUpdateApi(data: addCronData & CId) {
  return request<ApiResponseData<cronDataModel>>({
    url: "/cron/update",
    method: "post",
    data
  })
}
