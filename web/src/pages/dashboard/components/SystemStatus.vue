<template>
  <div class="system-status">
    <div class="status-header">
      <h3>系统状态</h3>
      <el-tag type="success" size="small">运行中</el-tag>
    </div>
    
    <div class="status-content">
      <div class="status-item">
        <div class="status-label">
          <el-icon><Monitor /></el-icon>
          <span>应用名称</span>
        </div>
        <div class="status-value">{{ info.appName }}</div>
      </div>
      
      <div class="status-item">
        <div class="status-label">
          <el-icon><CollectionTag /></el-icon>
          <span>版本</span>
        </div>
        <div class="status-value">{{ info.version }}</div>
      </div>
      
      <div class="status-item">
        <div class="status-label">
          <el-icon><Platform /></el-icon>
          <span>运行环境</span>
        </div>
        <div class="status-value">{{ info.os }} / {{ info.arch }}</div>
      </div>
      
      <div class="status-item">
        <div class="status-label">
          <el-icon><Cpu /></el-icon>
          <span>CPU 核心</span>
        </div>
        <div class="status-value">{{ info.numCpu }} 核</div>
      </div>
      
      <div class="status-item">
        <div class="status-label">
          <el-icon><Connection /></el-icon>
          <span>Go 协程</span>
        </div>
        <div class="status-value">{{ info.numGoroutine }}</div>
      </div>
      
      <div class="status-item">
        <div class="status-label">
          <el-icon><Timer /></el-icon>
          <span>启动时间</span>
        </div>
        <div class="status-value">{{ info.startTime }}</div>
      </div>
    </div>
    
    <div class="status-footer">
      <el-progress 
        :percentage="systemLoad" 
        :color="loadColor"
        :stroke-width="8"
        :show-text="true"
      />
      <div class="load-text">系统负载</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue"
import { Monitor, CollectionTag, Platform, Cpu, Connection, Timer } from "@element-plus/icons-vue"
import type { SystemInfo } from "@/api/sysMonitor/dashboard"

const props = defineProps<{
  info: SystemInfo
}>()

// 模拟系统负载（实际应该从后端获取）
const systemLoad = computed(() => {
  // 基于协程数估算负载
  const baseLoad = Math.min(props.info.numGoroutine / 10, 100)
  return Math.min(Math.round(baseLoad), 100)
})

const loadColor = computed(() => {
  if (systemLoad.value < 50) return "#67c23a"
  if (systemLoad.value < 80) return "#e6a23c"
  return "#f56c6c"
})
</script>

<style scoped lang="scss">
.system-status {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
}

.status-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  
  h3 {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: #303133;
  }
}

.status-content {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.status-label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #606266;
  font-size: 14px;
  
  .el-icon {
    color: #909399;
  }
}

.status-value {
  color: #303133;
  font-size: 14px;
  font-weight: 500;
  font-family: monospace;
}

.status-footer {
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid #ebeef5;
}

.load-text {
  text-align: center;
  font-size: 12px;
  color: #909399;
  margin-top: 8px;
}
</style>
