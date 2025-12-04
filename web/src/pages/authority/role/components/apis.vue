<script lang="ts" setup>
import type { ElTree as ElTree1 } from "element-plus"
import type { ApiTreeData } from "@/api/authority/api"
import type { CasbinInfo } from "@/api/base/casbin"
import { ref, watch } from "vue"
import { getElTreeApi } from "@/api/authority/api"
import { editCasbinApi } from "@/api/base/casbin"

const props = defineProps({
  id: {
    type: Number,
    default: 0
  }
})

const filterText = ref("")
const treeRef = ref<InstanceType<typeof ElTree1>>()

function filterNode(value: string, data: any) {
  if (!value) return true
  return data.apiGroup.includes(value)
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const apiDefaultProps = {
  children: "children",
  label(data: any) {
    return data.apiGroup
  }
}

const apiIds = ref<string[]>()
// const apiIds = ["/base/login,POST"]
const apisTreeData = ref<ApiTreeData[]>([])
function getTreeData() {
  getElTreeApi({ id: props.id })
    .then((res) => {
      apisTreeData.value = res.data.list
      apiIds.value = res.data.checkedKey
    })
    .catch(() => {})
}
getTreeData()

function editAuthority() {
  const casbinInfos: CasbinInfo[] = []
  for (const item of treeRef.value?.getCheckedNodes() as any[]) {
    if (item.path && item.method) {
      const casbinInfo: CasbinInfo = {
        path: item.path,
        method: item.method
      }
      casbinInfos.push(casbinInfo)
    }
  }
  editCasbinApi({ roleId: props.id, casbinInfos })
    .then((res) => {
      if (res.code === 0) {
        ElMessage({ type: "success", message: res.msg })
      }
    })
    .catch(() => {})
}
</script>

<template>
  <div>
    <div class="clearfix">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button type="primary" class="button" @click="editAuthority">
        更新
      </el-button>
    </div>
    <div class="tree-content">
      <ElTree
        ref="treeRef"
        :data="apisTreeData"
        :default-checked-keys="apiIds"
        default-expand-all
        node-key="key"
        highlight-current
        :props="apiDefaultProps"
        show-checkbox
        :filter-node-method="filterNode"
      />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.button {
  float: right;
  margin-right: 5%;
}
.tree-content {
  overflow: auto;
  height: calc(100vh - 160px);
  margin-top: 10px;
}
.clearfix::after {
  content: "";
  display: block;
  height: 0;
  clear: both;
  visibility: hidden;
}
.fitler {
  width: 80%;
}
</style>
