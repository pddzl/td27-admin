import { request } from "@/http/axios_n"

// 角色信息
export interface RoleInfo {
  id: number
  roleName: string
}

interface userData {
  username: string
  phone: string
  email: string
  active: boolean
  roleIds: number[] // 多角色支持
  deptId?: number // 部门ID（用于数据权限）
}

export interface userDataModel extends Td27Model {
  username: string
  phone: string
  email: string
  active: boolean
  deptId?: number
  deptName?: string // 部门名称
  // 多角色支持
  roles: RoleInfo[]
  roleName: string // 主角色名称（兼容旧版）
  roleId: number // 主角色ID（兼容旧版）
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
export function userDeleteApi(data: { username: string }) {
  return request<ApiResponseData<null>>({
    url: "/user/delete",
    method: "post",
    data
  })
}

// 添加用户
export function userCreateApi(data: userData & { password: string }) {
  return request<ApiResponseData<userDataModel>>({
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
export function switchActiveApi(data: { active: boolean, username: string }) {
  return request<ApiResponseData<null>>({
    url: "/user/switchActive",
    method: "post",
    data
  })
}

// 获取部门列表
export function deptListApi() {
  return request<ApiResponseData<any[]>>({
    url: "/dept/list",
    method: "get"
  })
}
