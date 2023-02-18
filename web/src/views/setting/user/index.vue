<template>
  <div class="app-container">
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
          <el-table-column prop="ID" label="ID" width="80" />
          <!-- <el-table-column prop="uuid" label="UUID" /> -->
          <el-table-column prop="username" label="用户名" align="center" />
          <el-table-column prop="phone" label="手机号" align="center" />
          <el-table-column prop="email" label="邮箱" align="center" />
          <el-table-column prop="active" label="状态" align="center">
            <template #default="scope">
              <el-switch v-model="scope.row.active" inline-prompt :active-value="true" :inactive-value="false" />
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" align="center" min-width="180px">
            <template #default="scope">
              <el-button type="primary" text icon="Setting" size="small">设置权限</el-button>
              <el-button type="primary" text icon="Edit" size="small">编辑</el-button>
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
    <el-dialog v-model="dialogVisible" title="新增用户" :before-close="handleClose" width="30%">
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
        label-position="left"
        style="width: 95%; margin-top: 15px"
      >
        <el-form-item label="角色名称" prop="roleName">
          <el-input v-model="formData.roleName" autocomplete="off" />
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
import { ref, reactive, watch } from "vue"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox } from "element-plus"
import { type UsersResponse, getUsersApi } from "@/api/system/user"
import { usePagination } from "@/hooks/usePagination"

const loading = ref<boolean>(false)
const { paginationData, handleCurrentChange, handleSizeChange } = usePagination()

const tableData = ref<UsersResponse[]>([])

const getTableData = async () => {
  loading.value = true
  const res = await getUsersApi({ page: paginationData.currentPage, pageSize: paginationData.pageSize })
  if (res.code === 0) {
    tableData.value = res.data.list
    paginationData.total = res.data.total
  }
  loading.value = false
}
getTableData()

const initForm = () => {
  formData.roleName = ""
}

const dialogVisible = ref<boolean>(false)
const handleClose = (done: Function) => {
  initForm()
  done()
}

const formRef = ref<FormInstance>()
const formData = reactive({
  roleName: ""
})
const formRules: FormRules = reactive({
  roleName: [{ required: true, trigger: "blur", message: "请填写角色名称" }]
})
const addDialog = () => {
  dialogVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogVisible.value = false
}

const operateAction = (formEl: FormInstance | undefined) => {
  // if (!formEl) return
  // formEl.validate(async (valid) => {
  //   if (valid) {
  //     const tempRole: reqRole = {
  //       roleName: formData.roleName
  //     }
  //     const res = await addRoleApi(tempRole)
  //     if (res.code === 0) {
  //       ElMessage({ type: "success", message: res.msg })
  //       const tempData: roleData = {
  //         ID: res.data.ID,
  //         roleName: res.data.roleName,
  //         menus: []
  //       }
  //       tableData.value.push(tempData)
  //     }
  //     initForm()
  //     dialogVisible.value = false
  //   }
  // })
}

const deleteRoleAction = (row: UsersResponse) => {
  // ElMessageBox.confirm("此操作将永久删除该角色, 是否继续?", "提示", {
  //   confirmButtonText: "确定",
  //   cancelButtonText: "取消",
  //   type: "warning"
  // }).then(() => {
  //   const index = tableData.value.indexOf(row)
  //   deleteRoleApi({ id: row.ID }).then((res) => {
  //     if (res.code === 0) {
  //       ElMessage({ type: "success", message: res.msg })
  //       tableData.value.splice(index, 1)
  //     }
  //   })
  // })
}

// 监听分页参数的变化
watch([() => paginationData.currentPage, () => paginationData.pageSize], getTableData, { immediate: true })
</script>

<style lang="scss" scoped>
.toolbar-wrapper {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
.table-wrapper {
  margin-bottom: 20px;
}
.pager-wrapper {
  display: flex;
  justify-content: flex-end;
}
</style>
