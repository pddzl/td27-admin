<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="CirclePlus" @click="addMenuAction(0)">新增</el-button>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" row-key="id">
          <el-table-column prop="id" label="ID" align="center" />
          <el-table-column prop="pid" label="父节点" align="center" />
          <el-table-column prop="meta.title" label="展示名称" align="center" />
          <el-table-column prop="name" label="路由名称" align="center" />
          <el-table-column prop="path" label="路由路径" width="130" align="center" />
          <el-table-column prop="meta.title" label="是否隐藏" align="center">
            <template #default="scope">
              <el-tag v-if="scope.row.meta.title" type="success" effect="plain">显示</el-tag>
              <el-tag v-else type="warning" effect="plain">隐藏</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="component" label="组件路径" min-width="180" align="center" />
          <el-table-column fixed="right" label="操作" width="280" align="center">
            <template #default="scope">
              <el-button type="primary" text icon="Plus" size="small" @click="addMenuAction(scope.row.id)"
                >添加子菜单</el-button
              >
              <el-button type="primary" text icon="Edit" size="small" @click="editMenuAction(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" text icon="Delete" size="small" @click="deleteMenuAction(scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle">
      <warning-bar title="新增菜单，需要在角色管理内配置权限才可使用" />
      <el-form :inline="true" label-position="top" label-width="85px">
        <el-form-item label="路由名称" style="width: 30%">
          <el-input />
        </el-form-item>
        <el-form-item label="路由路径" style="width: 30%">
          <el-input />
        </el-form-item>
        <el-form-item label="前端组件" style="width: 30%">
          <el-input />
        </el-form-item>
        <el-form-item label="重定向" style="width: 30%">
          <el-input />
        </el-form-item>
        <el-form-item label="展示名称" style="width: 30%">
          <el-input />
        </el-form-item>
        <el-form-item label="是否隐藏" style="width: 30%">
          <el-select>
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" style="width: 30%">
          <el-select />
        </el-form-item>
        <el-form-item label="是否固定" style="width: 30%">
          <el-select>
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { usePermissionStoreHook } from "@/store/modules/permission"
import { type MenusDataFormat } from "@/api/system/menu"
import WarningBar from "@/components/warningBar/warningBar.vue"

const loading = ref<boolean>(false)
const dialogVisible = ref<boolean>(false)
const tableData = ref<MenusDataFormat[]>([])
const permissionStore = usePermissionStoreHook()
tableData.value = permissionStore.menusDataFormatList

const dialogTitle = ref<string>("")
const addMenuAction = (id: number) => {
  dialogTitle.value = "新增菜单"
  dialogVisible.value = true
}
const editMenuAction = (row: any) => {}
const deleteMenuAction = (row: any) => {}
</script>

<style lang="scss">
.toolbar-wrapper {
  margin-bottom: 20px;
}

.table-wrapper {
  margin-bottom: 20px;
}
</style>
