<script setup lang="ts">
import type { DashboardStats, RecentOperation, SystemInfo } from "@/api/sysMonitor/dashboard"
import {
  Calendar,
  CircleCheck,
  DocumentChecked,
  Key,
  OfficeBuilding,
  User
} from "@element-plus/icons-vue"
import { ElMessage } from "element-plus"
import { onMounted, ref } from "vue"
import { useRouter } from "vue-router"
import {
  getDashboardStatisticsApi,
  getRecentOperationsApi,
  getSystemInfoApi
} from "@/api/sysMonitor/dashboard"

import { useUserStore } from "@/pinia/stores/user_n"
import QuickActions from "./components/QuickActions.vue"
import RecentActivity from "./components/RecentActivity.vue"
import StatsCard from "./components/StatsCard.vue"
import SystemStatus from "./components/SystemStatus.vue"
import WelcomeBanner from "./components/WelcomeBanner.vue"

const router = useRouter()
const userStore = useUserStore()

// Stats data
const stats = ref<Partial<DashboardStats>>({})

const statsList = [
  {
    key: "userCount" as const,
    label: "用户总数",
    icon: User,
    iconColor: "#409eff",
    iconBgColor: "#ecf5ff"
  },
  {
    key: "roleCount" as const,
    label: "角色总数",
    icon: Key,
    iconColor: "#67c23a",
    iconBgColor: "#f0f9eb"
  },
  {
    key: "deptCount" as const,
    label: "部门总数",
    icon: OfficeBuilding,
    iconColor: "#e6a23c",
    iconBgColor: "#fdf6ec"
  },
  {
    key: "operationCount" as const,
    label: "今日操作",
    icon: DocumentChecked,
    iconColor: "#909399",
    iconBgColor: "#f4f4f5"
  },
  {
    key: "cronCount" as const,
    label: "定时任务",
    icon: Calendar,
    iconColor: "#f56c6c",
    iconBgColor: "#fef0f0"
  },
  {
    key: "activeCronCount" as const,
    label: "活跃任务",
    icon: CircleCheck,
    iconColor: "#8e44ad",
    iconBgColor: "#f5eef8"
  }
]

// Recent operations
const recentOperations = ref<RecentOperation[]>([])

// System info
const systemInfo = ref<SystemInfo>({
  appName: "TD27 Admin",
  version: "v3.0.0",
  goVersion: "",
  os: "",
  arch: "",
  numCpu: 0,
  numGoroutine: 0,
  startTime: ""
})

// Fetch dashboard data
async function fetchDashboardData() {
  try {
    const [statsRes, operationsRes, systemRes] = await Promise.all([
      getDashboardStatisticsApi(),
      getRecentOperationsApi(),
      getSystemInfoApi()
    ])

    if (statsRes.code === 0) {
      stats.value = statsRes.data
    }

    if (operationsRes.code === 0) {
      recentOperations.value = operationsRes.data
    }

    if (systemRes.code === 0) {
      systemInfo.value = systemRes.data
    }
  } catch (error) {
    ElMessage.error("获取仪表盘数据失败")
    console.error(error)
  }
}

function handleViewAllOperations() {
  router.push("/sysMonitor/operationLog")
}

onMounted(() => {
  fetchDashboardData()
})
</script>

<template>
  <div class="dashboard-container">
    <!-- Welcome Banner -->
    <WelcomeBanner :username="userStore.userInfo.username || '访客'" />

    <!-- Stats Cards -->
    <el-row :gutter="16" class="stats-row">
      <el-col :xs="24" :sm="12" :md="8" :lg="4" v-for="stat in statsList" :key="stat.key">
        <StatsCard
          :icon="stat.icon"
          :value="stats[stat.key] || 0"
          :label="stat.label"
          :icon-color="stat.iconColor"
          :icon-bg-color="stat.iconBgColor"
        />
      </el-col>
    </el-row>

    <!-- Quick Actions -->
    <el-row class="section-row">
      <el-col :span="24">
        <QuickActions />
      </el-col>
    </el-row>

    <!-- Main Content -->
    <el-row :gutter="16" class="main-content">
      <el-col :xs="24" :lg="16">
        <RecentActivity
          :activities="recentOperations"
          @view-all="handleViewAllOperations"
        />
      </el-col>
      <el-col :xs="24" :lg="8">
        <SystemStatus :info="systemInfo" />
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.dashboard-container {
  padding: 16px;
}

.stats-row {
  margin-top: 16px;

  .el-col {
    margin-bottom: 16px;
  }
}

.section-row {
  margin-top: 8px;
  margin-bottom: 16px;
}

.main-content {
  .el-col {
    margin-bottom: 16px;
  }
}
</style>
