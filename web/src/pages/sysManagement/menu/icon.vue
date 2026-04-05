<script lang="ts" setup>
import { computed } from "vue"

/**
 * 1. Strongly typed icon names (must match your SvgIcon declaration)
 */
export type IconName = "link"
  | "menu"
  | "access"
  | "bug"
  | "config"
  | "dashboard"
  | "file"
  | "fullscreen-exit"
  | "fullscreen"
  | "keyboard-down"
  | "keyboard-enter"
  | "keyboard-esc"
  | "keyboard-up"
  | "load"
  | "lock"
  | "monitor"
  | "network"
  | "plus"
  | "radar"
  | "search"
  | "setting"

/**
 * 2. Props (v-model standard)
 */
const props = defineProps<{
  modelValue: IconName | "" // allow empty for clearable
}>()

/**
 * 3. Emits
 */
const emit = defineEmits<{
  (e: "update:modelValue", val: IconName | ""): void
}>()

/**
 * 4. Icon options (strictly typed)
 */
const options: IconName[] = [
  "dashboard",
  "setting",
  "lock",
  "menu",
  "bug",
  "network",
  "plus",
  "load",
  "config",
  "link",
  "access",
  "file",
  "monitor"
]

/**
 * 5. v-model bridge
 */
const iconValue = computed<IconName | "">({
  get: () => props.modelValue,
  set: val => emit("update:modelValue", val)
})

/**
 * 6. Safe icon fallback
 */
function getValidIconName(name: string): IconName {
  return options.includes(name as IconName)
    ? (name as IconName)
    : "setting"
}
</script>

<template>
  <el-select
    v-model="iconValue"
    clearable
    filterable
    placeholder="请选择图标"
    style="width: 100%"
  >
    <!-- 当前选中图标 -->
    <template #prefix>
      <SvgIcon
        v-if="iconValue"
        :name="getValidIconName(iconValue)"
        class="td27-icon"
      />
    </template>

    <!-- 下拉选项 -->
    <el-option
      v-for="item in options"
      :key="item"
      :label="item"
      :value="item"
      class="select__option_item"
    >
      <span class="td27-icon">
        <SvgIcon :name="item" />
      </span>
      <span>{{ item }}</span>
    </el-option>
  </el-select>
</template>

<style scoped>
.td27-icon {
  color: rgb(132, 146, 166);
  font-size: 14px;
  margin-right: 10px;
}

.select__option_item {
  display: flex;
  align-items: center;
}
</style>
