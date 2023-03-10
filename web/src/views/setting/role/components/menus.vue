<template>
  <div>
    <div class="clearfix">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button type="primary" class="button" @click="editRoleMenu">更新</el-button>
    </div>
    <div class="tree-content">
      <el-tree
        ref="treeRef"
        :data="menuTreeData"
        :default-checked-keys="menuIds"
        default-expand-all
        node-key="id"
        highlight-current
        :props="menuDefaultProps"
        show-checkbox
        :filter-node-method="filterNode"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue"
import { ElMessage, ElTree } from "element-plus"
import { type MenusData, getElTreeMenusApi } from "@/api/system/menu"
import { editRoleMenuApi } from "@/api/system/role"

const props = defineProps({
  id: {
    type: Number,
    default: 0
  }
})

const filterText = ref("")
const treeRef = ref<InstanceType<typeof ElTree>>()

const filterNode = (value: string, data: any) => {
  if (!value) return true
  return data.meta.title.includes(value)
}

watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const menuDefaultProps = {
  children: "children",
  label: function (data: any) {
    return data.meta.title
  }
}

const menuIds = ref<number[]>([])
// const menuIds = [2, 3, 4, 7, 8]
const menuTreeData = ref<MenusData[]>([])
const getTreeData = (id: number) => {
  getElTreeMenusApi({ id: id })
    .then((res) => {
      menuTreeData.value = res.data.list
      menuIds.value = res.data.menuIds
    })
    .catch(() => {})
}
getTreeData(props.id)

const editRoleMenu = () => {
  editRoleMenuApi({ roleId: props.id, ids: treeRef.value?.getCheckedKeys() as number[] })
    .then((res) => {
      if (res.code === 0) {
        ElMessage({ type: "success", message: res.msg })
      }
    })
    .catch(() => {})
}
</script>

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
