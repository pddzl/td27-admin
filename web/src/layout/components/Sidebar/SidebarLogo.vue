<script lang="ts" setup>
const props = defineProps({
  collapse: {
    type: Boolean,
    default: true
  }
})
</script>

<template>
  <div class="sidebar-logo-container" :class="{ collapse: props.collapse }">
    <transition name="sidebar-logo-fade">
      <router-link v-if="props.collapse" key="collapse" to="/">
        <img src="@/assets/layout/td27-logo.png" class="sidebar-logo" />
      </router-link>
      <router-link v-else key="expand" to="/">
        <img src="@/assets/layout/td27.png" class="sidebar-logo-text" />
      </router-link>
    </transition>
  </div>
</template>

<style lang="scss" scoped>
.sidebar-logo-container {
  position: relative;
  width: 100%;
  height: var(--base-header-height);
  line-height: var(--base-header-height);
  background-color: var(--base-sidebarlogo-bg-color);
  text-align: center;
  overflow: hidden;
  .sidebar-logo {
    display: none;
  }
  .sidebar-logo-text {
    height: 100%;
    width: 100%;
    vertical-align: middle;
  }
}

.collapse {
  .sidebar-logo {
    width: 32px;
    height: 32px;
    vertical-align: middle;
    display: inline-block;
  }
  .sidebar-logo-text {
    display: none;
  }
}

.tit-text {
  // display: inline-block;
  color: #fff;
  font-weight: 600;
  font-size: 20px;
  padding-left: 50px;
}
</style>
