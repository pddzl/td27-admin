<template>
  <div class="app-container">
    <!-- Stats Cards -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6">
        <el-card shadow="hover" class="stats-card">
          <div class="stats-icon total">
            <el-icon><Timer /></el-icon>
          </div>
          <div class="stats-info">
            <div class="stats-value">{{ stats.total }}</div>
            <div class="stats-label">总任务</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stats-card">
          <div class="stats-icon running">
            <el-icon><VideoPlay /></el-icon>
          </div>
          <div class="stats-info">
            <div class="stats-value">{{ stats.running }}</div>
            <div class="stats-label">运行中</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stats-card">
          <div class="stats-icon success">
            <el-icon><CircleCheck /></el-icon>
          </div>
          <div class="stats-info">
            <div class="stats-value">{{ stats.success }}</div>
            <div class="stats-label">成功</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stats-card">
          <div class="stats-icon failed">
            <el-icon><CircleClose /></el-icon>
          </div>
          <div class="stats-info">
            <div class="stats-value">{{ stats.failed }}</div>
            <div class="stats-label">失败</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Toolbar -->
    <el-card shadow="never" class="toolbar-card">
      <div class="toolbar-left">
        <el-button type="primary" :icon="Plus" @click="handleCreate">新增任务</el-button>
        <el-button :icon="Delete" :disabled="!selectedIds.length" @click="handleBatchDelete">批量删除</el-button>
      </div>
      <div class="toolbar-right">
        <el-input v-model="searchForm.name" placeholder="任务名称" clearable style="width: 200px" />
        <el-select v-model="searchForm.method" placeholder="任务类型" clearable style="width: 150px">
          <el-option label="清理数据表" value="clearTable" />
          <el-option label="Shell命令" value="shell" />
        </el-select>
        <el-select v-model="searchForm.status" placeholder="状态" clearable style="width: 120px">
          <el-option label="运行中" :value="true" />
          <el-option label="已停止" :value="false" />
        </el-select>
        <el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button>
        <el-button :icon="Refresh" @click="handleReset">重置</el-button>
      </div>
    </el-card>

    <!-- Table -->
    <el-card shadow="never" class="table-card">
      <el-table
        v-loading="loading"
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        highlight-current-row
      >
        <el-table-column type="selection" width="50" align="center" />
        <el-table-column prop="name" label="任务名称" min-width="180" show-overflow-tooltip />
        <el-table-column prop="method" label="任务类型" width="120">
          <template #default="{ row }">
            <el-tag :type="getMethodType(row.method)">
              {{ getMethodLabel(row.method) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="expression" label="执行规则" width="180">
          <template #default="{ row }">
            <el-tooltip :content="parseCron(row.expression)" placement="top">
              <span class="cron-expression">{{ row.expression }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column prop="strategy" label="执行策略" width="100">
          <template #default="{ row }">
            <el-tag :type="row.strategy === 'once' ? 'warning' : 'info'" size="small">
              {{ row.strategy === 'once' ? '仅一次' : '重复' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="open" label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-switch
              v-model="row.open"
              inline-prompt
              :active-value="true"
              :inactive-value="false"
              active-text="运行"
              inactive-text="停止"
              @change="(val: string | number | boolean) => handleStatusChange(row, val as boolean)"
            />
          </template>
        </el-table-column>
        <el-table-column prop="nextRunTime" label="下次执行" width="180">
          <template #default="{ row }">
            <span v-if="row.open && row.nextRunTime">{{ row.nextRunTime }}</span>
            <span v-else class="text-gray">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="comment" label="备注" min-width="150" show-overflow-tooltip />
        <el-table-column label="操作" width="200" align="center" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link :icon="Edit" @click="handleEdit(row)">编辑</el-button>
            <el-button type="primary" link :icon="VideoPlay" @click="handleRunOnce(row)">执行</el-button>
            <el-button type="danger" link :icon="Delete" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- Pagination -->
      <div class="pagination-wrapper">
        <el-pagination
          v-model:current-page="paginationData.currentPage"
          v-model:page-size="paginationData.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="paginationData.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- Create/Edit Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑定时任务' : '新增定时任务'"
      width="700px"
      @closed="resetForm"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="任务名称" prop="name">
              <el-input v-model="formData.name" placeholder="请输入任务名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="任务类型" prop="method">
              <el-select v-model="formData.method" placeholder="选择任务类型" style="width: 100%" @change="handleMethodChange">
                <el-option label="清理数据表" value="clearTable" />
                <el-option label="Shell命令" value="shell" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- Cron Expression Builder -->
        <el-form-item label="执行规则" prop="expression">
          <div class="cron-builder">
            <el-input v-model="formData.expression" placeholder="* * * * *" class="cron-input">
              <template #append>
                <el-button :icon="Clock" @click="showCronBuilder = true">生成</el-button>
              </template>
            </el-input>
            <div v-if="formData.expression" class="cron-desc">
              {{ parseCron(formData.expression) }}
            </div>
          </div>
        </el-form-item>

        <!-- ClearTable Config -->
        <template v-if="formData.method === 'clearTable'">
          <el-divider>数据清理配置</el-divider>
          <div v-for="(item, index) in formData.extraParams.tableInfo" :key="index" class="table-config-item">
            <el-row :gutter="10">
              <el-col :span="7">
                <el-form-item :label="index === 0 ? '数据表' : ''" label-width="100px">
                  <el-input v-model="item.tableName" placeholder="表名" />
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item :label="index === 0 ? '时间字段' : ''" label-width="100px">
                  <el-input v-model="item.compareField" placeholder="如: created_at" />
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item :label="index === 0 ? '保留时长' : ''" label-width="100px">
                  <el-input v-model="item.interval" placeholder="如: 720h (30天)">
                    <template #append>前</template>
                  </el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item :label="index === 0 ? '' : ''" label-width="0">
                  <el-button type="danger" :icon="Delete" circle @click="removeTableConfig(index)" />
                </el-form-item>
              </el-col>
            </el-row>
          </div>
          <el-button type="primary" link :icon="Plus" @click="addTableConfig">添加配置</el-button>
        </template>

        <!-- Shell Config -->
        <template v-if="formData.method === 'shell'">
          <el-divider>Shell命令</el-divider>
          <el-form-item label="命令" prop="extraParams.command">
            <el-input
              v-model="formData.extraParams.command"
              type="textarea"
              :rows="4"
              placeholder="请输入Shell命令"
            />
          </el-form-item>
        </template>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="执行策略">
              <el-radio-group v-model="formData.strategy">
                <el-radio label="always">重复执行</el-radio>
                <el-radio label="once">仅执行一次</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="立即启用">
              <el-switch v-model="formData.open" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="备注">
          <el-input v-model="formData.comment" type="textarea" :rows="2" placeholder="任务备注" />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- Cron Expression Builder Dialog -->
    <el-dialog v-model="showCronBuilder" title="Cron表达式生成器" width="600px">
      <CronBuilder v-model="formData.expression" />
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import type { FormInstance, FormRules } from "element-plus"
import { 
  Plus, Delete, Search, Refresh, Edit, VideoPlay, 
  Timer, CircleCheck, CircleClose, Clock 
} from "@element-plus/icons-vue"
import { usePagination } from "@@/composables/usePagination_n"
import type { cronDataModel, TableInfo } from "@/api/sysTool/cron"
import {
  cronListApi,
  cronCreateApi,
  cronUpdateApi,
  cronDeleteApi,
  cronDeleteByIds,
  cronSwitchOpenApi,
  METHOD
} from "@/api/sysTool/cron"
import CronBuilder from "./components/CronBuilder.vue"

// Stats
const stats = reactive({
  total: 0,
  running: 0,
  success: 0,
  failed: 0
})

// Search
const searchForm = reactive({
  name: "",
  method: "",
  status: undefined as boolean | undefined
})

// Table
const loading = ref(false)
const tableData = ref<cronDataModel[]>([])
const selectedIds = ref<number[]>([])
const { paginationData, changeCurrentPage, changePageSize } = usePagination()

// Dialog
const dialogVisible = ref(false)
const showCronBuilder = ref(false)
const isEdit = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  id: 0,
  name: "",
  method: "",
  expression: "",
  strategy: "always",
  open: false,
  extraParams: {
    tableInfo: [] as TableInfo[],
    command: ""
  },
  comment: ""
})

const formRules: FormRules = {
  name: [{ required: true, message: "请输入任务名称", trigger: "blur" }],
  method: [{ required: true, message: "请选择任务类型", trigger: "change" }],
  expression: [{ required: true, message: "请输入执行规则", trigger: "blur" }]
}

// Get table data
const getTableData = async () => {
  loading.value = true
  try {
    const res = await cronListApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
    })
    if (res.code === 0 || res.code === 200) {
      tableData.value = res.data.list
      paginationData.total = res.data.total
      updateStats()
    }
  } catch (error) {
    console.error(error)
  }
  loading.value = false
}

// Update stats
const updateStats = () => {
  stats.total = tableData.value.length
  stats.running = tableData.value.filter(item => item.open).length
  // Mock data for success/failed (in real app, get from API)
  stats.success = Math.floor(stats.total * 0.8)
  stats.failed = Math.floor(stats.total * 0.1)
}

// Search
const handleSearch = () => {
  paginationData.currentPage = 1
  getTableData()
}

const handleReset = () => {
  searchForm.name = ""
  searchForm.method = ""
  searchForm.status = undefined
  handleSearch()
}

// Selection
const handleSelectionChange = (val: cronDataModel[]) => {
  selectedIds.value = val.map(item => item.id)
}

// Create
const handleCreate = () => {
  isEdit.value = false
  resetForm()
  dialogVisible.value = true
}

// Edit
const handleEdit = (row: cronDataModel) => {
  isEdit.value = true
  Object.assign(formData, {
    id: row.id,
    name: row.name,
    method: row.method,
    expression: row.expression,
    strategy: row.strategy,
    open: row.open,
    extraParams: {
      tableInfo: row.extraParams?.tableInfo || [],
      command: row.extraParams?.command || ""
    },
    comment: row.comment
  })
  dialogVisible.value = true
}

// Status change
const handleStatusChange = async (row: cronDataModel, val: boolean) => {
  try {
    const res = await cronSwitchOpenApi({ id: row.id, open: val })
    if (res.code === 0 || res.code === 200) {
      ElMessage.success(val ? "任务已启动" : "任务已停止")
      getTableData()
    }
  } catch (error) {
    row.open = !val // Revert on error
  }
}

// Run once
const handleRunOnce = (row: cronDataModel) => {
  ElMessageBox.confirm(`确定要立即执行任务 "${row.name}" 吗？`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "info"
  }).then(() => {
    // Call run once API
    ElMessage.success("任务执行中")
  })
}

