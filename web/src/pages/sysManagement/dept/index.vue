<template>
  <div class="app-container">
    <!-- Search Form -->
    <el-card shadow="never" class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="部门名称">
          <el-input v-model="searchForm.deptName" placeholder="请输入部门名称" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="部门状态" clearable style="width: 120px">
            <el-option label="正常" :value="true" />
            <el-option label="禁用" :value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button>
          <el-button :icon="Refresh" @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Toolbar -->
    <el-card shadow="never" class="toolbar-card">
      <div class="toolbar-left">
        <el-button type="primary" :icon="Plus" @click="handleCreate">新增</el-button>
        <el-button type="warning" :icon="Check" @click="handleSaveSort">保存排序</el-button>
        <el-button :icon="Expand" @click="handleExpandAll">展开/折叠</el-button>
      </div>
    </el-card>

    <!-- Table -->
    <el-card shadow="never" class="table-card">
      <el-table
        ref="tableRef"
        v-loading="loading"
        :data="deptTree"
        row-key="id"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        border
        :default-expand-all="isExpandAll"
        highlight-current-row
      >
        <el-table-column prop="deptName" label="部门名称" min-width="250" show-overflow-tooltip>
          <template #default="{ row }">
            <span class="dept-name">{{ row.deptName }}</span>
          </template>
        </el-table-column>
        
        <el-table-column prop="sort" label="排序" width="120" align="center">
          <template #default="{ row }">
            <el-input-number 
              v-model="row.sort" 
              :min="0" 
              :max="999"
              controls-position="right"
              style="width: 90px"
              size="small"
            />
          </template>
        </el-table-column>
        
        <el-table-column prop="status" label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status ? 'success' : 'danger'">
              {{ row.status ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column prop="createdAt" label="创建时间" width="180" align="center">
          <template #default="{ row }">
            {{ formatDate(row.createdAt) }}
          </template>
        </el-table-column>
        
        <el-table-column label="操作" width="280" align="center" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link :icon="Edit" @click="handleUpdate(row)">修改</el-button>
            <el-button type="primary" link :icon="Plus" @click="handleAddChild(row)">新增</el-button>
            <el-button 
              v-if="!row.children || row.children.length === 0" 
              type="danger" 
              link 
              :icon="Delete" 
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      @closed="resetForm"
    >
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="上级部门">
          <el-tree-select
            v-model="formData.parentId"
            :data="deptOptions"
            :props="{ label: 'deptName', value: 'id', children: 'children' }"
            placeholder="请选择上级部门"
            clearable
            check-strictly
            :disabled="isEdit && formData.id === 1"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="部门名称" prop="deptName">
          <el-input v-model="formData.deptName" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="formData.sort" :min="0" :max="999" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="formData.status">
            <el-radio :label="true">正常</el-radio>
            <el-radio :label="false">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import type { FormInstance, FormRules, TableInstance } from "element-plus"
import { Plus, Edit, Delete, Search, Refresh, Check, Expand } from "@element-plus/icons-vue"
import {
  deptListApi,
  createDeptApi,
  updateDeptApi,
  deleteDeptApi,
  getElTreeDeptsApi,
  type Dept,
  type CreateDeptReq,
  type UpdateDeptReq
} from "@/api/sysManagement/dept"

// Search form
const searchForm = reactive({
  deptName: "",
  status: undefined as boolean | undefined
})

// Table data
const loading = ref(false)
const deptTree = ref<Dept[]>([])
const allDepts = ref<Dept[]>([])
const isExpandAll = ref(true)
const tableRef = ref<TableInstance>()

// Dialog
const dialogVisible = ref(false)
const dialogTitle = ref("")
const isEdit = ref(false)
const formRef = ref<FormInstance>()
const deptOptions = ref<Dept[]>([])

const formData = reactive<{
  id?: number
  deptName: string
  parentId: number | undefined
  sort: number
  status: boolean
}>({
  deptName: "",
  parentId: undefined,
  sort: 0,
  status: true
})

const formRules: FormRules = {
  deptName: [{ required: true, message: "请输入部门名称", trigger: "blur" }]
}

// Get department list
const getDeptList = async () => {
  loading.value = true
  try {
    const res = await deptListApi({
      deptName: searchForm.deptName,
      status: searchForm.status
    })
    if (res.code === 200) {
      deptTree.value = res.data
      allDepts.value = flattenTree(res.data)
    }
  } finally {
    loading.value = false
  }
}

// Flatten tree for options
const flattenTree = (tree: Dept[]): Dept[] => {
  const result: Dept[] = []
  const traverse = (nodes: Dept[]) => {
    nodes.forEach(node => {
      result.push(node)
      if (node.children) {
        traverse(node.children)
      }
    })
  }
  traverse(tree)
  return result
}

// Get department options for tree-select
const getDeptOptions = async () => {
  try {
    const res = await getElTreeDeptsApi()
    if (res.code === 200) {
      deptOptions.value = res.data.tree
    }
  } catch (error) {
    console.error(error)
  }
}

// Search
const handleSearch = () => {
  getDeptList()
}

// Reset
const handleReset = () => {
  searchForm.deptName = ""
  searchForm.status = undefined
  getDeptList()
}

// Expand/Collapse all
const handleExpandAll = () => {
  isExpandAll.value = !isExpandAll.value
  // Note: Element Plus table doesn't have a direct method to expand/collapse all
  // This is a simplified implementation
}

// Create
const handleCreate = () => {
  isEdit.value = false
  dialogTitle.value = "新增部门"
  resetForm()
  getDeptOptions()
  dialogVisible.value = true
}

// Update
const handleUpdate = (row: Dept) => {
  isEdit.value = true
  dialogTitle.value = "修改部门"
  Object.assign(formData, {
    id: row.id,
    deptName: row.deptName,
    parentId: row.parentId || undefined,
    sort: row.sort,
    status: row.status
  })
  getDeptOptions()
  dialogVisible.value = true
}

// Add child
const handleAddChild = (row: Dept) => {
  isEdit.value = false
  dialogTitle.value = "新增子部门"
  resetForm()
  formData.parentId = row.id
  getDeptOptions()
  dialogVisible.value = true
}

// Delete
const handleDelete = (row: Dept) => {
  ElMessageBox.confirm(`确定要删除部门 "${row.deptName}" 吗？`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    try {
      const res = await deleteDeptApi({ id: row.id! })
      if (res.code === 200) {
        ElMessage.success("删除成功")
        getDeptList()
      }
    } catch (error) {
      console.error(error)
    }
  })
}

// Submit form
const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          const res = await updateDeptApi(formData as UpdateDeptReq)
          if (res.code === 200) {
            ElMessage.success("更新成功")
            dialogVisible.value = false
            getDeptList()
          }
        } else {
          const res = await createDeptApi(formData as CreateDeptReq)
          if (res.code === 200) {
            ElMessage.success("创建成功")
            dialogVisible.value = false
            getDeptList()
          }
        }
      } catch (error) {
        console.error(error)
      }
    }
  })
}

