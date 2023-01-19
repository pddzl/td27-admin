import { ref } from "vue"
import { defineStore } from "pinia"
import layoutSettings from "@/config/layout"

export const useSettingsStore = defineStore("settings", () => {
  const fixedHeader = ref<boolean>(layoutSettings.fixedHeader)
  const showSettings = ref<boolean>(layoutSettings.showSettings)
  const showTagsView = ref<boolean>(layoutSettings.showTagsView)
  const showSidebarLogo = ref<boolean>(layoutSettings.showSidebarLogo)

  return {
    fixedHeader,
    showSettings,
    showTagsView,
    showSidebarLogo
  }
})