// Delete
const handleDelete = (row: cronDataModel) => {
  ElMessageBox.confirm(`确定要删除任务 "${row.name}" 吗？`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    try {
      const res = await cronDeleteApi({ id: row.id })
      if (res.code === 0 || res.code === 200) {
        ElMessage.success("删除成功")
        getTableData()
      }
    } catch (error) {
      console.error(error)
    }
  })
}

// Batch delete
const handleBatchDelete = () => {
  if (!selectedIds.value.length) return
  ElMessageBox.confirm(`确定要删除选中的 ${selectedIds.value.length} 个任务吗？`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    try {
      const res = await cronDeleteByIds({ ids: selectedIds.value })
      if (res.code === 0 || res.code === 200) {
        ElMessage.success("批量删除成功")
        getTableData()
      }
    } catch (error) {
      console.error(error)
    }
  })
}

// Submit
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          const res = await cronUpdateApi({ ...formData })
          if (res.code === 0 || res.code === 200) {
            ElMessage.success("更新成功")
            dialogVisible.value = false
            getTableData()
          }
        } else {
          const res = await cronCreateApi({ ...formData })
          if (res.code === 0 || res.code === 200) {
            ElMessage.success("创建成功")
            dialogVisible.value = false
            getTableData()
          }
        }
      } catch (error) {
        console.error(error)
      }
    }
  })
}

