import { request } from "@/http/axios_n"

export interface Dept {
  id?: number
  deptName: string
  parentId: number
  path?: string
  sort: number
  status: boolean
  createdAt?: string
  children?: Dept[]
}

export interface CreateDeptReq {
  deptName: string
  parentId: number
  path?: string
  sort: number
  status: boolean
}

export interface UpdateDeptReq {
  id: number
  deptName: string
  parentId: number
  path?: string
  sort: number
  status: boolean
}

export interface DeptListReq {
  deptName?: string
  status?: boolean
}

/** 获取部门列表（树形） */
export function deptListApi(data: DeptListReq) {
  return request<ApiResponseData<Dept[]>>({
    url: "/dept/list",
    method: "post",
    data
  })
}

/** 创建部门 */
export function createDeptApi(data: CreateDeptReq) {
  return request<ApiResponseData<null>>({
    url: "/dept/create",
    method: "post",
    data
  })
}

/** 更新部门 */
export function updateDeptApi(data: UpdateDeptReq) {
  return request<ApiResponseData<null>>({
    url: "/dept/update",
    method: "post",
    data
  })
}

/** 删除部门 */
export function deleteDeptApi(data: { id: number }) {
  return request<ApiResponseData<null>>({
    url: "/dept/delete",
    method: "post",
    data
  })
}

/** 获取部门树（用于选择器） */
export function getElTreeDeptsApi() {
  return request<ApiResponseData<{
    tree: Dept[]
    ids: number[]
  }>>({
    url: "/dept/getElTreeDepts",
    method: "post"
  })
}
