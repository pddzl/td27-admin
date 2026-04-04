<template>
  <div class="stats-card" :class="{ 'is-hoverable': hoverable }">
    <div class="stats-icon" :style="{ backgroundColor: iconBgColor, color: iconColor }">
      <el-icon :size="28">
        <component :is="icon" />
      </el-icon>
    </div>
    <div class="stats-content">
      <div class="stats-value">{{ value }}</div>
      <div class="stats-label">{{ label }}</div>
      <div v-if="trend !== undefined" class="stats-trend" :class="trend > 0 ? 'up' : 'down'">
        <el-icon :size="12">
          <CaretTop v-if="trend > 0" />
          <CaretBottom v-else />
        </el-icon>
        <span>{{ Math.abs(trend) }}%</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue"
import { CaretTop, CaretBottom } from "@element-plus/icons-vue"

interface Props {
  icon: any
  value: number | string
  label: string
  iconColor?: string
  iconBgColor?: string
  trend?: number
  hoverable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  iconColor: "#409eff",
  iconBgColor: "#ecf5ff",
  hoverable: true
})
</script>

<style scoped lang="scss">
.stats-card {
  display: flex;
  align-items: center;
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  transition: all 0.3s;
  
  &.is-hoverable:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
}

.stats-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  border-radius: 12px;
  margin-right: 16px;
}

.stats-content {
  flex: 1;
}

.stats-value {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  line-height: 1.2;
}

.stats-label {
  font-size: 13px;
  color: #909399;
  margin-top: 4px;
}

.stats-trend {
  display: inline-flex;
  align-items: center;
  gap: 2px;
  font-size: 12px;
  margin-top: 6px;
  padding: 2px 6px;
  border-radius: 4px;
  
  &.up {
    color: #67c23a;
    background-color: #f0f9eb;
  }
  
  &.down {
    color: #f56c6c;
    background-color: #fef0f0;
  }
}
</style>