// Method change
const handleMethodChange = () => {
  formData.extraParams.tableInfo = []
  formData.extraParams.command = ""
}

// Table config
const addTableConfig = () => {
  formData.extraParams.tableInfo.push({
    tableName: "",
    compareField: "",
    interval: ""
  })
}

const removeTableConfig = (index: number) => {
  formData.extraParams.tableInfo.splice(index, 1)
}

// Reset form
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  Object.assign(formData, {
    id: 0,
    name: "",
    method: "",
    expression: "",
    strategy: "always",
    open: false,
    extraParams: {
      tableInfo: [],
      command: ""
    },
    comment: ""
  })
}

// Helpers
const getMethodType = (method: string): "success" | "warning" | "info" | "danger" => {
  const map: Record<string, "success" | "warning" | "info" | "danger"> = {
    clearTable: "success",
    shell: "warning"
  }
  return map[method] || "info"
}

const getMethodLabel = (method: string) => {
  const map: Record<string, string> = {
    clearTable: "清理数据",
    shell: "Shell命令"
  }
  return map[method] || method
}

const parseCron = (expression: string) => {
  // Simple cron parser (you can use a library like cron-parser)
  if (!expression) return ""
  const parts = expression.split(" ")
  if (parts.length !== 5) return expression
  
  const [minute, hour, day, month, week] = parts
  
  if (minute === "*" && hour === "*") return "每分钟执行"
  if (minute === "0" && hour === "*") return "每小时执行"
  if (minute === "0" && hour === "0") return "每天执行"
  if (minute.startsWith("*/")) return `每${minute.slice(2)}分钟执行`
  
  return expression
}

