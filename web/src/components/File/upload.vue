<template>
    <el-upload
      :action="`${path}/file/upload`"
      :headers="{ 'x-token': useUserStoreHook().token }"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button type="primary" plain icon="upload">上传</el-button>
    </el-upload>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { ElMessage } from "element-plus"
import { useUserStoreHook } from "@/store/modules/user"

const path = ref(import.meta.env.VITE_BASE_API)

const filename = ref("")
const uploadSuccess = (res: ApiResponseData<{ path: string }>) => {
  filename.value = res.data.path
  ElMessage.success("上传成功")
}

const uploadError = () => {
  ElMessage({
    type: "error",
    message: "上传失败"
  })
}
</script>
