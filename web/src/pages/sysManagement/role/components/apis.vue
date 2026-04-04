<script lang="ts" setup>
import type { ElTree as ElTree1 } from "element-plus"
import type { ApiTreeData } from "@/api/sysManagement/api"
import { ref, watch } from "vue"
import { apiGetElTreeApi } from "@/api/sysManagement/api"
import { updateRoleAPIPermissionsApi } from "@/api/sysManagement/casbin"

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
  return data.apiGroup?.includes(value) || data.path?.includes(value)
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const apiDefaultProps = {
  children: "children",
  label(data: any) {
    if (data.path && data.method) {
      return `${data.path} [${data.method}]`
    }
    return data.apiGroup
  }
}

const apiIds = ref<number[]>([])  // 使用权限ID而不是key
const apisTreeData = ref<ApiTreeData[]>([])
const apiKeyToId = ref<Map<string, number>>(new Map())  // key -> id 映射

function getTreeData() {
  apiGetElTreeApi({ id: props.id })
    .then((res) => {
      apisTreeData.value = res.data.list
      apiIds.value = res.data.checkedIds || []  // 使用checkedIds
      
      // 构建 key -> id 映射
      apiKeyToId.value.clear()
      res.data.list.forEach((group: ApiTreeData) => {
        group.children?.forEach((api: any) => {
          if (api.id && api.key) {
            apiKeyToId.value.set(api.key, api.id)
          }
        })
      })
    })
    .catch(() => {})
}
getTreeData()

function editAuthority() {
  // 获取选中的节点
  const checkedNodes = treeRef.value?.getCheckedNodes(false, true) as any[]
  const apiPermissionIds: number[] = []
  
  for (const item of checkedNodes) {
    // 只添加API节点（有id的节点），不添加分组节点
    if (item.id && item.path && item.method) {
      apiPermissionIds.push(item.id)
    }
  }
  
  updateRoleAPIPermissionsApi({ roleId: props.id, apiPermissionIds })
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
      <el-button type="primary" class="button" @click="editAuthority">
        更新
      </el-button>
    </div>
    <div class="tree-content">
      <ElTree
        ref="treeRef"
        :data="apisTreeData"
        :default-checked-keys="apiIds"
        node-key="id"
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