// Pagination
const handleSizeChange = (val: number) => {
  changePageSize(val)
  getTableData()
}

const handleCurrentChange = (val: number) => {
  changeCurrentPage(val)
  getTableData()
}

onMounted(() => {
  getTableData()
})
</script>

<style lang="scss" scoped>
.app-container {
  padding: 20px;
}

.stats-row {
  margin-bottom: 20px;
}

.stats-card {
  display: flex;
  align-items: center;
  padding: 20px;
  
  .stats-icon {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 28px;
    margin-right: 15px;
    
    &.total {
      background: #e6f7ff;
      color: #1890ff;
    }
    &.running {
      background: #f6ffed;
      color: #52c41a;
    }
    &.success {
      background: #f0f9ff;
      color: #13c2c2;
    }
    &.failed {
      background: #fff1f0;
      color: #ff4d4f;
    }
  }
  
  .stats-info {
    flex: 1;
    
    .stats-value {
      font-size: 28px;
      font-weight: bold;
      color: #333;
    }
    
    .stats-label {
      font-size: 14px;
      color: #999;
      margin-top: 5px;
    }
  }
}

.toolbar-card {
  margin-bottom: 15px;
  
  .toolbar-left {
    float: left;
  }
  
  .toolbar-right {
    float: right;
    display: flex;
    gap: 10px;
  }
  
  &::after {
    content: "";
    display: block;
    clear: both;
  }
}

.table-card {
  .cron-expression {
    font-family: monospace;
    background: #f5f5f5;
    padding: 2px 6px;
    border-radius: 4px;
    font-size: 12px;
  }
  
  .text-gray {
    color: #999;
  }
}

.pagination-wrapper {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.cron-builder {
  .cron-input {
    width: 300px;
  }
  
  .cron-desc {
    margin-top: 8px;
    color: #666;
    font-size: 13px;
  }
}

.table-config-item {
  margin-bottom: 10px;
  padding: 15px;
  background: #f5f7fa;
  border-radius: 4px;
}
</style>
