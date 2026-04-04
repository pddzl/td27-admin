<template>
  <div class="welcome-banner">
    <div class="welcome-content">
      <div class="welcome-text">
        <h1>{{ greeting }}，{{ username }}！</h1>
        <p>{{ welcomeMessage }}</p>
      </div>
      <div class="welcome-time">
        <div class="time">{{ currentTime }}</div>
        <div class="date">{{ currentDate }}</div>
      </div>
    </div>
    <div class="welcome-decoration">
      <div class="decoration-circle circle-1"></div>
      <div class="decoration-circle circle-2"></div>
      <div class="decoration-circle circle-3"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from "vue"

const props = defineProps<{
  username: string
}>()

const now = ref(new Date())
let timer: number | null = null

const greeting = computed(() => {
  const hour = now.value.getHours()
  if (hour < 6) return "夜深了"
  if (hour < 9) return "早上好"
  if (hour < 12) return "上午好"
  if (hour < 14) return "中午好"
  if (hour < 18) return "下午好"
  return "晚上好"
})

const welcomeMessage = computed(() => {
  const messages = [
    "欢迎使用 TD27 Admin 管理系统",
    "今天是美好的一天，开始工作吧",
    "保持专注，高效完成任务",
    "有任何问题可以随时联系管理员"
  ]
  const day = now.value.getDay()
  return messages[day % messages.length]
})

const currentTime = computed(() => {
  return now.value.toLocaleTimeString("zh-CN", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit"
  })
})

const currentDate = computed(() => {
  return now.value.toLocaleDateString("zh-CN", {
    year: "numeric",
    month: "long",
    day: "numeric",
    weekday: "long"
  })
})

onMounted(() => {
  timer = window.setInterval(() => {
    now.value = new Date()
  }, 1000)
})

onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
  }
})
</script>

<style scoped lang="scss">
.welcome-banner {
  position: relative;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 32px;
  color: #fff;
  overflow: hidden;
}

.welcome-content {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.welcome-text {
  h1 {
    margin: 0 0 8px 0;
    font-size: 24px;
    font-weight: 600;
  }
  
  p {
    margin: 0;
    font-size: 14px;
    opacity: 0.9;
  }
}

.welcome-time {
  text-align: right;
  
  .time {
    font-size: 32px;
    font-weight: 300;
    font-family: "Roboto Mono", monospace;
    letter-spacing: 2px;
  }
  
  .date {
    font-size: 14px;
    opacity: 0.9;
    margin-top: 4px;
  }
}

.welcome-decoration {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
}

.decoration-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  
  &.circle-1 {
    width: 200px;
    height: 200px;
    top: -60px;
    right: 10%;
  }
  
  &.circle-2 {
    width: 120px;
    height: 120px;
    bottom: -30px;
    right: 25%;
    background: rgba(255, 255, 255, 0.05);
  }
  
  &.circle-3 {
    width: 80px;
    height: 80px;
    top: 20%;
    right: 5%;
    background: rgba(255, 255, 255, 0.08);
  }
}

@media (max-width: 768px) {
  .welcome-content {
    flex-direction: column;
    text-align: center;
    gap: 20px;
  }
  
  .welcome-time {
    text-align: center;
  }
}
</style>
