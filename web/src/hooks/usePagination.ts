import { reactive } from "vue"

interface DefaultPaginationData {
  total: number
  currentPage: number
  pageSizes: number[]
  pageSize: number
  layout: string
}

interface IPaginationData {
  total?: number
  currentPage?: number
  pageSizes?: number[]
  pageSize?: number
  layout?: string
}

/** 默认的分页参数 */
const defaultPaginationData: DefaultPaginationData = {
  total: 0,
  currentPage: 1,
  pageSizes: [10, 20, 50],
  pageSize: 10,
  layout: "total, sizes, prev, pager, next, jumper"
}

export function usePagination(initialPaginationData: IPaginationData = {}) {
  /** 合并分页参数 */
  const paginationData = reactive({ ...defaultPaginationData, ...initialPaginationData })

  /** 改变当前页码 */
  const changeCurrentPage = (value: number) => {
    paginationData.currentPage = value
  }

  /** 改变页面大小 */
  const changePageSize = (value: number) => {
    paginationData.pageSize = value
  }

  return { paginationData, changeCurrentPage, changePageSize }
}
