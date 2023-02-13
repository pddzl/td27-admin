<template>
  <div>
    <span class="td27-icon" style="position: absolute; z-index: 9999; padding: 3px 10px 0">
      <svg-icon :name="metaData.icon" />
    </span>
    <el-select
      v-model="metaData.icon"
      class="td27-select"
      style="width: 100%"
      clearable
      filterable
      placeholder="请选择"
    >
      <el-option
        v-for="item in options"
        :key="item.key"
        class="select__option_item"
        :label="item.key"
        :value="item.key"
      >
        <span class="td27-icon">
          <svg-icon :name="item.label" />
        </span>
        <span style="text-align: left">{{ item.key }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue"

const props = defineProps({
  meta: {
    default: function () {
      return {}
    },
    type: Object
  }
})

const options = reactive([
  { key: "dashboard", label: "dashboard" },
  { key: "lock", label: "lock" },
  { key: "menu", label: "menu" },
  { key: "bug", label: "bug" }
])

const metaData = ref(props.meta)
</script>

<style>
.td27-icon {
  margin-right: 10px;
}

.td27-select .el-input__inner {
  padding: 0 20px !important;
}
</style>
