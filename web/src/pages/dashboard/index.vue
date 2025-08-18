<script lang="ts" setup>
import { useUserStore } from "@/pinia/stores/user_n"
import Admin from "./components/Admin.vue"
import Editor from "./components/Editor.vue"

const userStore = useUserStore()

const isAdmin = userStore.userInfo.username === "admin"
</script>

<template>
  <component :is="isAdmin ? Admin : Editor" />
</template>
