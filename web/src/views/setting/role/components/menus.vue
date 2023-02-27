<template>
  <div>
    <div class="clearfix">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button type="primary" class="button">更新</el-button>
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
import { ElTree } from "element-plus"
import { type MenusData, getAllMenusApi } from "@/api/system/menu"

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

let menuIds: number[]
const menuTreeData = ref<MenusData[]>([])
const getTreeData = (id: number) => {
  getAllMenusApi({ id: id })
    .then((res) => {
      menuTreeData.value = res.data.list
      menuIds = res.data.menuIds
    })
    .catch(() => {})
}
getTreeData(props.id)
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
