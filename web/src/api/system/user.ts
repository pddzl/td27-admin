import { request } from "@/utils/service"

type UserInfoResponseData = IApiResponseData<{ username: string; roles: string[] }>

/** 获取用户详情 */
export function getUserInfoApi() {
  return request<UserInfoResponseData>({
    url: "/user/getUserInfo",
    method: "post",
    data: {}
  })
}

export interface UsersResponse {
  ID: number
  username: string
  phone: string
  email: string
  active: boolean
  role: string
}

export interface UsersResponsePageInfo {
  list: UsersResponse[]
  total: number
  page: number
  pageSize: number
}

type UsersResponseData = IApiResponseData<UsersResponsePageInfo>

/** 获取所有用户 */
export function getUsersApi(data: PageInfo) {
  return request<UsersResponseData>({
    url: "/user/getUsers",
    method: "post",
    data: data
  })
}

// 删除用户
export function deleteUserApi(data: reqId) {
  return request<IApiResponseData<null>>({
    url: "/user/deleteUser",
    method: "delete",
    data
  })
}

export interface reqUser {
  username: string
  password: string
  phone: string
  email: string
  active: boolean
  roleID: number
}

// 添加用户
export function addUserApi(data: reqUser) {
  return request<IApiResponseData<null>>({
    url: "/user/addUser",
    method: "post",
    data
  })
}
