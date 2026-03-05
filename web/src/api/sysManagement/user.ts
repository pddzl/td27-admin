import { request } from "@/http/axios_n"

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
export function userListApi(data: PageInfo) {
  return request<ApiResponseData<userListData>>({
    url: "/user/list",
    method: "post",
    data
  })
}

// 删除用户
export function userDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/user/delete",
    method: "post",
    data
  })
}

// 添加用户
export function userCreateApi(data: userData & { password: string }) {
  return request<ApiResponseData<null>>({
    url: "/user/create",
    method: "post",
    data
  })
}

// 编辑用户
export function userUpdateApi(data: userData & CId) {
  return request<ApiResponseData<userDataModel>>({
    url: "/user/update",
    method: "post",
    data
  })
}

// 修改用户密码
interface reqModifyPass {
  oldPassword: string
  newPassword: string
}

export function modifyPasswdApi(data: reqModifyPass & CId) {
  return request<ApiResponseData<null>>({
    url: "/user/modifyPasswd",
    method: "post",
    data
  })
}

// 切换用户状态
export function switchActiveApi(data: { active: boolean } & CId) {
  return request<ApiResponseData<null>>({
    url: "/user/switchActive",
    method: "post",
    data
  })
}
