<template>
  <div class="app-container">
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
        <el-table :data="tableData" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="40" />
          <el-table-column prop="id" label="ID" />
          <el-table-column prop="name" label="名称" />
          <el-table-column prop="strategy" label="策略">
            <template #default="scope">
              <el-tag type="success">{{ strategyFilter(scope.row.strategy) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="open" label="状态">
            <template #default="scope">
              <el-switch
                v-model="scope.row.open"
                inline-prompt
                :active-value="true"
                :inactive-value="false"
                active-text="开启"
                inactive-text="关闭"
                :before-change="() => switchAction(scope.row)"
              />
            </template>
          </el-table-column>
          <el-table-column prop="entryId" label="Cron ID" />
          <el-table-column label="操作" align="center" fixed="right" min-width="120">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" :before-close="handleClose" width="35%" destroy-on-close>
      <el-form ref="formRef" :model="opFormData" :rules="addFormRules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="opFormData.name" />
        </el-form-item>
        <el-form-item label="表达式" prop="expression">
          <el-input v-model="opFormData.expression" placeholder="second / min / hour / day / mon / week" />
        </el-form-item>
        <el-form-item label="方法" prop="method">
          <el-select
            v-model="opFormData.method"
            placeholder="请选择方法"
            clearable
            style="width: 100%"
            @change="changeMethod(opFormData.method)"
          >
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="参数" prop="extraParams" required>
          <div v-if="opFormData.method === 'clearTable'" style="width: 100%">
            <el-row
              v-for="(item, key) in opFormData.extraParams.tableInfo"
              :key="key"
              style="margin-bottom: 5px"
              justify="space-between"
            >
              <el-col :span="7">
                <el-input v-model="item.tableName" placeholder="表名称" />
              </el-col>
              <el-col :span="7">
                <el-input v-model="item.compareField" placeholder="比较字段" />
              </el-col>
              <el-col :span="7">
                <el-input v-model="item.interval" placeholder="时间间隔" />
              </el-col>
              <el-button
                type="danger"
                plain
                :icon="Delete"
                @click="removeTableInfo(item)"
                circle
                :disabled="key === 0"
              />
            </el-row>
            <el-button type="primary" plain :icon="Plus" @click="addTableInfo" style="width: 100%" />
          </div>
          <div v-else style="width: 100%">
            <el-input v-model="opFormData.extraParams.command" />
          </div>
        </el-form-item>
        <el-form-item label="策略" prop="strategy">
          <el-radio-group v-model="opFormData.strategy">
            <el-radio-button label="重复执行" value="always" />
            <el-radio-button label="执行一次" value="once" />
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="open">
          <el-switch v-model="opFormData.open" active-text="开启" inactive-text="关闭" />
        </el-form-item>
        <el-form-item label="描述" prop="comment">
          <el-input v-model="opFormData.comment" :rows="2" type="textarea" />
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
import { Plus, Delete } from "@element-plus/icons-vue"
import { reactive, ref } from "vue"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox, ElNotification } from "element-plus"
import { usePagination } from "@/hooks/usePagination"
import {
  type TableInfo,
  type cronDataModel,
  getCronListApi,
  addCronApi,
  switchCronApi,
  deleteCronApi,
  deleteCronByIds,
  editCronApi,
  METHOD
} from "@/api/sysTool/cron"
import { strategyFilter } from "./filter"

defineOptions({
  name: "Cron"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)

const methodOptions = [
  { value: "clearTable", label: "ClearTable" },
  { value: "shell", label: "SHELL" }
]

const tableData = ref<cronDataModel[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getCronListApi({
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
const handleSelectionChange = (val: cronDataModel[]) => {
  ids.value = val.map((item) => item.id)
  taskIds.value = ids.value.join(",")
}

// 对话框
const formRef = ref<FormInstance>()
const opFormData = reactive({
  name: "",
  method: "",
  expression: "",
  extraParams: {
    tableInfo: [] as TableInfo[],
    command: ""
  },
  strategy: "always",
  open: false,
  comment: ""
})

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}

let oKind: operationKind
const addFormRules: FormRules = reactive({
  name: [{ required: true, trigger: "blur", message: "名称不能为空" }],
  method: [{ required: true, trigger: "change", message: "方法不能为空" }],
  expression: [{ required: true, trigger: "blur", message: "表达式不能为空" }]
})

const initForm = () => {
  formRef.value?.resetFields()
  opFormData.name = ""
  opFormData.method = ""
  opFormData.expression = ""
  opFormData.strategy = "always"
  opFormData.extraParams = { tableInfo: [{ tableName: "", compareField: "", interval: "" }], command: "" }
  opFormData.open = false
  opFormData.comment = ""
}

const dialogVisible = ref(false)
const dialogTitle = ref("")
const handleClose = (done: Function) => {
  initForm()
  done()
}

const addDialog = () => {
  dialogTitle.value = "新增Cron"
  oKind = operationKind.Add
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
  initForm()
}

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  // 判断参数为空
  let empty
  if (opFormData.method === METHOD.ClearTable) {
    for (const element of opFormData.extraParams.tableInfo) {
      if (!element.tableName || !element.compareField || !element.interval) {
        empty = true
        break
      }
    }
  } else if (opFormData.method === METHOD.Shell) {
    if (!opFormData.extraParams.command) {
      empty = true
    }
  }
  if (empty) {
    ElNotification({
      title: "告警",
      message: "参数不能为空",
      type: "warning"
    })
    return
  }

  formEl
    .validate(async (valid) => {
      if (valid) {
        if (oKind === "Add") {
          const res = await addCronApi({ ...opFormData })
          if (res.code === 0) {
            ElMessage({ type: "success", message: res.msg })
            tableData.value.push(res.data)
          }
        } else if (oKind === "Edit") {
          const res = await editCronApi({ id: activeRow.id, ...opFormData })
          if (res.code === 0) {
            ElMessage({ type: "success", message: res.msg })
            // 修改对应数据
            const index = tableData.value.indexOf(activeRow)
            tableData.value.splice(index, 1, res.data)
          }
        }
        // 关闭对话框
        closeDialog()
      }
    })
    .catch(() => {})
}

// 删除cron
const handleDelete = (row: cronDataModel) => {
  ElMessageBox.confirm("此操作将永久删除该定时任务, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteCronApi({ id: row.id }).then((res) => {
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
let activeRow: cronDataModel
const editDialog = (row: any) => {
  dialogTitle.value = "编辑Cron"
  oKind = operationKind.Edit
  opFormData.name = row.name
  opFormData.method = row.method
  opFormData.expression = row.expression
  opFormData.strategy = row.strategy
  opFormData.extraParams = row.extraParams
  opFormData.open = row.open
  opFormData.comment = row.comment
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
  const res = await deleteCronByIds({ ids: ids.value })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: res.msg
    })
    deleteVisible.value = false
    getTableData()
  }
}

const addTableInfo = () => {
  const last = opFormData.extraParams.tableInfo[opFormData.extraParams.tableInfo.length - 1]
  if (last.tableName === "" || last.compareField === "" || last.interval === "") {
    return
  } else {
    opFormData.extraParams.tableInfo.push({ tableName: "", compareField: "", interval: "" })
  }
}

const removeTableInfo = (item: TableInfo) => {
  const index = opFormData.extraParams.tableInfo.indexOf(item)
  if (index !== -1) {
    opFormData.extraParams.tableInfo.splice(index, 1)
  }
}

const switchAction = (row: cronDataModel) => {
  return new Promise<boolean>((resolve, reject) => {
    switchCronApi({ id: row.id, open: !row.open })
      .then((res) => {
        if (res.code === 0) {
          if (!row.open) {
            ElMessage({ type: "success", message: "开启成功" })
          } else {
            ElMessage({ type: "success", message: "关闭成功" })
          }
          row.entryId = res.data.entryId
          return resolve(true)
        } else {
          return reject(false)
        }
      })
      .catch(() => {
        return reject(false)
      })
  })
}

const changeMethod = (method: string) => {
  if (method === "clearTable") {
    opFormData.extraParams.tableInfo = [{ tableName: "", compareField: "", interval: "" }]
    opFormData.extraParams.command = ""
  } else {
    opFormData.extraParams.tableInfo = []
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
