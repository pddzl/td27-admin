<script lang="ts" setup>
import type { ElTree } from "element-plus"
import type { ApiChild, ApiTreeData } from "@/api/sysManagement/api"
import type { ServiceToken } from "@/api/sysTool/service_token"
import { usePagination } from "@@/composables/usePagination_n"
import { formatDateTime } from "@@/utils/datetime"
import { apiGetElTreeApi } from "@/api/sysManagement/api"
import {
  createServiceTokenApi,
  deleteServiceTokenApi,
  getServiceTokenDetailApi,
  listServiceTokenApi,
  updateServiceTokenApi
} from "@/api/sysTool/service_token"

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

// Table data
const tableData = ref<ServiceToken[]>([])
const total = ref(0)
const loading = ref(false)
const searchForm = ref({
  name: "",
  status: undefined as boolean | undefined
})

// Dialog
const dialogVisible = ref(false)
const dialogTitle = ref("")
const isEdit = ref(false)
const saving = ref(false)

const apiIds = ref<number[]>([])

// Form
const form = ref<{
  id?: number
  name: string
  status?: boolean
  expiresAt?: number
  apiIds: number[] // API IDs (domain_id)
  apiKeys: string[] // Tree keys for checked APIs
}>({
  name: "",
  expiresAt: undefined,
  apiIds: [],
  apiKeys: [],
  status: true
})
const formRef = ref()

// API Tree
const apiTreeRef = ref<InstanceType<typeof ElTree>>()
const apiTreeData = ref<ApiTreeData[]>([])
const apiTreeLoading = ref(false)

// Generated token (only shown once after creation)
const generatedToken = ref("")
const showTokenDialog = ref(false)

// Rules
const rules = {
  name: [
    { required: true, message: "请输入令牌名称", trigger: "blur" },
    { max: 100, message: "长度不能超过100个字符", trigger: "blur" }
  ]
}

// Load table data
async function loadData() {
  loading.value = true
  try {
    const res = await listServiceTokenApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize,
      name: searchForm.value.name || undefined,
      status: searchForm.value.status
    })
    if (res.code === 0) {
      tableData.value = res.data.list
      total.value = res.data.total
    }
  } finally {
    loading.value = false
  }
}

// Load API tree
async function loadApiTree(id: number) {
  apiTreeLoading.value = true
  try {
    const res = await apiGetElTreeApi({ id, from_source: "token" })
    if (res.code === 0) {
      apiTreeData.value = res.data.list || []
      apiIds.value = res.data.checkedIds || []
      // console.log("API Tree loaded:", apiTreeData.value)
    } else {
      ElMessage.error(res.msg || "加载API列表失败")
    }
  } catch (err) {
    console.error("Load API tree error:", err)
    ElMessage.error("加载API列表失败")
  } finally {
    apiTreeLoading.value = false
  }
}

// Search
function handleSearch() {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  loadData()
}

// 分页
function handleSizeChange(value: number) {
  changePageSize(value)
  loadData()
}

function handleCurrentChange(value: number) {
  changeCurrentPage(value)
  loadData()
}

// Reset
function handleReset() {
  searchForm.value = { name: "", status: undefined }
  handleSearch()
}

// Add new
function handleAdd() {
  isEdit.value = false
  dialogTitle.value = "创建服务令牌"
  form.value = {
    name: "",
    expiresAt: undefined,
    apiIds: [],
    apiKeys: [],
    status: true
  }
  dialogVisible.value = true
  // loadApiTree()
}

// Edit
async function handleEdit(row: ServiceToken) {
  isEdit.value = true
  dialogTitle.value = "编辑服务令牌"
  dialogVisible.value = true
  loadApiTree(row.id)

  const res = await getServiceTokenDetailApi({ id: row.id })
  if (res.code === 0) {
    form.value.id = res.data.id
    form.value.name = res.data.name
    form.value.status = res.data.status
    form.value.expiresAt = res.data.expiresAt || undefined
    form.value.apiIds = res.data.apiIds || []
    // Tree uses string keys (API IDs), backend returns API IDs as numbers
    form.value.apiKeys = res.data.apiIds?.map(id => id.toString()) || []
  }
}

// Delete
function handleDelete(row: ServiceToken) {
  ElMessageBox.confirm(`确定要删除令牌 "${row.name}" 吗？删除后将无法恢复。`, "提示", {
    type: "warning"
  }).then(async () => {
    const res = await deleteServiceTokenApi(row.id)
    if (res.code === 0) {
      ElMessage.success("删除成功")
      loadData()
    }
  })
}

// Get checked API IDs from tree
function getCheckedApiIds(): number[] {
  const checkedNodes = apiTreeRef.value?.getCheckedNodes(false, true) as ApiChild[]
  if (!checkedNodes) return []

  return checkedNodes
    .filter(node => node.id && node.path && node.method)
    .map(node => node.id)
}

