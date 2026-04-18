import { request } from "@/http/axios_n"

// 服务令牌响应
export interface ServiceToken {
  id: number
  name: string
  status: boolean
  expiresAt: number | null
  apiCount: number
  createdAt: number
}

// 创建服务令牌响应（包含明文token）
export interface CreateServiceTokenResp extends ServiceToken {
  token: string // 仅创建时返回一次
}

// 服务令牌详情
export interface ServiceTokenDetail extends ServiceToken {
  apiIds: number[]
}

// 创建服务令牌请求
export interface CreateServiceTokenReq {
  name: string
  expiresAt?: number // 过期时间戳(秒)，不传表示永不过期
  apiIds: number[]
}

// 更新服务令牌请求
export interface UpdateServiceTokenReq {
  id: number
  name: string
  status: boolean
  expiresAt?: number
  apiIds: number[]
}

// 列表查询请求
export interface ListServiceTokenReq extends PageInfo {
  name?: string
  status?: boolean
}

// 创建服务令牌
export function createServiceTokenApi(data: CreateServiceTokenReq) {
  return request<ApiResponseData<CreateServiceTokenResp>>({
    url: "/serviceToken/create",
    method: "post",
    data
  })
}

// 更新服务令牌
export function updateServiceTokenApi(data: UpdateServiceTokenReq) {
  return request<ApiResponseData<null>>({
    url: "/serviceToken/update",
    method: "post",
    data
  })
}

// 删除服务令牌
export function deleteServiceTokenApi(id: number) {
  return request<ApiResponseData<null>>({
    url: "/serviceToken/delete",
    method: "post",
    data: { id }
  })
}

// 获取服务令牌详情
export function getServiceTokenDetailApi(data: { id: number }) {
  return request<ApiResponseData<ServiceTokenDetail>>({
    url: "/serviceToken/detail",
    method: "post",
    data
  })
}

// 获取服务令牌列表
export function listServiceTokenApi(data: ListServiceTokenReq) {
  return request<ApiResponseData<ApiListData<ServiceToken[]>>>({
    url: "/serviceToken/list",
    method: "post",
    data
  })
}
