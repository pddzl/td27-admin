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

interface AddCronData {
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

export interface CronData {
  ID: number
  CreatedAt: string
  UpdatedAt: string
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

type CronDataList = ApiListData<CronData[]>

// 分页获取cron
export function getCronListApi(data?: PageInfo) {
  return request<ApiResponseData<CronDataList>>({
    url: "/cron/getCronList",
    method: "post",
    data
  })
}

// 添加cron
export function addCronApi(data: AddCronData) {
  return request<ApiResponseData<CronData>>({
    url: "/cron/addCron",
    method: "post",
    data
  })
}

// 切换cron
export function switchCronApi(data: { id: number; open: boolean }) {
  return request<ApiResponseData<{ entryId: number }>>({
    url: "/cron/switchOpen",
    method: "post",
    data
  })
}

// 删除
export function deleteCronApi(data: { id: number }) {
  return request<ApiResponseData<null>>({
    url: "/cron/deleteCron",
    method: "post",
    data
  })
}

// 批量删除
export function deleteCronByIds(data: { ids: number[] }) {
  return request<ApiResponseData<null>>({
    url: "/cron/deleteCronByIds",
    method: "post",
    data
  })
}

export interface editData extends AddCronData {
  id: number
}

// 编辑
export function editCronApi(data: editData) {
  return request<ApiResponseData<CronData>>({
    url: "/cron/editCron",
    method: "post",
    data
  })
}
