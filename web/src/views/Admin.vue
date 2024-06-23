<script type="javascript" setup>
import {ElMessage, ElMessageBox} from "element-plus";
import {useTokenStore} from "@/stores/token.js";
import router from "@/router/index.js";

const tokenStore = useTokenStore()
const logout = () => {
  ElMessageBox.confirm(
      '确认退出登录吗？',
      '温馨提示',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(async () => {
        // 清空pinia中的信息
        tokenStore.removeToken()
        // 跳转到登录界面
        router.push('/login')
        ElMessage.success('成功退出登录')
      })
}
</script>

<template>
  <el-container class="layout-container">
    <!--左侧菜单-->
    <el-aside width="200px">
      <el-menu active-text-color="#ffd04b" background-color="#00A2E8"  text-color="#fff" router>
        <el-menu-item index="/admin/dashboard">
          <span>仪表盘</span>
        </el-menu-item>
        <el-menu-item index="/admin/category">
          <span>商品分类</span>
        </el-menu-item>
        <el-menu-item index="/admin/goods">
          <span>商品管理</span>
        </el-menu-item>
        <el-menu-item index="/admin/pay-method">
          <span>支付方式</span>
        </el-menu-item>
        <el-menu-item index="/admin/order">
          <span>购买记录</span>
        </el-menu-item>
        <el-menu-item index="/admin/procurement">
          <span>进货记录</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <!--右侧主区域-->
    <el-container>
      <!--头部区域-->
      <el-header>
        <div>小型超市商品管理系统</div>
        <el-button type="primary" text @click="logout">
          退出登录
        </el-button>
      </el-header>
      <!-- 中间区域 -->
      <el-main>
        <router-view></router-view>
      </el-main>
      <!-- 底部区域 -->
      <el-footer>© 2024 by 张亮鹏</el-footer>
    </el-container>
  </el-container>
</template>

<style lang="scss" scoped>
.layout-container {
  height: 100vh;

  .el-aside {
    background-color: #00A2E8;

    &__logo {
      height: 120px;
      background: url('@/assets/logo.png') no-repeat center / 120px auto;
    }

    .el-menu {
      border-right: none;
    }
  }

  .el-header {
    background-color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .el-dropdown__box {
      display: flex;
      align-items: center;

      .el-icon {
        color: #999;
        margin-left: 10px;
      }

      &:active,
      &:focus {
        outline: none;
      }
    }
  }

  .el-footer {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    color: #666;
  }
}
</style>