<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="CirclePlus">新增</el-button>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column prop="ID" label="ID" />
          <el-table-column prop="roleName" label="名称" align="center" />
          <el-table-column fixed="right" label="操作" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Setting" size="small">设置权限</el-button>
              <el-button type="primary" text icon="Edit" size="small">编辑</el-button>
              <el-button type="danger" text icon="Delete" size="small" :disabled="scope.row.roleName === 'root'"
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
import { ref } from "vue"
import { roleData, getRoles } from "@/api/system/role"

const loading = ref<boolean>(false)
const tableData = ref<roleData[]>([])

const getTableData = async () => {
  const res = await getRoles()
  if (res.code === 0) {
    tableData.value = res.data
  }
}
getTableData()
</script>

<style lang="scss" scoped>
.toolbar-wrapper {
  margin-bottom: 20px;
}
.table-wrapper {
  margin-bottom: 20px;
}
</style>
