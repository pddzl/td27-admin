<script lang="ts" setup>
import type { FormInstance } from "element-plus"
import type { DictModel } from "@/api/sysManagement/dict"
import {
  dictCreateApi,
  dictDeleteApi,
  DictListApi,
  dictUpdateApi
} from "@/api/sysManagement/dict"

import DictDetail from "./dictDetail.vue"

defineOptions({
  name: "Dict"
})

const selectID = ref(0)

const formData = reactive({
  cn_name: "",
  en_name: ""
})
const rules = ref({
  cn_name: [
    {
      required: true,
      message: "请输入字典名（中）",
      trigger: "blur"
    }
  ],
  en_name: [
    {
      required: true,
      message: "请输入字典名（英）",
      trigger: "blur"
    }
  ]
})

const tableData = ref<DictModel[]>([])

async function getTableData() {
  const res = await DictListApi()
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
let activeRow: DictModel
async function editDictApiFunc(row: DictModel) {
  activeRow = row
  oKind = operationKind.Edit
  formData.cn_name = row.cn_name
  formData.en_name = row.en_name
  dialogVisible.value = true
}

function closeDialog() {
  formData.cn_name = ""
  formData.en_name = ""
  oKind = operationKind.Add
  dialogVisible.value = false
}

async function delDictApiFunc(id: number) {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await dictDeleteApi({ id })
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
        res = await dictCreateApi(formData)
        if (res.code === 0) {
          ElMessage.success(res.msg)
          if (!tableData.value) {
            tableData.value = []
          }
          tableData.value.push(res.data)
        }
        break
      case operationKind.Edit:
        res = await dictUpdateApi({ id: activeRow.id, ...formData })
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
    <div class="flex gap-4 p-2">
      <div
        class="flex-none w-52 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900 rounded p-4"
      >
        <div class="flex justify-between items-center">
          <span class="text font-bold">字典列表</span>
          <el-button type="primary" @click="openDialog">
            新增
          </el-button>
        </div>
        <el-scrollbar class="mt-4" max-height="calc(100vh - 240px)">
          <div
            v-for="dictionary in tableData"
            :key="dictionary.id"
            class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 dark:hover:bg-blue-900 bg-gray-50 dark:bg-gray-800 gap-4"
            :class="
              selectID === dictionary.id
                ? 'text-active'
                : 'text-slate-700 dark:text-slate-50'
            "
            @click="toDetail(dictionary.id)"
          >
            <span class="max-w-[160px] truncate">{{ dictionary.cn_name }}</span>
            <div class="min-w-[40px]">
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
      <div class="flex-1 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900">
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
        <el-form-item label="字典名（中）" prop="cn_name">
          <el-input
            v-model="formData.cn_name"
            placeholder="请输入字典名（中）"
            clearable
          />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="en_name">
          <el-input
            v-model="formData.en_name"
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
