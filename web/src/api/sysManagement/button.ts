import { request } from "@/http/axios_n"

export interface ButtonData {
  id: number
  buttonCode: string
  buttonName: string
  description: string
  pagePath: string
}

export interface ButtonDto extends ButtonData {
  hasPermission: boolean
}

export function createButtonApi(data: Omit<ButtonData, "id">) {
  return request<ApiResponseData<ButtonData>>({
    url: "/button/create",
    method: "post",
    data
  })
}

export function updateButtonApi(data: ButtonData) {
  return request<ApiResponseData<null>>({
    url: "/button/update",
    method: "post",
    data
  })
}

export function deleteButtonApi(id: number) {
  return request<ApiResponseData<null>>({
    url: "/button/delete",
    method: "post",
    data: { id }
  })
}

export function listButtonApi(data: { page: number, pageSize: number, pagePath?: string }) {
  return request<ApiResponseData<ApiListData<ButtonData[]>>>({
    url: "/button/list",
    method: "post",
    data
  })
}

export function getPageButtonsApi(pagePath: string) {
  return request<ApiResponseData<ButtonDto[]>>({
    url: "/button/page",
    method: "get",
    params: { pagePath }
  })
}

export function checkButtonApi(buttonCode: string) {
  return request<ApiResponseData<boolean>>({
    url: "/button/check",
    method: "post",
    data: { buttonCode }
  })
}

export function batchCheckButtonApi(buttonCodes: string[]) {
  return request<ApiResponseData<Record<string, boolean>>>({
    url: "/button/batchCheck",
    method: "post",
    data: { buttonCodes }
  })
}

export function getUserButtonsApi() {
  return request<ApiResponseData<string[]>>({
    url: "/button/user",
    method: "get"
  })
}
