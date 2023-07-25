<template>
  <div class="navigation-bar">
    <Hamburger v-if="!isTop || isMobile" :is-active="sidebar.opened" class="hamburger" @toggle-click="toggleSidebar" />
    <Breadcrumb v-if="!isTop || isMobile" class="breadcrumb" />
    <Sidebar v-if="isTop && !isMobile" class="sidebar" />
    <div class="right-menu">
      <Screenfull v-if="showScreenfull" class="right-menu-item" />
      <ThemeSwitch v-if="showThemeSwitch" class="right-menu-item" />
      <el-dropdown class="right-menu-item">
        <el-button plain
          >{{ userStore.username }}<el-icon class="el-icon--right"><arrow-down /></el-icon
        ></el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <a target="_blank" href="https://github.com/pddzl/td27-admin">
              <el-dropdown-item>GitHub</el-dropdown-item>
            </a>
            <el-dropdown-item @click="toPersonal">个人中心</el-dropdown-item>
            <el-dropdown-item divided @click="logout">
              <span style="display: block">退出登录</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"
import { useRouter } from "vue-router"
import { storeToRefs } from "pinia"
import { useAppStore } from "@/store/modules/app"
import { useUserStore } from "@/store/modules/user"
import Breadcrumb from "../Breadcrumb/index.vue"
import Sidebar from "../Sidebar/index.vue"
import Hamburger from "../Hamburger/index.vue"
import ThemeSwitch from "@/components/ThemeSwitch/index.vue"
import Screenfull from "@/components/Screenfull/index.vue"
import { joinInBlacklistApi } from "@/api/system/jwt"
import { useSettingsStore } from "@/store/modules/settings"
import { DeviceEnum } from "@/constants/app-key"

const router = useRouter()
const appStore = useAppStore()
const settingsStore = useSettingsStore()
const userStore = useUserStore()

const { sidebar, device } = storeToRefs(appStore)
const { layoutMode, showThemeSwitch, showScreenfull } = storeToRefs(settingsStore)

const isTop = computed(() => layoutMode.value === "top")
const isMobile = computed(() => device.value === DeviceEnum.Mobile)

/** 切换侧边栏 */
const toggleSidebar = () => {
  appStore.toggleSidebar(false)
}

const logout = () => {
  // userStore.logout()
  // router.push("/login")
  joinInBlacklistApi()
    .then(() => {
      userStore.logout()
      router.push("/login")
    })
    .catch(() => {})
}

const toPersonal = () => {
  router.push({ name: "Profile" })
}
</script>

<style lang="scss" scoped>
.navigation-bar {
  height: var(--base-navigationbar-height);
  overflow: hidden;
  background: var(--base-header-bg-color);
  display: flex;
  justify-content: space-between;
  .hamburger {
    display: flex;
    align-items: center;
    height: 100%;
    padding: 0 15px;
    cursor: pointer;
  }
  .breadcrumb {
    flex: 1;
    // 参考 Bootstrap 的响应式设计将宽度设置为 576
    @media screen and (max-width: 576px) {
      display: none;
    }
  }
  .sidebar {
    flex: 1;
    // 设置 min-width 是为了让 Sidebar 里的 el-menu 宽度自适应
    min-width: 0px;
    :deep(.el-menu) {
      background-color: transparent;
    }
    :deep(.el-sub-menu) {
      &.is-active {
        .el-sub-menu__title {
          color: var(--el-menu-active-color) !important;
        }
      }
    }
  }
  .right-menu {
    margin-right: 10px;
    height: 100%;
    display: flex;
    align-items: center;
    color: #606266;
    .right-menu-item {
      padding: 0 10px;
      cursor: pointer;
      .right-menu-avatar {
        display: flex;
        align-items: center;
        .el-avatar {
          margin-right: 10px;
        }
        span {
          font-size: 16px;
        }
      }
    }
  }
}
</style>
