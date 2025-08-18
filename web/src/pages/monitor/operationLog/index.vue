<script lang="ts" setup>
import type { orDataModel } from "@/api/monitor/operationLog"
import { usePagination } from "@@/composables/usePagination_n"
import { formatDateTime } from "@@/utils/datetime"
import { reactive, ref } from "vue"
import VueJsonPretty from "vue-json-pretty"
import { deleteOrApi, deleteOrByIdsApi, getOrListApi } from "@/api/monitor/operationLog"
import "vue-json-pretty/lib/styles.css"

defineOptions({
  name: "OperationRecord"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)
const searchFormData = reactive({
  path: "",
  method: "",
  status: 0
})

const methodOptions = [
  { value: "GET", label: "GET" },
  { value: "POST", label: "POST" },
  { value: "PUT", label: "PUT" },
  { value: "DELETE", label: "DELETE" }
]

function handleSearch() {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  getTableData()
}

function resetSearch() {
  searchFormData.path = ""
  searchFormData.method = ""
  searchFormData.status = 0
}

const tableData = ref<orDataModel[]>([])

async function getTableData() {
  loading.value = true
  try {
    const res = await getOrListApi({
      path: searchFormData.path || undefined,
      method: searchFormData.method || undefined,
      status: searchFormData.status || undefined,
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
    })
    if (res.code === 0) {
      tableData.value = res.data.list
      paginationData.total = res.data.total
    }
  } catch (error) {
    console.log(error)
  }
  loading.value = false
}
getTableData()

// 分页
function handleSizeChange(value: number) {
  changePageSize(value)
  getTableData()
}

function handleCurrentChange(value: number) {
  changeCurrentPage(value)
  getTableData()
}

// 删除操作记录
function deleteOrFunc(row: orDataModel) {
  deleteOrApi({ id: row.id }).then((res) => {
    if (res.code === 0) {
      ElMessage({ type: "success", message: res.msg })
      const index = tableData.value.indexOf(row)
      tableData.value.splice(index, 1)
    }
  })
}

// 批量删除
const multipleSelection = ref<orDataModel[]>([])
function handleSelectionChange(val: orDataModel[]) {
  multipleSelection.value = val
}

async function deleteByIdsFunc() {
  const ids: number[] = []
  multipleSelection.value
  && multipleSelection.value.forEach((item) => {
    ids.push(item.id)
  })
  if (ids.length === 0) {
    ElNotification({
      title: "警告",
      message: "请选择记录",
      type: "warning"
    })
    return
  }
  const res = await deleteOrByIdsApi({ ids })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功"
    })
    getTableData()
  }
}

function typeFilter(effect: number) {
  const structure: Record<string, "success" | "info" | "warning" | "danger"> = {
    2: "success",
    3: "success",
    4: "warning",
    5: "danger"
  }
  const key = String(effect)[0]
  return structure[key] || "info"
}
</script>

<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form :inline="true" :model="searchFormData">
        <el-form-item prop="path" label="路径">
          <el-input v-model="searchFormData.path" placeholder="路径" />
        </el-form-item>
        <el-form-item prop="method" label="方法">
          <el-select v-model="searchFormData.method" placeholder="方法" clearable style="width: 100px">
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item prop="group" label="状态值">
          <el-input-number
            v-model="searchFormData.status"
            :min="0"
            :max="600"
            controls-position="right"
            placeholder="状态值"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleSearch">
            查询
          </el-button>
          <el-button icon="Refresh" @click="resetSearch">
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-popover placement="top" width="160">
          <p>确定要删除此记录吗</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link>
              取消
            </el-button>
            <el-button type="primary" @click="deleteByIdsFunc">
              确定
            </el-button>
          </div>
          <template #reference>
            <el-button type="danger" plain icon="delete" :disabled="!multipleSelection.length">
              删除
            </el-button>
          </template>
        </el-popover>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" style="width: 100%" row-key="id" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="40" />
          <el-table-column type="expand">
            <template #default="props">
              <div style="padding: 0px 18px">
                <el-tabs>
                  <el-tab-pane label="请求信息">
                    <VueJsonPretty
                      :data="props.row.reqParam !== '{}' ? JSON.parse(props.row.reqParam) : null"
                      :show-line="true"
                    />
                  </el-tab-pane>
                  <el-tab-pane label="响应信息">
                    <VueJsonPretty :data="JSON.parse(props.row.respData)" :show-line="true" />
                  </el-tab-pane>
                </el-tabs>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="id" label="ID" width="60" />
          <el-table-column prop="userName" label="用户" />
          <el-table-column prop="ip" label="IP" width="120" />
          <!-- <el-table-column prop="userAgent" label="UserAgent" min-width="140" /> -->
          <el-table-column prop="path" label="路径" min-width="150" />
          <el-table-column prop="status" label="状态码">
            <template #default="scope">
              <el-tag :type="typeFilter(scope.row.status)">
                {{ scope.row.status }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="method" label="请求方法" />
          <el-table-column prop="respTime" label="响应时间(ms)" min-width="120" />
          <el-table-column prop="createdAt" label="创建时间" min-width="180">
            <template #default="scope">
              {{ formatDateTime(scope.row.createdAt) }}
            </template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" min-width="100">
            <template #default="scope">
              <el-popover placement="top" width="160">
                <p>确定要删除此记录吗</p>
                <div style="text-align: right; margin-top: 8px">
                  <el-button type="primary" link>
                    取消
                  </el-button>
                  <el-button type="primary" @click="deleteOrFunc(scope.row)">
                    确定
                  </el-button>
                </div>
                <template #reference>
                  <el-button type="danger" size="small" link icon="Delete">
                    删除
                  </el-button>
                </template>
              </el-popover>
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
          :current-page="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 5px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