// Submit
async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  saving.value = true
  try {
    const apiIds = getCheckedApiIds()

    if (isEdit.value && form.value.id) {
      const res = await updateServiceTokenApi({
        id: form.value.id,
        name: form.value.name,
        status: form.value.status!,
        expiresAt: form.value.expiresAt ? Number(form.value.expiresAt) : undefined,
        apiIds
      })
      if (res.code === 0) {
        ElMessage.success("更新成功")
        dialogVisible.value = false
        loadData()
      }
    } else {
      const res = await createServiceTokenApi({
        name: form.value.name,
        expiresAt: form.value.expiresAt ? Number(form.value.expiresAt) : undefined,
        apiIds
      })
      if (res.code === 0) {
        generatedToken.value = res.data.token
        showTokenDialog.value = true
        dialogVisible.value = false
        loadData()
      }
    }
  } finally {
    saving.value = false
  }
}

// Copy token
function copyToken() {
  navigator.clipboard.writeText(generatedToken.value)
  ElMessage.success("已复制到剪贴板")
}

// Status change
async function handleStatusChange(row: ServiceToken) {
  const detailRes = await getServiceTokenDetailApi({ id: row.id })
  if (detailRes.code !== 0) return

  const res = await updateServiceTokenApi({
    id: row.id,
    name: row.name,
    status: row.status,
    expiresAt: row.expiresAt || undefined,
    apiIds: detailRes.data.apiIds
  })
  if (res.code === 0) {
    ElMessage.success("状态更新成功")
  } else {
    row.status = !row.status
  }
}

// Format expiration
function formatExpires(expiresAt: number | null) {
  if (!expiresAt) return "永不过期"
  return formatDateTime(expiresAt)
}

onMounted(() => {
  loadData()
})
</script>

<template>
  <div class="app-container">
    <el-card class="mb-[5px]" shadow="never">
      <el-form :model="searchForm" inline>
        <el-form-item label="令牌名称">
          <el-input v-model="searchForm.name" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="全部" clearable style="width: 120px">
            <el-option label="启用" :value="true" />
            <el-option label="禁用" :value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            搜索
          </el-button>
          <el-button @click="handleReset">
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never">
      <template #header>
        <div class="flex justify-between items-center">
          <el-button type="primary" icon="Plus" @click="handleAdd">
            创建令牌
          </el-button>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="loadData" />
          </el-tooltip>
        </div>
      </template>

      <el-table v-loading="loading" :data="tableData" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="令牌名称" min-width="120px" />
        <el-table-column label="状态" width="100px">
          <template #default="{ row }">
            <el-switch v-model="row.status" @change="handleStatusChange(row)" />
          </template>
        </el-table-column>
        <el-table-column label="权限数量" width="100px">
          <template #default="{ row }">
            <el-tag>{{ row.apiCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="过期时间" width="180">
          <template #default="{ row }">
            {{ formatExpires(row.expiresAt) }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">
              编辑
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="mt-5 flex justify-end">
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="700px" destroy-on-close>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="令牌名称" prop="name">
          <el-input v-model="form.name" placeholder="例如：第三方系统集成" />
        </el-form-item>

        <el-form-item label="过期时间">
          <el-date-picker
            v-model="form.expiresAt"
            type="datetime"
            placeholder="选择过期时间（留空永不过期）"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="X"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="API权限">
          <div class="api-tree-container">
            <el-tree
              v-if="!apiTreeLoading && apiTreeData.length > 0"
              ref="apiTreeRef"
              :data="apiTreeData"
              node-key="key"
              :default-checked-keys="apiIds"
              default-expand-all
              show-checkbox
              :props="{
                label: (data: any) => data.description || data.key,
                children: 'children',
              }"
            />
            <el-empty v-else-if="!apiTreeLoading && apiTreeData.length === 0" description="暂无API数据，请先创建API" />
            <el-skeleton v-else :rows="6" animated />
          </div>
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">
          取消
        </el-button>
        <el-button type="primary" :loading="saving" @click="handleSubmit">
          保存
        </el-button>
      </template>
    </el-dialog>

    <el-dialog
      v-model="showTokenDialog"
      title="服务令牌创建成功"
      width="600px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <el-alert
        title="请立即复制保存"
        description="此令牌仅显示一次，关闭后将无法再次查看。请妥善保管。"
        type="warning"
        :closable="false"
        show-icon
        style="margin-bottom: 20px"
      />

      <div class="token-display">
        <code>{{ generatedToken }}</code>
        <el-button type="primary" @click="copyToken">
          复制
        </el-button>
      </div>

      <div class="token-usage">
        <h4>使用说明</h4>
        <p>在请求头中添加以下字段：</p>
        <pre>X-Service-Token: {{ generatedToken }}</pre>
      </div>

      <template #footer>
        <el-button type="primary" @click="showTokenDialog = false">
          我已保存
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.api-tree-container {
  max-height: 400px;
  overflow-y: auto;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  padding: 10px;
}

.token-display {
  display: flex;
  gap: 10px;
  align-items: center;
  background: #f5f7fa;
  padding: 15px;
  border-radius: 4px;
  margin-bottom: 20px;

  code {
    flex: 1;
    font-size: 14px;
    word-break: break-all;
    font-family: "Courier New", monospace;
  }
}

.token-usage {
  background: #f5f7fa;
  padding: 15px;
  border-radius: 4px;

  h4 {
    margin-top: 0;
    margin-bottom: 10px;
  }

  pre {
    background: #2d2d2d;
    color: #f8f8f2;
    padding: 10px;
    border-radius: 4px;
    overflow-x: auto;
  }
}
</style>
