<script lang="ts" setup>
import type { ElTree as ElTree1 } from "element-plus"
import type { ApiChild, ApiTreeData } from "@/api/sysManagement/api"
import { ref, watch } from "vue"
import { apiGetElTreeApi } from "@/api/sysManagement/api"
import { rebuildRolePermissionApi } from "@/api/sysManagement/role_permission"

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
  return data.group_cn?.includes(value) || data.group_en?.includes(value) || data.description?.includes(value)
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const apiDefaultProps = {
  children: "children",
  label(data: any) {
    if (data.description) {
      return data.description
    }
    return data.key
  }
}

const apiIds = ref<number[]>([]) // 使用权限ID而不是key
const apisTreeData = ref<ApiTreeData[]>([])

function getTreeData() {
  apiGetElTreeApi({ id: props.id, from_source: "role" })
    .then((res) => {
      apisTreeData.value = res.data.list
      apiIds.value = res.data.checkedIds || [] // 使用checkedIds
    })
    .catch(() => {})
}
getTreeData()

function rebuildHandle() {
  // 获取选中的节点
  const checkedNodes = treeRef.value?.getCheckedNodes(false, true) as ApiChild[]
  const domainIds: number[] = []

  for (const item of checkedNodes) {
    // 只添加API节点（有id的节点），不添加分组节点
    if (item.key && item.path && item.method) {
      domainIds.push(item.id)
    }
  }

  rebuildRolePermissionApi({ role_id: props.id, domain_ids: domainIds, domain: "api" })
    .then((res) => {
      if (res.code === 0 || res.code === 200) {
        ElMessage({ type: "success", message: "更新成功" })
      }
    })
    .catch(() => {})
}
</script>

<template>
  <div>
    <div class="clearfix">
      <el-input v-model="filterText" class="fitler" placeholder="筛选API" />
      <el-button type="primary" class="button" @click="rebuildHandle">
        更新
      </el-button>
    </div>
    <div class="tree-content">
      <ElTree
        ref="treeRef"
        :data="apisTreeData"
        :default-checked-keys="apiIds"
        node-key="key"
        default-expand-all
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
