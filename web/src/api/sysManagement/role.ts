import { request } from "@/http/axios_n"

export interface roleData {
  roleName: string
  parentId?: number // 父角色ID（支持角色层级）
}

export interface roleDataModel extends roleData, Td27Model {
  parentId?: number
  parent?: roleDataModel
  children?: roleDataModel[]
  menus?: any[]
}

// 数据结构 - List
export type roleListData = ListData<roleDataModel[]>

// 获取角色列表
export function roleListApi(data: PageInfo) {
  return request<ApiResponseData<roleListData>>({
    url: "/role/list",
    method: "post",
    data
  })
}

// 获取角色树（包含层级关系）
export function roleTreeApi() {
  return request<ApiResponseData<roleDataModel[]>>({
    url: "/role/tree",
    method: "get"
  })
}

// 创建角色
export function roleCreateApi(data: roleData) {
  return request<ApiResponseData<roleDataModel>>({
    url: "/role/create",
    method: "post",
    data
  })
}

// 更新角色
export function roleUpdateApi(data: roleData & CId) {
  return request<ApiResponseData<roleDataModel>>({
    url: "/role/update",
    method: "post",
    data
  })
}

// 删除角色
export function roleDeleteApi(data: CId) {
  return request<ApiResponseData<null>>({
    url: "/role/delete",
    method: "post",
    data
  })
}

// 设置角色继承关系
export function setRoleInheritanceApi(data: { childRoleId: number, parentRoleId: number }) {
  return request<ApiResponseData<null>>({
    url: "/role/set-inheritance",
    method: "post",
    data
  })
}
