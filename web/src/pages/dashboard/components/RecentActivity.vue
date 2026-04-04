<template>
  <div class="recent-activity">
    <div class="activity-header">
      <h3>最近活动</h3>
      <el-button link type="primary" @click="$emit('viewAll')">查看全部</el-button>
    </div>
    
    <div class="activity-list">
      <div v-for="item in activities" :key="item.id" class="activity-item">
        <div class="activity-avatar">
          <el-avatar :size="36" :style="{ backgroundColor: getAvatarColor(item.userName) }">
            {{ getInitials(item.userName) }}
          </el-avatar>
        </div>
        <div class="activity-content">
          <div class="activity-title">
            <span class="username">{{ item.userName }}</span>
            <span class="action">{{ getActionText(item.method) }}</span>
            <span class="path">{{ item.path }}</span>
          </div>
          <div class="activity-meta">
            <el-tag :type="getStatusType(item.status)" size="small">
              {{ item.status }}
            </el-tag>
            <span class="time">{{ formatTime(item.createdAt) }}</span>
            <span v-if="item.respTime > 0" class="resp-time" :class="getRespTimeClass(item.respTime)">
              {{ item.respTime }}ms
            </span>
          </div>
        </div>
      </div>
      
      <el-empty v-if="activities.length === 0" description="暂无活动记录" />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { RecentOperation } from "@/api/sysMonitor/dashboard"

defineProps<{
  activities: RecentOperation[]
}>()

defineEmits<{
  (e: "viewAll"): void
}>()

const avatarColors = ["#409eff", "#67c23a", "#e6a23c", "#f56c6c", "#909399", "#8e44ad"]

function getInitials(name: string): string {
  if (!name) return "?"
  return name.slice(0, 2).toUpperCase()
}

function getAvatarColor(name: string): string {
  if (!name) return avatarColors[0]
  let hash = 0
  for (let i = 0; i < name.length; i++) {
    hash = name.charCodeAt(i) + ((hash << 5) - hash)
  }
  return avatarColors[Math.abs(hash) % avatarColors.length]
}

function getActionText(method: string): string {
  const actionMap: Record<string, string> = {
    GET: "访问",
    POST: "创建",
    PUT: "更新",
    PATCH: "修改",
    DELETE: "删除"
  }
  return actionMap[method.toUpperCase()] || "操作"
}

function getStatusType(status: number): "success" | "warning" | "danger" | "info" {
  if (status >= 200 && status < 300) return "success"
  if (status >= 400 && status < 500) return "warning"
  if (status >= 500) return "danger"
  return "info"
}

function getRespTimeClass(ms: number): string {
  if (ms < 100) return "fast"
  if (ms < 500) return "normal"
  return "slow"
}

function formatTime(timeStr: string): string {
  const date = new Date(timeStr)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)
  
  if (minutes < 1) return "刚刚"
  if (minutes < 60) return `${minutes}分钟前`
  if (hours < 24) return `${hours}小时前`
  if (days < 7) return `${days}天前`
  
  return date.toLocaleDateString("zh-CN")
}
</script>

<style scoped lang="scss">
.recent-activity {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
}

.activity-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  
  h3 {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.activity-list {
  max-height: 400px;
  overflow-y: auto;
}

.activity-item {
  display: flex;
  align-items: flex-start;
  padding: 12px 0;
  border-bottom: 1px solid #ebeef5;
  
  &:last-child {
    border-bottom: none;
  }
  
  &:hover {
    background-color: #f5f7fa;
    margin: 0 -20px;
    padding-left: 20px;
    padding-right: 20px;
  }
}

.activity-avatar {
  margin-right: 12px;
  flex-shrink: 0;
}

.activity-content {
  flex: 1;
  min-width: 0;
}

.activity-title {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 6px;
  font-size: 14px;
  margin-bottom: 6px;
  
  .username {
    font-weight: 600;
    color: #303133;
  }
  
  .action {
    color: #606266;
  }
  
  .path {
    color: #909399;
    font-family: monospace;
    font-size: 12px;
    background: #f5f7fa;
    padding: 2px 6px;
    border-radius: 4px;
    max-width: 200px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
}

.activity-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  
  .time {
    font-size: 12px;
    color: #909399;
  }
  
  .resp-time {
    font-size: 12px;
    font-family: monospace;
    
    &.fast {
      color: #67c23a;
    }
    
    &.normal {
      color: #e6a23c;
    }
    
    &.slow {
      color: #f56c6c;
    }
  }
}
</style>
