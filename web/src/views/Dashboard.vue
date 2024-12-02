<script setup>
import {ref} from "vue";
import {goodsPageService} from "@/api/goods.js";
import {categoryListService} from "@/api/category.js";
import {orderPageService} from "@/api/order.js";
import {monthRange, todayRange, weekRange, yearRange, yesterdayRange} from "@/utils/time-range.js";

// 获取商品分类信息
const categories = ref([])
const categoryList = async () => {
  let result = await categoryListService()
  categories.value = result.data
}
categoryList()

// 商品总库存量
const totalStock = ref(0.0)

// 获取所有商品
const goods = ref([])
const goodsList = async () => {
  for (let i = 1; true; i++) {
    const data = {
      pageNum: i,
      pageSize: 100
    }
    const result = await goodsPageService(data)
    if(result.data.items && result.data.items.length > 0) {
      for(let idx in result.data.items) {
        const item = result.data.items[idx]
        goods.value.push(item)
        totalStock.value += Number(item.stock)
      }
      if(result.data.items.length < data.pageSize) {
        break
      }
    } else {
      break
    }
  }
}
goodsList()

// 总销售额
const totalSale = ref(0.0)
const todaySale = ref(0.0)
const yesterdaySale = ref(0.0)
const weekSale = ref(0.0)
const monthSale = ref(0.0)
const yearSale = ref(0.0)
// 获取所有订单
const orders = ref([])
const orderList = async () => {
  for (let i=1;true;i++) {
    const data = {
      pageNum: i,
      pageSize: 100
    }
    const result = await orderPageService(data);
    if(result.data.items && result.data.items.length>0) {
      for(let idx in result.data.items) {
        const item = result.data.items[idx]
        orders.value.push(item)
        totalSale.value += Number(item.totalPaid)
        const payTime = new Date(item.payTime)
        if(payTime>=todayRange[0] && payTime <= todayRange[1]) {
          todaySale.value += Number(item.totalPaid)
        }
        if(payTime>=yesterdayRange[0] && payTime<=yesterdayRange[1]) {
          yesterdaySale.value += Number(item.totalPaid)
        }
        if(payTime>=weekRange[0] && payTime<=weekRange[1]) {
          weekSale.value += Number(item.totalPaid)
        }
        if(payTime>=monthRange[0] && payTime<=monthRange[1]) {
          monthSale.value += Number(item.totalPaid)
        }
        if(payTime>=yearRange[0] && payTime<=yearRange[1]) {
          yearSale.value += Number(item.totalPaid)
        }
      }
      if(result.data.items.length < data.pageSize) {
        break
      }
    } else {
      break
    }
  }
}
orderList()


</script>

<template>
  <el-card class="page-container">
    <!--商品统计-->
    <el-card class="goods-container">
      <div class="header">
        商品统计
      </div>
      <el-row :gutter="16">
        <!--分类-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="categories.length">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  商品总分类数
                </div>
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--商品总数-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="goods.length">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  商品总数量
                </div>
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--总库存-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="totalStock" :precision="3">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  商品总库存
                </div>
              </template>
            </el-statistic>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <!--销售统计-->
    <el-card class="order-container">
      <div class="header">
        销售统计
      </div>
      <!--总销售额-->
      <el-row :gutter="16">
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="totalSale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  总销售额
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--今日销售额-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="todaySale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  今日销售额
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--昨日销售额-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="yesterdaySale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  昨日销售额
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--近一周销售额-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="weekSale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  近一周销售额(截至昨日)
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--近一月销售额-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="monthSale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  近一月销售额(截至昨日)
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
        <!--近一年销售额-->
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="yearSale" :precision="2">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  近一年销售额(截至昨日)
                </div>
              </template>
              <template #prefix>
                ￥
              </template>
            </el-statistic>
          </div>
        </el-col>
      </el-row>
    </el-card>
  </el-card>
</template>

<style scoped lang="scss">
.page-container {
  min-height: 100%;
  box-sizing: border-box;
  .goods-container {
    box-sizing: border-box;
    margin-bottom: 20px;
    .header {
      color: red;
      font-size: 20px;
    }
    .el-statistic {
      --el-statistic-content-font-size: 28px;
    }
    .statistic-card {
      height: 100%;
      padding: 16px;
      border-radius: 6px;
      background-color: var(--el-bg-color-overlay);
    }
  }
  .order-container {
    box-sizing: border-box;
    margin-bottom: 20px;
    .header {
      color: red;
      font-size: 20px;
    }
    :global(h2#card-usage ~ .example .example-showcase) {
      background-color: var(--el-fill-color) !important;
    }

    .el-statistic {
      --el-statistic-content-font-size: 28px;
    }

    .statistic-card {
      height: 100%;
      padding: 20px;
      border-radius: 4px;
      background-color: var(--el-bg-color-overlay);
    }
  }
}
</style>