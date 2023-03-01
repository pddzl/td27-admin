<template>
  <div class="app-container">
    <el-card>
      <el-tabs>
        <el-tab-pane label="基本信息">
          <div class="form-container">
            <el-form label-width="120px" :model="userInfoForm">
              <el-form-item label="创建时间">
                <el-input style="width: 400px" v-model="userInfoForm.createdAt" disabled />
              </el-form-item>
              <el-form-item label="用户名" required>
                <el-input style="width: 400px" v-model="userInfoForm.username" />
              </el-form-item>
              <el-form-item label="角色">
                <el-input style="width: 400px" v-model="userInfoForm.role" disabled />
              </el-form-item>
              <el-form-item label="手机号码">
                <el-input style="width: 400px" v-model="userInfoForm.phone" />
              </el-form-item>
              <el-form-item label="邮箱">
                <el-input style="width: 400px" v-model="userInfoForm.email" />
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
        <el-tab-pane label="修改密码">
          <div class="form-container">
            <el-form label-width="120px">
              <el-form-item label="旧密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
              <el-form-item label="新密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
              <el-form-item label="确认密码" required>
                <el-input style="width: 400px" />
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue"
import { formatDateTime } from "@/utils/index"
import { useUserStore } from "@/store/modules/user"

const userStore = useUserStore()

const userInfoForm = reactive({
  id: 0,
  createdAt: "",
  username: "",
  phone: "",
  email: "",
  role: ""
})

const getCache = () => {
  userInfoForm.id = userStore.userInfo.id
  userInfoForm.createdAt = formatDateTime(userStore.userInfo.createdAt)
  userInfoForm.username = userStore.userInfo.username
  userInfoForm.phone = userStore.userInfo.phone
  userInfoForm.email = userStore.userInfo.email
  userInfoForm.role = userStore.userInfo.role
}
getCache()
</script>

<style lang="scss" scoped>
.form-container {
  padding: 20px;
}
</style>
