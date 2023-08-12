<script lang="ts" setup>
import { watchEffect } from "vue"
import { storeToRefs } from "pinia"
import { useAppStore } from "@/store/modules/app"
import { useSettingsStore } from "@/store/modules/settings"
import useResize from "./hooks/useResize"
import LeftMode from "./LeftMode.vue"
import TopMode from "./TopMode.vue"
import LeftTopMode from "./LeftTopMode.vue"
import { Settings, RightPanel } from "./components"
import { DeviceEnum } from "@/constants/app-key"
import { getCssVariableValue, setCssVariableValue } from "@/utils"

/** Layout 布局响应式 */
useResize()

const appStore = useAppStore()
const settingsStore = useSettingsStore()

const { showSettings, layoutMode, showTagsView } = storeToRefs(settingsStore)

//#region 隐藏标签栏时删除其高度，是为了让 Logo 组件高度和 Header 区域高度始终一致
const cssVariableName = "--base-tagsview-height"
const v3TagsviewHeight = getCssVariableValue(cssVariableName)
watchEffect(() => {
  showTagsView.value
    ? setCssVariableValue(cssVariableName, v3TagsviewHeight)
    : setCssVariableValue(cssVariableName, "0px")
})
//#endregion
</script>

<template>
  <div>
    <!-- 左侧模式 -->
    <LeftMode v-if="layoutMode === 'left' || appStore.device === DeviceEnum.Mobile" />
    <!-- 顶部模式 -->
    <TopMode v-else-if="layoutMode === 'top'" />
    <!-- 混合模式 -->
    <LeftTopMode v-else-if="layoutMode === 'left-top'" />
    <!-- 右侧设置面板 -->
    <RightPanel v-if="showSettings">
      <Settings />
    </RightPanel>
  </div>
</template>
