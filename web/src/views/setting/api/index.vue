<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form ref="searchFormRef" :inline="true" :model="searchFormData">
        <el-form-item prop="path" label="路径">
          <el-input v-model="searchFormData.path" placeholder="路径" />
        </el-form-item>
        <el-form-item prop="group" label="API组">
          <el-input v-model="searchFormData.group" placeholder="API组" />
        </el-form-item>
        <el-form-item prop="method" label="方法">
          <el-input v-model="searchFormData.method" placeholder="方法" />
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
          <el-table-column prop="description" label="描述" />
          <el-table-column prop="method" label="请求方法" />
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
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"

const loading = ref(false)
const searchFormData = reactive({
  path: "",
  group: "",
  method: "",
  description: ""
})

const handleSearch = () => {}

const resetSearch = () => {}

const addDialog = () => {}

const tableData = ref([])

const getTableData = () => {}
</script>
