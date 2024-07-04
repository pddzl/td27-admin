<template>
  <div class="app-container">
    <el-card shadow="never" class="search-wrapper">
      <el-form ref="searchFormRef" :inline="true" :model="searchFormData">
        <el-form-item prop="name" label="名称">
          <el-input v-model="searchFormData.name" placeholder="名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-upload
          :action="`${path}/file/upload`"
          :before-upload="checkFile"
          :headers="{ 'x-token': useUserStoreHook().token }"
          :on-error="uploadError"
          :on-success="uploadSuccess"
          :show-file-list="false"
          class="upload-btn"
        >
          <el-button type="primary" plain icon="upload">上传</el-button>
        </el-upload>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" @sort-change="handleSortChange">
          <el-table-column prop="id" label="ID" sortable="custom" />
          <el-table-column prop="fileName" label="名称" min-width="200" />
          <!-- <el-table-column prop="fullPath" label="路径" /> -->
          <el-table-column prop="mime" label="MIME" />
          <el-table-column prop="createdAt" label="创建时间">
            <template #default="scope">
              {{ formatDateTime(scope.row.createdAt) }}
            </template>
          </el-table-column>
          <el-table-column label="操作" width="180px" fixed="right" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Download" size="small" @click="handleDownload(scope.row.fileName)"
                >下载</el-button
              >
              <el-button type="danger" text icon="Delete" size="small" @click="handleDelete(scope.row)">删除</el-button>
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
import { ElMessage, ElMessageBox, type UploadProps } from "element-plus"
import { useUserStoreHook } from "@/store/modules/user"
import { usePagination } from "@/hooks/usePagination"
import { type fileDataModel, getFileListApi, downloadApi, deleteApi } from "@/api/fileM/file"
import { formatDateTime } from "@/utils/index"

defineOptions({
  name: "File"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const path = ref(import.meta.env.VITE_BASE_API)

const uploadSuccess = (res: ApiResponseData<fileDataModel>) => {
  if (res.code === 0) {
    tableData.value.push(res.data)
    ElMessage.success("上传成功")
  } else {
    ElMessage.error(res.msg)
  }
}

const checkFile: UploadProps["beforeUpload"] = (rawFile) => {
  if (rawFile.type !== "text/csv") {
    ElMessage.error("文件格式必须为csv")
    return false
  }
  return true
}

const uploadError = () => {
  ElMessage({
    type: "error",
    message: "上传失败"
  })
}

const handleSearch = () => {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  getTableData()
}

const searchFormData = reactive({
  name: "",
  orderKey: "",
  // 默认升序
  desc: false
})

// 排序
const handleSortChange = (column: any) => {
  searchFormData.orderKey = column.prop
  if (column.order === "descending") {
    searchFormData.desc = true
  } else {
    searchFormData.desc = false
  }
  getTableData()
}

const loading = ref(false)
const tableData = ref<fileDataModel[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getFileListApi({
      name: searchFormData.name || undefined,
      orderKey: searchFormData.orderKey || undefined,
      desc: searchFormData.desc || undefined,
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

// 下载文件
const handleDownload = (fileName: string) => {
  downloadApi({ name: fileName }).then((res) => {
    const blobData = res as Blob
    const url = window.URL.createObjectURL(new Blob([blobData], { type: "text/csv" }))
    const a = document.createElement("a")
    a.style.display = "none"
    a.href = url
    a.setAttribute("download", `${fileName}`)
    document.body.appendChild(a)
    a.click()
    window.URL.revokeObjectURL(a.href)
    document.body.removeChild(a)
  })
}

// 删除文件
const handleDelete = (row: fileDataModel) => {
  ElMessageBox.confirm("此操作将永久删除该文件, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(async () => {
      const res = await deleteApi({ name: row.fileName })
      if (res.code === 0) {
        ElMessage.success(res.msg)
        const index = tableData.value.indexOf(row)
        tableData.value.splice(index, 1)
      } else {
        ElMessage.error(res.msg)
      }
    })
    .catch(() => {})
}
</script>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 5px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
