<script setup lang="ts">
const props = defineProps<{
  modelValue: string
  type: "second" | "minute" | "hour" | "day" | "month" | "week" | "year"
}>()

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void
  (e: "change"): void
}>()

// Type config
const typeConfig: Record<string, { label: string, min: number, max: number }> = {
  second: { label: "秒", min: 0, max: 59 },
  minute: { label: "分钟", min: 0, max: 59 },
  hour: { label: "小时", min: 0, max: 23 },
  day: { label: "日", min: 1, max: 31 },
  month: { label: "月", min: 1, max: 12 },
  week: { label: "周", min: 1, max: 7 },
  year: { label: "年", min: 2024, max: 2100 }
}

const typeLabel = computed(() => typeConfig[props.type].label)
const min = computed(() => typeConfig[props.type].min)
const max = computed(() => typeConfig[props.type].max)

// Generate options for specific selection
const options = computed(() => {
  const opts = []
  for (let i = min.value; i <= max.value; i++) {
    let label = String(i)
    if (props.type === "week") {
      const weekDays = ["", "周日", "周一", "周二", "周三", "周四", "周五", "周六"]
      label = weekDays[i] || String(i)
    }
    opts.push({ value: i, label })
  }
  return opts
})

// Radio value
const radioValue = ref("*")

// Cycle config
const cycleStart = ref(min.value)
const cycleEnd = ref(max.value)

// Interval config
const intervalStart = ref(min.value)
const intervalStep = ref(1)

// Specific values
const specificValues = ref<number[]>([])

// Parse modelValue
function parseValue(val: string) {
  if (val === "*") {
    radioValue.value = "*"
  } else if (val === "?") {
    radioValue.value = "?"
  } else if (val.includes("-")) {
    radioValue.value = "cycle"
    const [start, end] = val.split("-").map(Number)
    cycleStart.value = start
    cycleEnd.value = end
  } else if (val.includes("/")) {
    radioValue.value = "interval"
    const [start, step] = val.split("/").map(Number)
    intervalStart.value = start
    intervalStep.value = step
  } else if (val.includes(",")) {
    radioValue.value = "specific"
    specificValues.value = val.split(",").map(Number)
  } else if (val) {
    radioValue.value = "specific"
    specificValues.value = [Number(val)]
  }
}

watch(() => props.modelValue, val => parseValue(val), { immediate: true })

// Build value
function buildValue() {
  switch (radioValue.value) {
    case "*":
      return "*"
    case "?":
      return "?"
    case "cycle":
      return `${cycleStart.value}-${cycleEnd.value}`
    case "interval":
      return `${intervalStart.value}/${intervalStep.value}`
    case "specific":
      return specificValues.value.sort((a, b) => a - b).join(",")
    default:
      return "*"
  }
}

// Watch changes
watch([radioValue, cycleStart, cycleEnd, intervalStart, intervalStep, specificValues], () => {
  const newValue = buildValue()
  emit("update:modelValue", newValue)
  emit("change")
}, { deep: true })

function onRadioChange() {
  // Reset values when radio changes
  if (radioValue.value === "specific" && specificValues.value.length === 0) {
    specificValues.value = [min.value]
  }
}
</script>

<template>
  <div class="cron-tab">
    <el-radio-group v-model="radioValue" @change="onRadioChange">
      <el-radio :value="'*'">
        每{{ typeLabel }}
      </el-radio>
      <el-radio :value="'?'">
        不指定
      </el-radio>
      <el-radio :value="'cycle'">
        周期
      </el-radio>
      <el-radio :value="'interval'">
        间隔
      </el-radio>
      <el-radio :value="'specific'">
        指定
      </el-radio>
    </el-radio-group>

    <div class="config-panel">
      <!-- Cycle -->
      <div v-if="radioValue === 'cycle'" class="config-item">
        <span>从</span>
        <el-input-number v-model="cycleStart" :min="min" :max="max" size="small" />
        <span>到</span>
        <el-input-number v-model="cycleEnd" :min="min" :max="max" size="small" />
        <span>{{ typeLabel }}</span>
      </div>

      <!-- Interval -->
      <div v-if="radioValue === 'interval'" class="config-item">
        <span>从</span>
        <el-input-number v-model="intervalStart" :min="min" :max="max" size="small" />
        <span>{{ typeLabel }}开始，每</span>
        <el-input-number v-model="intervalStep" :min="1" :max="max" size="small" />
        <span>{{ typeLabel }}执行一次</span>
      </div>

      <!-- Specific -->
      <div v-if="radioValue === 'specific'" class="config-item">
        <span>指定{{ typeLabel }}:</span>
        <el-select-v2
          v-model="specificValues"
          :options="options"
          placeholder="请选择"
          multiple
          collapse-tags
          collapse-tags-tooltip
          style="width: 300px"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.cron-tab {
  padding: 15px;

  .el-radio-group {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .config-panel {
    margin-top: 20px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 4px;

    .config-item {
      display: flex;
      align-items: center;
      gap: 10px;
      flex-wrap: wrap;

      span {
        color: #666;
      }
    }
  }
}
</style>
