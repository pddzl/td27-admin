<template>
  <div class="navigation-bar">
    <Hamburger :is-active="sidebar.opened" class="hamburger" @toggle-click="toggleSidebar" />
    <Breadcrumb class="breadcrumb" />
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
import { useRouter } from "vue-router"
import { storeToRefs } from "pinia"
import { useAppStore } from "@/store/modules/app"
import { useUserStore } from "@/store/modules/user"
import Breadcrumb from "../Breadcrumb/index.vue"
import Hamburger from "../Hamburger/index.vue"
import ThemeSwitch from "@/components/ThemeSwitch/index.vue"
import Screenfull from "@/components/Screenfull/index.vue"
import { joinInBlacklistApi } from "@/api/system/jwt"
import { useSettingsStore } from "@/store/modules/settings"

const router = useRouter()
const appStore = useAppStore()
const settingsStore = useSettingsStore()
const userStore = useUserStore()

const { sidebar } = storeToRefs(appStore)
const { showThemeSwitch, showScreenfull } = storeToRefs(settingsStore)

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
  background: #fff;
  .hamburger {
    display: flex;
    align-items: center;
    height: 100%;
    float: left;
    padding: 0 15px;
    cursor: pointer;
  }
  .breadcrumb {
    float: left;
    // 参考 Bootstrap 的响应式设计将宽度设置为 576
    @media screen and (max-width: 576px) {
      display: none;
    }
  }
  .right-menu {
    float: right;
    margin-right: 10px;
    height: 100%;
    display: flex;
    align-items: center;
    color: #606266;
    .right-menu-item {
      padding: 0 10px;
      cursor: pointer;
    }
  }
}
.el-button.is-plain {
  --el-button-hover-border-color: #ffffff;
}
</style>
