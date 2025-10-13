<script lang="ts" setup>
import type { CascaderOption, FormInstance } from "element-plus"
import type { dictDetailDataModel } from "@/api/sysSet/dictDetail"
import {
  addDictDetailApi,
  delDictDetailApi,
  editDictDetailApi,
  getDictDetailApi,
  getDictDetailFlatApi
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
  dictId: props.dictId,
  parentId: undefined as number | undefined,
  description: ""
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

const treeOptions = ref<CascaderOption[]>([])

function mapTreeOptions(list: dictDetailDataModel[]): CascaderOption[] {
  return list.map(node => ({
    value: node.id,
    label: node.label,
    children: node.children?.length ? mapTreeOptions(node.children) : undefined
  }))
}

async function setTreeOptions() {
  if (!props.dictId) return
  const res = await getDictDetailApi({
    page: 1,
    pageSize: 1000,
    dictId: props.dictId
  })
  if (res.code === 0) {
    treeOptions.value = mapTreeOptions(res.data.list)
  }
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

async function getTableDataFlat() {
  if (!props.dictId) return
  const res = await getDictDetailFlatApi({
    dictId: props.dictId
  })
  if (res.code === 0) {
    console.log("flat", res.data)
  }
}
getTableDataFlat()

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
  formData.parentId = row.parentId
  formData.description = row.description
  dialogVisible.value = true
}

function closeDialog() {
  formData.label = ""
  formData.value = ""
  formData.sort = 0
  formData.parentId = undefined
  formData.description = ""
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
          // todo
          // it just make effect in no parent dictDetail
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

function openAddDialog(parent?: dictDetailDataModel) {
  oKind = operationKind.Add
  formRef.value?.clearValidate()
  formData.label = ""
  formData.value = ""
  formData.sort = 0
  formData.dictId = props.dictId
  formData.parentId = parent ? parent.id : undefined
  dialogVisible.value = true
}

watch(
  () => props.dictId,
  () => {
    getTableData()
    setTreeOptions()
  }
)
</script>

<template>
  <div>
    <el-card shadow="never">
      <div class="toolbar-wrapper">
        <el-button type="primary" icon="plus" @click="openAddDialog()">
          新增字典项
        </el-button>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table
          :data="tableData" style="width: 100%" row-key="id" border default-expand-all
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <el-table-column prop="label" label="展示值" />
          <el-table-column prop="value" label="字典值" />
          <el-table-column prop="sort" label="排序标记" />
          <el-table-column prop="description" label="描述" />
          <el-table-column label="创建日期" width="180">
            <template #default="scope">
              {{ formatDateTime(scope.row.createdAt) }}
            </template>
          </el-table-column>
          <el-table-column label="更新日期" width="180">
            <template #default="scope">
              {{ formatDateTime(scope.row.updatedAt) }}
            </template>
          </el-table-column>
          <el-table-column label="操作" width="260">
            <template #default="scope">
              <el-button type="primary" icon="edit" link @click="editDictDetailApiFunc(scope.row)">
                编辑
              </el-button>
              <el-button type="primary" icon="Plus" link @click="openAddDialog(scope.row)">
                新增子项
              </el-button>
              <el-button type="danger" icon="Delete" link @click="delDictDetailApiFunc(scope.row.id)">
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pager-wrapper">
        <el-pagination
          background :layout="paginationData.layout" :page-sizes="paginationData.pageSizes"
          :total="paginationData.total" :page-size="paginationData.pageSize" :current-page="paginationData.currentPage"
          @size-change="handleSizeChange" @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <el-dialog v-model="dialogVisible" :show-close="false" :before-close="closeDialog" width="600">
      <el-form ref="formRef" :model="formData" :rules="rules">
        <el-form-item label="父级" prop="parentId">
          <el-cascader
            v-model="formData.parentId" style="width: 100%" :options="treeOptions" placeholder="选择父级字典项" clearable
            :props="{ checkStrictly: true, emitPath: false }"
          />
        </el-form-item>
        <el-form-item label="展示值" prop="label">
          <el-input v-model="formData.label" placeholder="请输入展示值" clearable />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input v-model="formData.value" placeholder="请输入字典值" clearable :disabled="oKind === operationKind.Edit" />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number v-model.number="formData.sort" placeholder="排序标记" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input type="textarea" v-model="formData.description" placeholder="描述" />
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
