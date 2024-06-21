<script lang="js" setup>
import {ref} from "vue";
import {userLoginService} from "@/api/user.js";
import {useTokenStore} from "@/stores/token.js";
import {ElMessage} from "element-plus";
import router from "@/router/index.js";

const loginData = ref({
  username: '',
  password: ''
})

const tokenStore = useTokenStore()
const login = async () => {
  let result = await userLoginService(loginData.value)
  tokenStore.setToken(result.data.token)
  ElMessage.success('登录成功')
  await router.push('/admin')
}
</script>

<template>
  <el-row class="login-page">
    <el-col :span="6" :offset="3" class="form">
      <el-form :model="loginData" label-width="auto" style="max-width: 600px" require-asterisk-position="right">
        <el-form-item label="用户名" required>
          <el-input v-model="loginData.username" />
        </el-form-item>
        <el-form-item label="密码" required>
          <el-input v-model="loginData.password" type="password" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="login" auto-insert-space class="button">登录</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped>
.login-page {
  height: 100vh;
  background-color: #fff;

  .form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    user-select: none;
    margin: 0 auto;

    .button {
      width: 30%;
      margin: 0 auto;
    }

    .flex {
      width: 100%;
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>