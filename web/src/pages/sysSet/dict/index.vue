<script lang="ts" setup>
import type { FormInstance } from "element-plus"
import type { dictDataModel } from "@/api/sysSet/dict"
import {
  addDictApi,
  delDictApi,
  editDictApi,
  getDictApi
} from "@/api/sysSet/dict"

import DictDetail from "./dictDetail.vue"

defineOptions({
  name: "Dict"
})

const selectID = ref(0)

const formData = reactive({
  chName: "",
  enName: ""
})
const rules = ref({
  chName: [
    {
      required: true,
      message: "请输入字典名（中）",
      trigger: "blur"
    }
  ],
  enName: [
    {
      required: true,
      message: "请输入字典名（英）",
      trigger: "blur"
    }
  ]
})

const tableData = ref<dictDataModel[]>([])

async function getTableData() {
  const res = await getDictApi()
  if (res.code === 0) {
    tableData.value = res.data
    selectID.value = res.data[0].id
  }
}

getTableData()

function toDetail(id: number) {
  selectID.value = id
}

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}

let oKind: operationKind
const dialogVisible = ref(false)
let activeRow: dictDataModel
async function editDictApiFunc(row: dictDataModel) {
  activeRow = row
  oKind = operationKind.Edit
  formData.chName = row.chName
  formData.enName = row.enName
  dialogVisible.value = true
}

function closeDialog() {
  formData.chName = ""
  formData.enName = ""
  oKind = operationKind.Add
  dialogVisible.value = false
}

async function delDictApiFunc(id: number) {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await delDictApi({ id })
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
    let res
    switch (oKind) {
      case operationKind.Add:
        res = await addDictApi(formData)
        if (res.code === 0) {
          ElMessage.success(res.msg)
          if (!tableData.value) {
            tableData.value = []
          }
          tableData.value.push(res.data)
        }
        break
      case operationKind.Edit:
        res = await editDictApi({ id: activeRow.id, ...formData })
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
function openDialog() {
  oKind = operationKind.Add
  formRef.value && formRef.value.clearValidate()
  dialogVisible.value = true
}
</script>

<template>
  <div class="app-container">
    <div style="display: flex !important;     gap: 1rem!important;    padding: .5rem!important;">
      <div class="first">
        <div style="display: flex!important;align-items: center!important;justify-content: space-between!important;">
          <span style="font-weight: 700 !important;">字典列表</span>
          <el-button type="primary" @click="openDialog">
            新增
          </el-button>
        </div>
        <el-scrollbar style="margin-top: 1rem !important;" max-height="calc(100vh - 240px)">
          <div
            v-for="dictionary in tableData"
            :key="dictionary.id"
            class="first-1"
            :class="
              selectID === dictionary.id
                ? 'text-active'
                : 'text-slate-700 dark:text-slate-50'
            "
            @click="toDetail(dictionary.id)"
          >
            <span class="max-w-160px truncate">{{ dictionary.chName }}</span>
            <div style="min-width: 40px!important;">
              <el-icon
                class="text-blue-500"
                @click.stop="editDictApiFunc(dictionary)"
              >
                <Edit />
              </el-icon>
              <el-icon
                class="ml-2 text-red-500"
                @click="delDictApiFunc(dictionary.id)"
              >
                <Delete />
              </el-icon>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <div class="second">
        <DictDetail :dict-id="selectID" />
      </div>
    </div>
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
        <el-form-item label="字典名（中）" prop="chName">
          <el-input
            v-model="formData.chName"
            placeholder="请输入字典名（中）"
            clearable
          />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="enName">
          <el-input
            v-model="formData.enName"
            placeholder="请输入字典名（英）"
            clearable
            :disabled="oKind === operationKind.Edit"
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

<style lang="scss" scoped>
.first {
  // flex: none !important;
  border-radius: 0.25rem !important;
  width: 13rem !important;
  padding: 1rem !important;
  --tw-bg-opacity: 1 !important;
  background-color: rgb(255 255 255 / var(--tw-bg-opacity, 1)) !important;
}

.first-1 {
  padding-top: 1rem !important;
  padding-bottom: 1rem !important;
  padding-left: 0.5rem !important;
  padding-right: 0.5rem !important;
  --tw-bg-opacity: 1 !important;
  background-color: rgb(249 250 251 / var(--tw-bg-opacity, 1)) !important;
  border-radius: 0.25rem !important;
  gap: 1rem !important;
  justify-content: space-between !important;
  align-items: center !important;
  display: flex !important;
  margin-top: 0.5rem !important;
}

.second {
  flex: 1 1 0% !important;
}

.truncate {
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.max-w-160px {
  max-width: 160px !important;
}

.text-blue-500 {
  --tw-text-opacity: 1 !important;
  color: rgb(59 130 246 / var(--tw-text-opacity, 1)) !important;
}

.text-red-500 {
  --tw-text-opacity: 1 !important;
  color: rgb(239 68 68 / var(--tw-text-opacity, 1)) !important;
}

.ml-2 {
  margin-left: 0.5rem !important;
}

.text-slate-700 {
  --tw-text-opacity: 1 !important;
  color: rgb(51 65 85 / var(--tw-text-opacity, 1)) !important;
}

.text-active {
  color: var(--el-color-primary) !important;
}
</style>
