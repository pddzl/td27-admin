<script lang="ts" setup>
import type { FormInstance } from "element-plus"
import type { dictDetailDataModel } from "@/api/sysSet/dictDetail"
import {
  addDictDetailApi,
  delDictDetailApi,
  editDictDetailApi,
  getDictDetailApi
} from "@/api/sysSet/dictDetail"
import { usePagination } from "@/common/composables/usePagination_n"
import { formatDateTime } from "@/common/utils/datetime"

defineOptions({
  name: "DictDetail"
})

const props = defineProps({
  dictId: {
    type: Number,
    default: 0
  }
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const formData = reactive({
  label: "",
  value: "",
  sort: 0,
  dictId: props.dictId
})

const rules = ref({
  label: [
    {
      required: true,
      message: "请输入展示值",
      trigger: "blur"
    }
  ],
  value: [
    {
      required: true,
      message: "请输入字典值",
      trigger: "blur"
    }
  ],
  sort: [
    {
      required: true,
      message: "排序标记",
      trigger: "blur"
    }
  ]
})

const tableData = ref<dictDetailDataModel[]>([])

// 分页
function handleSizeChange(value: number) {
  changePageSize(value)
  getTableData()
}

function handleCurrentChange(value: number) {
  changeCurrentPage(value)
  getTableData()
}

// 查询
async function getTableData() {
  if (!props.dictId) return
  const res = await getDictDetailApi({
    page: paginationData.currentPage,
    pageSize: paginationData.pageSize,
    dictId: props.dictId
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    paginationData.total = res.data.total
  }
}

getTableData()

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}
let oKind: operationKind
let activeRow: dictDetailDataModel
const dialogVisible = ref(false)
async function editDictDetailApiFunc(row: dictDetailDataModel) {
  activeRow = row
  oKind = operationKind.Edit
  formData.label = row.label
  formData.value = row.value
  formData.sort = row.sort
  dialogVisible.value = true
}

function closeDialog() {
  formData.label = ""
  formData.value = ""
  formData.sort = 0
  oKind = operationKind.Add
  dialogVisible.value = false
}

async function delDictDetailApiFunc(id: number) {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await delDictDetailApi({ id })
    if (res.code === 0) {
      ElMessage({ type: "success", message: res.msg })
      const index = tableData.value.indexOf(activeRow)
      tableData.value.splice(index, 1)
    }
  })
}

const formRef = ref<FormInstance>()
async function operateAction(formEl: FormInstance | undefined) {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (!valid) return
    formData.dictId = props.dictId
    let res
    switch (oKind) {
      case operationKind.Add:
        res = await addDictDetailApi(formData)
        if (res.code === 0) {
          ElMessage.success(res.msg)
          if (!tableData.value) {
            tableData.value = []
          }
          tableData.value.push(res.data)
        }
        break
      case operationKind.Edit:
        res = await editDictDetailApi({ id: activeRow.id, ...formData })
        if (res.code === 0) {
          ElMessage.success(res.msg)
          const index = tableData.value.indexOf(activeRow)
          tableData.value[index] = res.data
        }
        break
      default:
        break
    }
    closeDialog()
  })
}
function addDialog() {
  oKind = operationKind.Add
  formRef.value && formRef.value.clearValidate()
  dialogVisible.value = true
}

watch(
  () => props.dictId,
  () => {
    getTableData()
  }
)
</script>

<template>
  <div>
    <el-card shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="plus" @click="addDialog">
          新增字典项
        </el-button>
      </div>
      <div class="table-wrapper">
        <el-table
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="id"
        >
          <el-table-column type="selection" width="55" />
          <el-table-column align="left" label="展示值" prop="label" />
          <el-table-column align="left" label="字典值" prop="value" />
          <el-table-column
            align="left"
            label="排序标记"
            prop="sort"
            width="120"
          />
          <el-table-column align="left" label="创建日期" width="180">
            <template #default="scope">
              {{ formatDateTime(scope.row.createdAt) }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="操作">
            <template #default="scope">
              <el-button
                type="primary"
                link
                icon="edit"
                @click="editDictDetailApiFunc(scope.row)"
              >
                编辑
              </el-button>
              <el-button
                type="danger"
                link
                icon="delete"
                @click="delDictDetailApiFunc(scope.row.id)"
              >
                删除
              </el-button>
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
          :current-page="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <el-dialog
      v-model="dialogVisible"
      :show-close="false"
      :before-close="closeDialog"
      width="600"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
      >
        <el-form-item label="展示值" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="请输入展示值"
            clearable
          />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input
            v-model="formData.value"
            placeholder="请输入字典值"
            clearable
            :disabled="oKind === operationKind.Edit"
          />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number
            v-model.number="formData.sort"
            placeholder="排序标记"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button @click="closeDialog">
            取 消
          </el-button>
          <el-button type="primary" @click="operateAction(formRef)">
            确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
