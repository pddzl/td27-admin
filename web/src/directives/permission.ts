import type { Directive, DirectiveBinding } from "vue"
import { checkPermissionsBatch } from "@/composables/usePermission"

const cache: Record<string, boolean> = {}

async function updateEl(el: HTMLElement, binding: DirectiveBinding<string | string[]>) {
  const codes = Array.isArray(binding.value) ? binding.value : [binding.value]
  const modifier = binding.modifiers
  
  const unchecked = codes.filter(code => cache[code] === undefined)
  if (unchecked.length) {
    const result = await checkPermissionsBatch(unchecked)
    Object.assign(cache, result)
  }
  
  const hasPermission = codes.some(code => cache[code] === true)
  
  if (!hasPermission) {
    if (modifier.disable) {
      el.setAttribute("disabled", "true")
      ;(el as any).classList?.add("is-disabled")
    } else {
      el.style.display = "none"
    }
  } else {
    el.removeAttribute("disabled")
    ;(el as any).classList?.remove?.("is-disabled")
    el.style.display = ""
  }
}

export const vPermission: Directive = {
  mounted(el, binding) {
    updateEl(el, binding)
  },
  updated(el, binding) {
    updateEl(el, binding)
  }
}
