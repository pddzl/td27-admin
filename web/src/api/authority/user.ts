import { request } from "@/utils/service"

interface userData {
  username: string
  phone: string
  email: string
  active: boolean
  roleId: number
}

export interface userDataModel extends userData, Td27Model {
  roleName: string
}

// 数据结构 - List
export type userListData = ListData<userDataModel[]>

/** 获取用户详情 */
export function getUserInfoApi() {
  return request<ApiResponseData<userDataModel>>({
    url: "/user/getUserInfo",
    method: "get"
  })
}

/** 获取所有用户 */
export function getUsersApi(data: PageInfo) {
  return request<ApiResponseData<userListData>>({
    url: "/user/getUsers",
    method: "post",
    data: data
  })
}

// 删除用户
export function deleteUserApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/user/deleteUser",
    method: "post",
    data
  })
}

// 添加用户
export function addUserApi(data: userData & { password: string }) {
  return request<ApiResponseData<null>>({
    url: "/user/addUser",
    method: "post",
    data
  })
}

// 编辑用户
export function editUserApi(data: userData & CId) {
  return request<ApiResponseData<userDataModel>>({
    url: "/user/editUser",
    method: "post",
    data
  })
}

// 修改用户密码
interface reqModifyPass {
  oldPassword: string
  newPassword: string
}

export function modifyPassApi(data: reqModifyPass & CId) {
  return request<ApiResponseData<null>>({
    url: "/user/modifyPass",
    method: "post",
    data
  })
}

// 切换用户状态
export function SwitchActiveApi(data: { active: boolean } & CId) {
  return request<ApiResponseData<null>>({
    url: "/user/switchActive",
    method: "post",
    data
  })
}