// Save sort
const handleSaveSort = async () => {
  // Collect all sort changes
  const updates: { id: number; sort: number }[] = []
  const collectSort = (nodes: Dept[]) => {
    nodes.forEach(node => {
      updates.push({ id: node.id!, sort: node.sort })
      if (node.children) {
        collectSort(node.children)
      }
    })
  }
  collectSort(deptTree.value)
  
  // Batch update sort
  try {
    // Note: You may need to add a batch update API
    for (const item of updates) {
      await updateDeptApi({
        id: item.id,
        deptName: allDepts.value.find(d => d.id === item.id)?.deptName || "",
        parentId: allDepts.value.find(d => d.id === item.id)?.parentId || 0,
        sort: item.sort,
        status: allDepts.value.find(d => d.id === item.id)?.status || true
      })
    }
    ElMessage.success("排序保存成功")
    getDeptList()
  } catch (error) {
    console.error(error)
  }
}

// Reset form
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  Object.assign(formData, {
    deptName: "",
    parentId: undefined,
    sort: 0,
    status: true
  })
}

// Format date
const formatDate = (date: string | undefined) => {
  if (!date) return "-"
  return new Date(date).toLocaleString()
}

onMounted(() => {
  getDeptList()
})
</script>

<style lang="scss" scoped>
.app-container {
  padding: 20px;
}

.search-card {
  margin-bottom: 15px;
}

.toolbar-card {
  margin-bottom: 15px;
  
  .toolbar-left {
    display: flex;
    gap: 10px;
  }
}

.table-card {
  .dept-name {
    font-weight: 500;
  }
}

:deep(.el-table__row) {
  .el-input-number {
    .el-input__inner {
      text-align: center;
    }
  }
}
</style>
