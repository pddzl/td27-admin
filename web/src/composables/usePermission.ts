import { ref, onMounted } from "vue"
import { batchCheckButtonApi, getUserButtonsApi } from "@/api/sysManagement/button"

const permissionCache = ref<Record<string, boolean>>({})
let loadingPromise: Promise<void> | null = null

async function loadPermissions() {
  if (loadingPromise) return loadingPromise
  
  loadingPromise = (async () => {
    try {
      const res = await getUserButtonsApi()
      if (res.code === 0) {
        const perms: Record<string, boolean> = {}
        res.data.forEach((code: string) => {
          perms[code] = true
        })
        permissionCache.value = perms
      }
    } catch (error) {
      console.error("加载按钮权限失败", error)
    } finally {
      loadingPromise = null
    }
  })()
  
  return loadingPromise
}

export function usePermission() {
  onMounted(() => {
    loadPermissions()
  })

  function hasPermission(buttonCode: string) {
    return permissionCache.value[buttonCode] === true
  }

  function hasAnyPermission(buttonCodes: string[]) {
    return buttonCodes.some(code => permissionCache.value[code] === true)
  }

  function hasAllPermissions(buttonCodes: string[]) {
    return buttonCodes.every(code => permissionCache.value[code] === true)
  }

  async function checkPageButtons(pagePath: string) {
    // This is handled by getPageButtonsApi per page
    return loadPermissions()
  }

  return {
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    checkPageButtons,
    permissionCache
  }
}

export async function checkPermissionsBatch(buttonCodes: string[]) {
  if (!buttonCodes.length) return {}
  try {
    const res = await batchCheckButtonApi(buttonCodes)
    if (res.code === 0) {
      return res.data
    }
  } catch (error) {
    console.error("批量检查权限失败", error)
  }
  return buttonCodes.reduce((acc, code) => {
    acc[code] = false
    return acc
  }, {} as Record<string, boolean>)
}
