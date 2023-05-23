import { request } from "@/utils/service"

// type UserInfoResponseData = ApiResponseData<{ username: string; roles: string[] }>

/** 获取用户详情 */
export function getUserInfoApi() {
  return request<ApiResponseData<UsersResponse>>({
    url: "/user/getUserInfo",
    method: "get"
  })
}

export interface UsersResponse {
  createdAt: string
  ID: number
  username: string
  phone: string
  email: string
  active: boolean
  roleId: number
  role: string
}

export interface UsersResponsePageInfo {
  list: UsersResponse[]
  total: number
  page: number
  pageSize: number
}

type UsersResponseData = ApiResponseData<UsersResponsePageInfo>

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
  return request<ApiResponseData<null>>({
    url: "/user/deleteUser",
    method: "post",
    data
  })
}

export interface reqUser {
  username: string
  password: string
  phone: string
  email: string
  active: boolean
  roleId: number
}

// 添加用户
export function addUserApi(data: reqUser) {
  return request<ApiResponseData<null>>({
    url: "/user/addUser",
    method: "post",
    data
  })
}

interface reqEditUser {
  id: number
  username: string
  phone: string
  email: string
  active: boolean
  roleId: number
}

// 编辑用户
export function editUserApi(data: reqEditUser) {
  return request<ApiResponseData<UsersResponse>>({
    url: "/user/editUser",
    method: "post",
    data
  })
}

// 修改用户密码
interface reqModifyPass {
  id: number
  oldPassword: string
  newPassword: string
}

export function modifyPassApi(data: reqModifyPass) {
  return request<ApiResponseData<null>>({
    url: "/user/modifyPass",
    method: "post",
    data
  })
}

// 切换用户状态
interface reqSwitchActive {
  id: number
  active: boolean
}

export function SwitchActiveApi(data: reqSwitchActive) {
  return request<ApiResponseData<null>>({
    url: "/user/switchActive",
    method: "post",
    data
  })
}
