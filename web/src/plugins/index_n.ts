import type { App } from "vue"
import { installElementPlusIcons } from "./element-plus-icons"
import { installSvgIcon } from "./svg-icon"
import { installVxeTable } from "./vxe-table"

export function installPlugins(app: App) {
  installElementPlusIcons(app)
  installSvgIcon(app)
  installVxeTable(app)
}
