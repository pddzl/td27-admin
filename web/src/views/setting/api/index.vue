<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form ref="searchFormRef" :inline="true" :model="searchFormData">
        <el-form-item prop="path" label="路径">
          <el-input v-model="searchFormData.path" placeholder="路径" />
        </el-form-item>
        <el-form-item prop="group" label="API组">
          <el-input v-model="searchFormData.apiGroup" placeholder="API组" />
        </el-form-item>
        <el-form-item prop="method" label="方法">
          <el-select v-model="searchFormData.method" placeholder="方法">
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item prop="description" label="描述">
          <el-input v-model="searchFormData.description" placeholder="描述" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleSearch">查询</el-button>
          <el-button icon="Refresh" @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="addDialog">新增</el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column prop="id" label="ID" />
          <el-table-column prop="path" label="路径" />
          <el-table-column prop="group" label="分组" />
          <el-table-column prop="method" label="请求方法" />
          <el-table-column prop="description" label="描述" />
          <el-table-column label="操作">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
              <el-button
                type="danger"
                text
                icon="Delete"
                size="small"
                @click="deleteRoleAction(scope.row)"
                :disabled="scope.row.roleName === 'root'"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pager-wrapper">
        <el-pagination
          background
          :layout="paginationData.layout"
          :page-sizes="paginationData.pageSizes"
          :total="paginationData.total"
          :page-size="paginationData.pageSize"
          :currentPage="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"
import { usePagination } from "@/hooks/usePagination"
import { type ApiData, getApis } from "@/api/system/api"

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)
const searchFormData = reactive({
  path: "",
  apiGroup: "",
  method: "",
  description: ""
})

const methodOptions = [
  { value: "GET", label: "GET" },
  { value: "POST", label: "POST" },
  { value: "PUT", label: "PUT" },
  { value: "DELETE", label: "DELETE" }
]

const handleSearch = () => {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  getTableData()
}

const resetSearch = () => {}

const addDialog = () => {}

const tableData = ref<ApiData[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getApis({
      path: searchFormData.path || undefined,
      apiGroup: searchFormData.apiGroup || undefined,
      method: searchFormData.method || undefined,
      description: searchFormData.description || undefined,
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
    })
    if (res.code === 0) {
      tableData.value = res.data.list
      paginationData.total = res.data.total
    }
  } catch (error) {
    //
  }
  loading.value = false
}
getTableData()

// 分页
const handleSizeChange = (value: number) => {
  changePageSize(value)
  getTableData()
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  getTableData()
}
</script>
