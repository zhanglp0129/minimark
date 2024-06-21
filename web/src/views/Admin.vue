<script type="javascript" setup>

</script>

<template>
  <el-container class="layout-container">
    <!--左侧菜单-->
    <el-aside width="200px">
      <el-menu active-text-color="#ffd04b" background-color="#232323"  text-color="#fff" router>
        <el-menu-item index="/admin/dashboard">
          <el-icon>
            <Management />
          </el-icon>
          <span>仪表盘</span>
        </el-menu-item>
        <el-menu-item index="/admin/category">
          <el-icon>
            <Management />
          </el-icon>
          <span>商品分类</span>
        </el-menu-item>
        <el-menu-item index="/admin/goods">
          <el-icon>
            <Management />
          </el-icon>
          <span>商品管理</span>
        </el-menu-item>
        <el-menu-item index="/admin/pay-method">
          <el-icon>
            <Management />
          </el-icon>
          <span>支付方式</span>
        </el-menu-item>
        <el-menu-item index="/admin/order">
          <el-icon>
            <Management />
          </el-icon>
          <span>购买商品</span>
        </el-menu-item>
        <el-menu-item index="/admin/procurement">
          <el-icon>
            <Management />
          </el-icon>
          <span>进货商品</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <!--右侧主区域-->
    <el-container>
      <!--头部区域-->
      <el-header>
        <div>小型超市商品管理系统</div>
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
    background-color: #232323;

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