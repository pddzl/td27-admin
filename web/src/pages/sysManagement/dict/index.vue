<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus";
import type { DictModel } from "@/api/sysManagement/dict";
import { reactive, ref } from "vue";
import {
  dictCreateApi,
  dictDeleteApi,
  DictListApi,
  dictUpdateApi,
} from "@/api/sysManagement/dict";
import { usePagination } from "@/common/composables/usePagination_n";

import DictDetail from "./dictDetail.vue";

defineOptions({
  name: "Dict",
});

const { paginationData, changeCurrentPage, changePageSize } = usePagination();

const drawer = ref(false);
const currentId = ref<number | null>(null);
const dictName = ref("");

function openDrawer(id: number, name: string) {
  currentId.value = id;
  drawer.value = true;
  dictName.value = name;
}

// function rowClick(row: DictModel) {
//   openDrawer(row.id, row.cn_name)
// }

const loading = ref(false);
const searchFormData = reactive({
  cn_name: "",
  en_name: "",
});

function handleSearch() {
  paginationData.currentPage = 1;
  paginationData.pageSize = 10;
  getTableData();
}

function resetSearch() {
  searchFormData.cn_name = "";
  searchFormData.en_name = "";
}

const tableData = ref<ListData<DictModel[]>>({
  list: [],
  total: 0,
  page: 0,
  pageSize: 0,
});

async function getTableData() {
  loading.value = true;
  try {
    const res = await DictListApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize,
    });
    if (res.code === 0) {
      tableData.value.list = res.data.list;
      paginationData.total = res.data.total;
    }
  } catch (error) {
    console.log(error);
  }
  loading.value = false;
}
getTableData();

// 分页
function handleSizeChange(value: number) {
  changePageSize(value);
  getTableData();
}

function handleCurrentChange(value: number) {
  changeCurrentPage(value);
  getTableData();
}

// 对话框
const formRef = ref<FormInstance>();
const opFormData = reactive({
  cn_name: "",
  en_name: "",
});

enum operationEnum {
  CREATE = "CREATE",
  UPDATE = "UPDATE",
}

const oKind: Ref<operationEnum | ""> = ref("");
const addFormRules: FormRules = reactive({
  cn_name: [{ required: true, trigger: "blur", message: "中文名称不能为空" }],
  en_name: [{ required: true, trigger: "blur", message: "英文名称不能为空" }],
});

function initForm() {
  formRef.value?.resetFields();
  opFormData.cn_name = "";
  opFormData.en_name = "";
}

const dialogVisible = ref(false);
const dialogTitle = ref("");
function handleClose(done: () => void) {
  initForm();
  done();
}

function createDialog() {
  dialogTitle.value = "新增字典";
  oKind.value = operationEnum.CREATE;
  dialogVisible.value = true;
}

function closeDialog() {
  dialogVisible.value = false;
  initForm();
}

function operateHandle(formEl: FormInstance | undefined) {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      if (oKind.value === operationEnum.CREATE) {
        const res = await dictCreateApi({ ...opFormData });
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg });
          tableData.value.list.push(res.data);
        }
      } else if (oKind.value === operationEnum.UPDATE) {
        const res = await dictUpdateApi({ id: activeRow.id, ...opFormData });
        if (res.code === 0) {
          ElMessage({ type: "success", message: res.msg });
          // 修改对应数据
          const index = tableData.value.list.indexOf(activeRow);
          tableData.value.list[index].cn_name = opFormData.cn_name;
          // tableData.value.list[index].en_name = opFormData.en_name
        }
      }
      closeDialog();
    }
  });
}

async function deleteHandle(id: number) {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const res = await dictDeleteApi({ id });
    if (res.code === 0) {
      ElMessage({ type: "success", message: res.msg });
      const index = tableData.value.list.indexOf(activeRow);
      tableData.value.list.splice(index, 1);
    }
  });
}

// 编辑dialog
let activeRow: DictModel;
function editDialogHandle(row: DictModel) {
  dialogTitle.value = "更新字典";
  oKind.value = operationEnum.UPDATE;
  opFormData.cn_name = row.cn_name;
  opFormData.en_name = row.en_name;
  activeRow = row;
  dialogVisible.value = true;
}
</script>

<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form :inline="true" :model="searchFormData">
        <el-form-item prop="path" label="中文名称">
          <el-input v-model="searchFormData.cn_name" placeholder="中文名称" />
        </el-form-item>
        <el-form-item prop="group" label="英文名称">
          <el-input v-model="searchFormData.en_name" placeholder="英文名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleSearch">
            查询
          </el-button>
          <el-button icon="Refresh" @click="resetSearch"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="createDialog">
            新增
          </el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button
              type="primary"
              icon="RefreshRight"
              circle
              plain
              @click="getTableData"
            />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData.list">
          <!-- <el-table-column type="selection" width="60" /> -->
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="cn_name" label="中文名称">
            <template #default="scope">
              <el-button
                type="primary"
                link
                @click="openDrawer(scope.row.id, scope.row.cn_name)"
              >
                {{ scope.row.cn_name }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column prop="en_name" label="英文名称" />
          <el-table-column label="操作">
            <template #default="scope">
              <el-button
                type="primary"
                text
                icon="Edit"
                size="small"
                @click="editDialogHandle(scope.row)"
              >
                编辑
              </el-button>
              <el-button
                type="danger"
                text
                icon="Delete"
                size="small"
                @click="deleteHandle(scope.row.id)"
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
      :title="dialogTitle"
      :before-close="handleClose"
      width="30%"
    >
      <el-form
        ref="formRef"
        :model="opFormData"
        :rules="addFormRules"
        label-width="80px"
      >
        <el-form-item label="中文名称" prop="cn_name">
          <el-input
            v-model="opFormData.cn_name"
            placeholder="请输入字典名（中）"
          />
        </el-form-item>
        <el-form-item label="英文名称" prop="en_name">
          <el-input
            v-model="opFormData.en_name"
            placeholder="请输入字典名（英）"
            :disabled="oKind === operationEnum.UPDATE"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog"> 取消 </el-button>
          <el-button type="primary" @click="operateHandle(formRef)">
            确认
          </el-button>
        </div>
      </template>
    </el-dialog>
    <DictDetail
      v-model:drawer="drawer"
      :dict-id="currentId"
      :dict-name="dictName"
    />
  </div>
</template>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 5px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
