<script lang="ts" setup>
import { ref, onMounted } from "vue"
import {
  createButtonApi,
  updateButtonApi,
  deleteButtonApi,
  listButtonApi,
  type ButtonData
} from "@/api/sysManagement/button"

const tableData = ref<ButtonData[]>([])
const total = ref(0)
const loading = ref(false)
const searchForm = ref({
  pagePath: ""
})
const pagination = ref({
  page: 1,
  pageSize: 10
})

const dialogVisible = ref(false)
const dialogTitle = ref("")
const isEdit = ref(false)
const saving = ref(false)

const form = ref<Partial<ButtonData>>({
  buttonCode: "",
  buttonName: "",
  description: "",
  pagePath: ""
})
const formRef = ref()

const rules = {
  buttonCode: [
    { required: true, message: "请输入按钮代码", trigger: "blur" },
    { max: 100, message: "长度不能超过100", trigger: "blur" }
  ],
  buttonName: [
    { required: true, message: "请输入按钮名称", trigger: "blur" },
    { max: 100, message: "长度不能超过100", trigger: "blur" }
  ],
  pagePath: [
    { required: true, message: "请输入页面路径", trigger: "blur" },
    { max: 200, message: "长度不能超过200", trigger: "blur" }
  ]
}

async function loadData() {
  loading.value = true
  try {
    const res = await listButtonApi({
      page: pagination.value.page,
      pageSize: pagination.value.pageSize,
      pagePath: searchForm.value.pagePath || undefined
    })
    if (res.code === 0) {
      tableData.value = res.data.list
      total.value = res.data.total
    }
  } finally {
    loading.value = false
  }
}

function handleSearch() {
  pagination.value.page = 1
  loadData()
}

function handleReset() {
  searchForm.value.pagePath = ""
  handleSearch()
}

function handleAdd() {
  isEdit.value = false
  dialogTitle.value = "创建按钮"
  form.value = {
    buttonCode: "",
    buttonName: "",
    description: "",
    pagePath: ""
  }
  dialogVisible.value = true
}

function handleEdit(row: ButtonData) {
  isEdit.value = true
  dialogTitle.value = "编辑按钮"
  form.value = { ...row }
  dialogVisible.value = true
}

function handleDelete(row: ButtonData) {
  ElMessageBox.confirm(`确定删除 "${row.buttonName}" 吗？`, "提示", {
    type: "warning"
  }).then(async () => {
    const res = await deleteButtonApi(row.id)
    if (res.code === 0) {
      ElMessage.success("删除成功")
      loadData()
    }
  })
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  saving.value = true
  try {
    if (isEdit.value && form.value.id) {
      const res = await updateButtonApi(form.value as ButtonData)
      if (res.code === 0) {
        ElMessage.success("更新成功")
        dialogVisible.value = false
        loadData()
      }
    } else {
      const res = await createButtonApi(form.value as Omit<ButtonData, "id">)
      if (res.code === 0) {
        ElMessage.success("创建成功")
        dialogVisible.value = false
        loadData()
      }
    }
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<template>
  <div class="page-container">
    <el-card class="search-card" shadow="never">
      <el-form :model="searchForm" inline>
        <el-form-item label="页面路径">
          <el-input v-model="searchForm.pagePath" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="table-card" shadow="never">
      <template #header>
        <div class="card-header">
          <span>按钮权限列表</span>
          <el-button type="primary" @click="handleAdd">+ 创建按钮</el-button>
        </div>
      </template>

      <el-table v-loading="loading" :data="tableData" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="buttonCode" label="按钮代码" min-width="150" />
        <el-table-column prop="buttonName" label="按钮名称" min-width="120" />
        <el-table-column prop="pagePath" label="页面路径" min-width="150" />
        <el-table-column prop="description" label="描述" min-width="150" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next"
          @size-change="loadData"
          @current-change="loadData"
        />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px" destroy-on-close>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="按钮代码" prop="buttonCode">
          <el-input v-model="form.buttonCode" placeholder="如: user_create" />
        </el-form-item>
        <el-form-item label="按钮名称" prop="buttonName">
          <el-input v-model="form.buttonName" placeholder="如: 新建用户" />
        </el-form-item>
        <el-form-item label="页面路径" prop="pagePath">
          <el-input v-model="form.pagePath" placeholder="如: /sysManagement/user" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.description" type="textarea" :rows="2" />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saving" @click="handleSubmit">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.page-container {
  padding: 20px;
}
.search-card {
  margin-bottom: 20px;
}
.table-card {
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}
.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
