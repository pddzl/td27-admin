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
  id: number
  uuid: string
  username: string
  phone: number
  email: string
  active: boolean
  roles: [roleName: string]
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
