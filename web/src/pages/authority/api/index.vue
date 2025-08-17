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
          <el-select v-model="searchFormData.method" placeholder="方法" clearable style="width: 100px">
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
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px">
              <el-button text @click="deleteVisible = false">取消</el-button>
              <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button
                icon="delete"
                type="danger"
                plain
                :disabled="!ids.length"
                style="margin-left: 10px"
                @click="deleteVisible = true"
                >删除</el-button
              >
            </template>
          </el-popover>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" @sort-change="handleSortChange" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="60" />
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="path" label="路径" sortable="custom" />
          <el-table-column prop="apiGroup" label="分组" sortable="custom" />
          <el-table-column prop="method" label="请求方法" sortable="custom" />
          <el-table-column prop="description" label="描述" />
          <el-table-column label="操作">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
              <el-button type="danger" text icon="Delete" size="small" @click="handleDeleteApi(scope.row)"
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" :before-close="handleClose" width="38%">
      <warning-bar title="新增接口，需要在角色管理内配置权限才可使用" />
      <el-form ref="formRef" :model="opFormData" :rules="addFormRules" label-width="80px">
        <el-form-item label="API路径" prop="path">
          <el-input v-model="opFormData.path" />
        </el-form-item>
        <el-form-item label="请求方法" prop="method">
          <el-select v-model="opFormData.method" placeholder="请选择方法" :clearable="true" style="width: 100%">
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="API分组" prop="apiGroup">
          <el-input v-model="opFormData.apiGroup" />
        </el-form-item>
        <el-form-item label="API描述" prop="description">
          <el-input v-model="opFormData.description" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="operateAction(formRef)">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox, ElNotification } from "element-plus"
import { usePagination } from "@/hooks/usePagination"
import {
  type ApiDataModel,
  getApisApi,
  addApiApi,
  deleteApiApi,
  deleteApiByIdApi,
  editApiApi
} from "@/api/authority/api"
import WarningBar from "@/components/WarningBar/warningBar.vue"

defineOptions({
  name: "Api"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)
const searchFormData = reactive({
  path: "",
  apiGroup: "",
  method: "",
  description: "",
  orderKey: "",
  // 默认升序
  desc: false
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

const resetSearch = () => {
  searchFormData.path = ""
  searchFormData.apiGroup = ""
  searchFormData.method = ""
  searchFormData.description = ""
  searchFormData.orderKey = ""
  searchFormData.desc = false
}

const tableData = ref<ApiDataModel[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getApisApi({
      path: searchFormData.path || undefined,
      apiGroup: searchFormData.apiGroup || undefined,
      method: searchFormData.method || undefined,
      description: searchFormData.description || undefined,
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

const ids = ref<number[]>([])
const taskIds = ref("") // for csv
const handleSelectionChange = (val: ApiDataModel[]) => {
  ids.value = val.map((item) => item.id)
  taskIds.value = ids.value.join(",")
}

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

// 对话框
const formRef = ref<FormInstance>()
const opFormData = reactive({
  path: "",
  apiGroup: "",
  method: "",
  description: ""
})

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}

let oKind: operationKind
const addFormRules: FormRules = reactive({
  path: [{ required: true, trigger: "blur", message: "路径不能为空" }],
  apiGroup: [{ required: true, trigger: "blur", message: "分组不能为空" }],
  method: [{ required: true, trigger: "change", message: "方法不能为空" }],
  description: [{ required: true, trigger: "blur", message: "描述不能为空" }]
})

const initForm = () => {
  formRef.value?.resetFields()
  opFormData.path = ""
  opFormData.apiGroup = ""
  opFormData.method = ""
  opFormData.description = ""
}

const dialogVisible = ref(false)
const dialogTitle = ref("")
const handleClose = (done: Function) => {
  initForm()
  done()
}

const addDialog = () => {
  dialogTitle.value = "新增接口"
  oKind = operationKind.Add
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
  initForm()
}

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      if (oKind === "Add") {
        const res = await addApiApi({ ...opFormData })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          tableData.value.push(res.data)
        }
      } else if (oKind === "Edit") {
        const res = await editApiApi({ id: activeRow.id, ...opFormData })
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          // 修改对应数据
          const index = tableData.value.indexOf(activeRow)
          tableData.value[index].apiGroup = opFormData.apiGroup
          tableData.value[index].path = opFormData.path
          tableData.value[index].description = opFormData.description
          tableData.value[index].method = opFormData.method
        }
      }
      closeDialog()
    }
  })
}

// 删除api
const handleDeleteApi = (row: ApiDataModel) => {
  ElMessageBox.confirm("此操作将永久删除所有角色下该api, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteApiApi({ id: row.id }).then((res) => {
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg })
          const index = tableData.value.indexOf(row)
          tableData.value.splice(index, 1)
        }
      })
    })
    .catch(() => {})
}

// 编辑dialog
let activeRow: ApiDataModel
const editDialog = (row: ApiDataModel) => {
  dialogTitle.value = "编辑接口"
  oKind = operationKind.Edit
  opFormData.apiGroup = row.apiGroup
  opFormData.description = row.description
  opFormData.method = row.method
  opFormData.path = row.path
  activeRow = row
  dialogVisible.value = true
}

const deleteVisible = ref(false)
const onDelete = async () => {
  if (ids.value.length === 0) {
    ElNotification({
      title: "警告",
      message: "请选择记录",
      type: "warning"
    })
    return
  }
  const res = await deleteApiByIdApi({ ids: ids.value })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: res.msg
    })
    deleteVisible.value = false
    getTableData()
  }
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
