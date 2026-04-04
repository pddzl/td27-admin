import { request } from "@/http/axios_n"

// 统计数据
export interface DashboardStats {
  userCount: number
  roleCount: number
  deptCount: number
  operationCount: number
  cronCount: number
  activeCronCount: number
}

// 最近操作记录
export interface RecentOperation {
  id: number
  userName: string
  path: string
  method: string
  status: number
  respTime: number
  createdAt: string
}

// 系统信息
export interface SystemInfo {
  appName: string
  version: string
  goVersion: string
  os: string
  arch: string
  numCpu: number
  numGoroutine: number
  startTime: string
}

// 获取统计数据
export function getDashboardStatisticsApi() {
  return request<ApiResponseData<DashboardStats>>({
    url: "/dashboard/statistics",
    method: "get"
  })
}

// 获取最近操作记录
export function getRecentOperationsApi() {
  return request<ApiResponseData<RecentOperation[]>>({
    url: "/dashboard/recent-operations",
    method: "get"
  })
}

// 获取系统信息
export function getSystemInfoApi() {
  return request<ApiResponseData<SystemInfo>>({
    url: "/dashboard/system-info",
    method: "get"
  })
}
