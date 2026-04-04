<template>
  <div class="quick-actions">
    <h3>快捷操作</h3>
    <div class="actions-grid">
      <div 
        v-for="action in actions" 
        :key="action.name"
        class="action-item"
        @click="handleAction(action)"
      >
        <div class="action-icon" :style="{ backgroundColor: action.bgColor, color: action.color }">
          <el-icon :size="24">
            <component :is="action.icon" />
          </el-icon>
        </div>
        <span class="action-name">{{ action.name }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router"
import {
  User,
  Key,
  SetUp,
  Document,
  Calendar,
  Folder,
  Monitor,
  Warning
} from "@element-plus/icons-vue"

const router = useRouter()

interface Action {
  name: string
  icon: any
  route?: string
  bgColor: string
  color: string
  handler?: () => void
}

const actions: Action[] = [
  {
    name: "用户管理",
    icon: User,
    route: "/sysManagement/user",
    bgColor: "#ecf5ff",
    color: "#409eff"
  },
  {
    name: "角色权限",
    icon: Key,
    route: "/sysManagement/role",
    bgColor: "#f0f9eb",
    color: "#67c23a"
  },
  {
    name: "菜单管理",
    icon: SetUp,
    route: "/sysManagement/menu",
    bgColor: "#fdf6ec",
    color: "#e6a23c"
  },
  {
    name: "操作日志",
    icon: Document,
    route: "/sysMonitor/operationLog",
    bgColor: "#f4f4f5",
    color: "#909399"
  },
  {
    name: "定时任务",
    icon: Calendar,
    route: "/sysTool/cron",
    bgColor: "#fef0f0",
    color: "#f56c6c"
  },
  {
    name: "文件管理",
    icon: Folder,
    route: "/sysTool/file",
    bgColor: "#f5f7fa",
    color: "#606266"
  },
  {
    name: "系统监控",
    icon: Monitor,
    route: "/sysMonitor/dashboard",
    bgColor: "#e6f7ff",
    color: "#1890ff"
  },
  {
    name: "异常页面",
    icon: Warning,
    route: "/error/404",
    bgColor: "#fff2f0",
    color: "#ff4d4f"
  }
]

function handleAction(action: Action) {
  if (action.handler) {
    action.handler()
  } else if (action.route) {
    router.push(action.route)
  }
}
</script>

<style scoped lang="scss">
.quick-actions {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
  
  h3 {
    margin: 0 0 16px 0;
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  
  &:hover {
    background-color: #f5f7fa;
    transform: translateY(-2px);
  }
}

.action-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border-radius: 12px;
  transition: all 0.3s;
}

.action-name {
  font-size: 13px;
  color: #606266;
  text-align: center;
}
</style>
