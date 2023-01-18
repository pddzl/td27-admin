<template>
  <section class="app-main">
    <router-view v-slot="{ Component }">
      <transition name="fade-transform" mode="out-in">
        <!-- <keep-alive> -->
        <component :is="Component" :key="key" />
        <!-- </keep-alive> -->
      </transition>
    </router-view>
  </section>
</template>

<script lang="ts" setup>
import { computed } from "vue"
import { useRoute } from "vue-router"

const route = useRoute()
const key = computed(() => {
  return route.path
})
</script>

<style lang="scss" scoped>
.app-main {
  min-height: calc(100vh - var(--td27-navigationbar-height));
  width: 100%;
  position: relative;
  overflow: hidden;
  background-color: var(--td27-body-bg-color);
}

.fixed-header + .app-main {
  padding-top: var(--td27-navigationbar-height);
  height: 100vh;
  overflow: auto;
}

.hasTagsView {
  .app-main {
    min-height: calc(100vh - var(--td27-header-height));
  }
  .fixed-header + .app-main {
    padding-top: var(--td27-header-height);
  }
}
</style>
