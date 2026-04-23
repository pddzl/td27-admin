<script setup lang="ts">
import { CopyDocument } from "@element-plus/icons-vue"
import CronTab from "./CronTab.vue"

const props = defineProps<{
  modelValue: string
  showSecond?: boolean
  showYear?: boolean
}>()

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void
}>()

const activeTab = ref("minute")

// Cron parts
const second = ref("*")
const minute = ref("*")
const hour = ref("*")
const day = ref("*")
const month = ref("*")
const week = ref("?")
const year = ref("*")

// Parse initial value
function parseExpression(expr: string) {
  if (!expr) return
  const parts = expr.split(" ")
  if (parts.length >= 5) {
    if (props.showSecond && parts.length >= 6) {
      second.value = parts[0]
      minute.value = parts[1]
      hour.value = parts[2]
      day.value = parts[3]
      month.value = parts[4]
      week.value = parts[5] || "?"
      year.value = parts[6] || "*"
    } else {
      minute.value = parts[0]
      hour.value = parts[1]
      day.value = parts[2]
      month.value = parts[3]
      week.value = parts[4] || "?"
    }
  }
}

watch(() => props.modelValue, val => parseExpression(val), { immediate: true })

// Build expression
const cronExpression = computed(() => {
  if (props.showSecond) {
    return `${second.value} ${minute.value} ${hour.value} ${day.value} ${month.value} ${week.value}`
  }
  return `${minute.value} ${hour.value} ${day.value} ${month.value} ${week.value}`
})

function onChange() {
  emit("update:modelValue", cronExpression.value)
}

// Copy expression
function copyExpression() {
  navigator.clipboard.writeText(cronExpression.value)
  ElMessage.success("已复制到剪贴板")
}

// Quick options
const quickOptions = [
  { label: "每分钟", value: "0 * * * * ?" },
  { label: "每小时", value: "0 0 * * * ?" },
  { label: "每天0点", value: "0 0 0 * * ?" },
  { label: "每天12点", value: "0 0 12 * * ?" },
  { label: "每周一", value: "0 0 0 ? * 2" },
  { label: "每月1日", value: "0 0 0 1 * ?" },
  { label: "工作日", value: "0 0 0 ? * 2-6" }
]

function applyQuick(value: string) {
  parseExpression(value)
  onChange()
}

// Calculate next 5 execution times
const nextTimes = computed(() => {
  const times: string[] = []
  try {
    // Simple calculation for demo (in real app, use a cron parser library)
    const now = new Date()
    for (let i = 1; i <= 5; i++) {
      const next = new Date(now.getTime() + i * 60 * 60 * 1000)
      times.push(next.toLocaleString())
    }
  } catch (e) {
    // Invalid cron
  }
  return times
})
</script>

<template>
  <div class="cron-builder">
    <el-tabs v-model="activeTab" type="border-card">
      <el-tab-pane label="秒" name="second" v-if="showSecond">
        <CronTab v-model="second" type="second" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="分钟" name="minute">
        <CronTab v-model="minute" type="minute" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="小时" name="hour">
        <CronTab v-model="hour" type="hour" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="日" name="day">
        <CronTab v-model="day" type="day" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="月" name="month">
        <CronTab v-model="month" type="month" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="周" name="week">
        <CronTab v-model="week" type="week" @change="onChange" />
      </el-tab-pane>
      <el-tab-pane label="年" name="year" v-if="showYear">
        <CronTab v-model="year" type="year" @change="onChange" />
      </el-tab-pane>
    </el-tabs>

    <div class="cron-preview">
      <div class="preview-item">
        <label>Cron表达式:</label>
        <el-input v-model="cronExpression" readonly class="expression-input">
          <template #append>
            <el-button :icon="CopyDocument" @click="copyExpression">
              复制
            </el-button>
          </template>
        </el-input>
      </div>
      <div class="preview-item">
        <label>执行时间预览:</label>
        <div class="next-times">
          <el-tag v-for="(time, index) in nextTimes" :key="index" class="time-tag">
            {{ time }}
          </el-tag>
        </div>
      </div>
    </div>

    <div class="quick-select">
      <label>快捷选择:</label>
      <el-space wrap>
        <el-button v-for="item in quickOptions" :key="item.value" size="small" @click="applyQuick(item.value)">
          {{ item.label }}
        </el-button>
      </el-space>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.cron-builder {
  .cron-preview {
    margin-top: 20px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 4px;

    .preview-item {
      margin-bottom: 15px;

      &:last-child {
        margin-bottom: 0;
      }

      label {
        display: block;
        margin-bottom: 8px;
        font-weight: 500;
        color: #666;
      }

      .expression-input {
        width: 300px;
      }

      .next-times {
        display: flex;
        flex-wrap: wrap;
        gap: 8px;

        .time-tag {
          margin: 0;
        }
      }
    }
  }

  .quick-select {
    margin-top: 20px;
    padding-top: 20px;
    border-top: 1px solid #e4e7ed;

    label {
      display: block;
      margin-bottom: 10px;
      font-weight: 500;
      color: #666;
    }
  }
}
</style>
