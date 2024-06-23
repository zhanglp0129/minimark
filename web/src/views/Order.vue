<script setup>

import {ref} from "vue";
import {
  orderAddGoodsService,
  orderCreateService,
  orderDeleteService,
  orderPageService, orderUpdateGoodsService,
  orderUpdateService
} from "@/api/order.js";
import {payMethodListService} from "@/api/pay-method.js";
import {Delete, Edit, Plus} from "@element-plus/icons-vue";
import {goodsFindService, goodsPageService} from "@/api/goods.js";
import {ElMessage, ElMessageBox, ElSwitch} from "element-plus";

// 获取支付方式
const payMethods = ref([])
const payMethodList = async () => {
  let result = await payMethodListService()
  payMethods.value = result.data
}
payMethodList()

// 获取所有商品
const goods = ref([])
const goodsList = async () => {
  for (let i = 1;; i++) {
    const data = {
      pageNum: i,
      pageSize: 100
    }
    const result = await goodsPageService(data)
    if(result.data.items && result.data.items.length > 0) {
      goods.value.push(...result.data.items)
    } else {
      break
    }
  }
}
goodsList()

// 订单分页查询
const orders = ref([])
const total = ref(0)
const searchData = ref({
  pageNum: 1,
  pageSize: 10,
  payMethodId: null,
  payTimeRange: null,
  totalPayRange: null
})
const clearSearchData = () => {
  searchData.value.payMethodId = null
  searchData.value.payTimeRange = null
  searchData.value.totalPayRange = null
}
const orderPage = async () => {
  const result = await orderPageService(searchData.value)
  orders.value = result.data.items
  total.value = result.data.total
}
orderPage()

//当每页条数发生了变化
const onSizeChange = (size) => {
  searchData.value.pageSize = size
  orderPage()
}
//当前页码发生变化
const onCurrentChange = (num) => {
  searchData.value.pageNum = num
  orderPage()
}

// 查看订单中的商品详情
const goodsInfo = ref([])
const orderGoodsInfo = async (row) => {
  goodsInfo.value = []
  let goodsIds = []
  for (let item of row.goods) {
    goodsInfo.value.push({
      id: item.goodsId,
      price: item.price,
      quantity: item.quantity,
      orderId: item.orderId
    })
    goodsIds.push(item.goodsId)
  }
  const result = await goodsFindService(goodsIds)

  // 设置商品名称
  for(let i=0;i<result.data.length;i++){
    for(let j=0;j<goodsInfo.value.length;j++){
      if(goodsInfo.value[i].id===result.data[i].id){
        goodsInfo.value[i].name = result.data[i].name
        break
      }
    }
  }
}

// 添加订单和编辑订单基本信息
const orderAddUpdateModel = ref({
  totalPaid: null,
  payMethodId: null,
  payTime: null,
  goods: [],
  changeStock: true
})
const orderAddUpdateRules = ref({
  goods: [
    {required: true, message: '请输入商品信息', trigger: 'blur'},
  ]
})
const clearOrderAddUpdateModel = () => {
  orderAddUpdateModel.value.totalPaid = null
  orderAddUpdateModel.value.payMethodId = null
  orderAddUpdateModel.value.payTime = null
  orderAddUpdateModel.value.goods = []
  orderAddUpdateModel.value.changeStock = true
}
const isAdd = ref(false) // 判断是否正在添加订单
const title = ref('')
const addUpdateDialogVisible = ref(false) // 控制对话框的显示
// 添加订单对话框
const addDialog = () => {
  isAdd.value = true
  title.value = '添加购买记录'
  addUpdateDialogVisible.value = true
  clearOrderAddUpdateModel()
  orderAddUpdateModel.value.goods = [{
    goodsId: null,
    price: null,
    quantity: null,
  }]
}
const orderAdd = async () => {
  await orderCreateService(orderAddUpdateModel.value)
  ElMessage.success("添加成功")
  await orderPage()
}
// 修改订单对话框
const orderId = ref(null)
const updateDialog = (row) => {
  isAdd.value = false
  title.value = '修改购买记录'
  addUpdateDialogVisible.value = true
  // 拷贝数据
  orderId.value = row.id
  for(let key in orderAddUpdateModel.value) {
    orderAddUpdateModel.value[key] = row[key]
  }
  orderAddUpdateModel.value.goods = null
}
const orderUpdate = async () => {
  await orderUpdateService(orderId.value, orderAddUpdateModel.value)
  ElMessage.success("修改成功")
  await orderPage()
}

// 订单添加商品
const orderAddGoodsDialogVisible = ref(false)
const orderGoods = ref([])
const addOrderGoodsDialog = (row)=>{
  orderAddGoodsDialogVisible.value = true
  orderGoods.value = [{
    goodsId: null,
    price: null,
    quantity: null
  }]
  title.value = '购买记录添加商品'
  orderId.value = row.id
}
const orderAddGoods = async () => {
  await orderAddGoodsService(orderId.value, orderGoods.value)
  ElMessage.success("添加成功")
  await orderPage()
}

// 订单修改商品
const orderUpdateGoodsDialogVisible = ref(false)
const orderUpdateGoodsModel = ref({
  orderId: null,
  goodsId: null,
  price: null,
  quantity: null,
})
const updateOrderGoodsDialog = (row) => {
  orderUpdateGoodsDialogVisible.value = true
  title.value = '修改购买记录中的商品'
  orderUpdateGoodsModel.value.orderId = row.orderId
  orderUpdateGoodsModel.value.goodsId = row.id
  orderUpdateGoodsModel.value.price = row.price
  orderUpdateGoodsModel.value.quantity = row.quantity
}
const orderUpdateGoods = async () => {
  await orderUpdateGoodsService(orderUpdateGoodsModel.value.orderId, [orderUpdateGoodsModel.value])
  ElMessage.success("修改成功")
  await orderPage()
}

// 订单删除商品
const goodsId = ref(null)
const orderDeleteGoodsDialogVisible = ref(false)
const deleteOrderGoodsDialog = (row) => {
  orderDeleteGoodsDialogVisible.value = true
  title.value = '删除购买记录中的商品'
  changeStock.value = true
  orderUpdateGoodsModel.value.orderId = row.orderId
  orderUpdateGoodsModel.value.goodsId = row.id
  orderUpdateGoodsModel.value.price = 0
  orderUpdateGoodsModel.value.quantity = 0
}
const orderDeleteGoods = async () => {
  await orderUpdateGoodsService(orderUpdateGoodsModel.value.orderId, [orderUpdateGoodsModel.value])
  ElMessage.success("删除成功")
  await orderPage()
}

// 删除订单
const deleteDialogVisible = ref(false)
const changeStock = ref(true)
const deleteDialog = (row) => {
  deleteDialogVisible.value = true
  changeStock.value = true
  title.value = '删除购买记录'
  orderId.value = row.id
}
const deleteOrder = async ()=>{
  await orderDeleteService(orderId.value, changeStock.value)
  ElMessage.success("删除成功")
  await orderPage()
}


</script>

<template>
  <el-card class="page-container">
    <!--头部-->
    <template #header>
      <div class="header">
        <span>购买记录</span>
        <div class="extra">
          <el-button type="primary" @click="addDialog">添加购买记录</el-button>
        </div>
      </div>
    </template>

    <!--搜索表单-->
    <el-form :inline="true" class="search-form">
      <el-form-item label="支付方式">
        <el-select placeholder="请选择" v-model="searchData.payMethodId" clearable>
          <el-option
              v-for="c in payMethods"
              :key="c.id"
              :label="c.name"
              :value="c.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="购买时间">
        <el-date-picker
            v-model="searchData.payTimeRange"
            type="datetimerange"
            unlink-panels
            range-separator="-"
            start-placeholder="开始"
            end-placeholder="结束"
            clearable
            format="YYYY-MM-DD HH:mm:ss"
        />
      </el-form-item>

      <!--提交表单-->
      <el-form-item>
        <el-button type="primary" @click="orderPage">搜索</el-button>
        <el-button @click="clearSearchData();orderPage()">重置</el-button>
      </el-form-item>
    </el-form>

    <!--订单列表-->
    <el-table :data="orders" stripe style="width: 100%">
      <el-table-column prop="id" label="购买记录id"/>
      <el-table-column prop="totalPaid" label="实付金额"/>
      <el-table-column label="商品信息">
        <template #default="{row}">
          <el-popover effect="light" trigger="click" placement="right-start" width="auto">
            <template #default>
              <el-table :data="goodsInfo" stripe style="width: 100%;font-size: smaller">
                <el-table-column prop="name" label="商品名称"/>
                <el-table-column prop="quantity" label="购买数量"/>
                <el-table-column prop="price" label="购买单价"/>
                <el-table-column label="操作" width="100">
                  <template #default="{row}">
                    <el-button :icon="Edit" circle plain type="primary" size="small" @click="updateOrderGoodsDialog(row)"/>
                    <el-button :icon="Delete" circle plain type="danger" size="small" @click="deleteOrderGoodsDialog(row)"/>
                  </template>
                </el-table-column>
              </el-table>
            </template>
            <template #reference>
              <el-tag style="cursor: pointer;" @click="orderGoodsInfo(row)">共{{ row.goods.length }}件</el-tag>
            </template>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column prop="payTime" label="购买时间"/>
      <el-table-column prop="payMethod.name" label="支付方式"/>

      <el-table-column label="操作" width="200px">
        <template #default="{ row }">   <!--获取当前行数据-->
          <el-button :icon="Plus" circle plain type="success" @click="addOrderGoodsDialog(row)"/>
          <el-button :icon="Edit" circle plain type="primary" @click="updateDialog(row)"/>
          <el-button :icon="Delete" circle plain type="danger" @click="deleteDialog(row)"/>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="没有数据" />
      </template>
    </el-table>

    <!--分页条-->
    <el-pagination v-model:current-page="searchData.pageNum" v-model:page-size="searchData.pageSize" :page-sizes="[5, 10, 15, 20]"
                   layout="jumper, total, sizes, prev, pager, next" background :total="total" @size-change="onSizeChange"
                   @current-change="onCurrentChange" style="margin-top: 20px; justify-content: flex-end" />

  </el-card>

  <!--添加和修改订单对话框-->
  <el-dialog v-model="addUpdateDialogVisible" :title="title" width="45%">
    <el-form :model="orderAddUpdateModel" :rules="orderAddUpdateRules" label-width="120px" style="padding-right: 30px">
      <el-form-item label="实付金额" prop="totalPaid">
        <el-input-number v-model="orderAddUpdateModel.totalPaid" :precision="2" :step="0.01" placeholder="请输入" />
      </el-form-item>
      <el-form-item label="支付方式">
        <el-select placeholder="请选择" v-model="orderAddUpdateModel.payMethodId" clearable>
          <el-option
              v-for="p in payMethods"
              :key="p.id"
              :label="p.name"
              :value="p.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="购买时间" prop="payTime">
        <el-date-picker
            v-model="orderAddUpdateModel.payTime"
            type="datetime"
            placeholder="请选择时间"
            clearable
        />
      </el-form-item>
      <el-form-item label="购买商品" prop="goods" v-if="isAdd">
        <el-form inline :model="orderAddUpdateModel.goods">
        <div v-for="(v,i) in orderAddUpdateModel.goods" :key="i" style="margin-bottom: 16px">
          <el-form-item :label="'商品'+(i+1)" :prop="'goods.'+i+'.goodsId'" :rules="[{ required: true, message: '请选择商品',trigger: 'blur'}]">
            <el-select placeholder="请选择" v-model="v.goodsId" style="--el-select-width: 300px;">
              <el-option
                  v-for="g in goods"
                  :key="g.id"
                  :label="g.name"
                  :value="g.id"/>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button v-if="i===0" type="primary" @click="orderAddUpdateModel.goods.push({goodsId: null, price: null, quantity: null})">添加</el-button>
            <el-button v-else type="danger" @click="orderAddUpdateModel.goods.splice(i, 1)">删除</el-button>
          </el-form-item>
          <el-form-item label="数量" :rules="[{ required: true, message: '请输入数量',trigger: 'blur'}]">
            <el-input-number v-model="v.quantity" placeholder="请输入"/>
          </el-form-item>
          <el-form-item label="购买单价" :prop="v.price">
            <el-input-number v-model="v.price" :precision="2" :step="0.01" placeholder="请输入" />
          </el-form-item>
        </div>
        </el-form>
      </el-form-item>
      <el-form-item label="修改库存" prop="changeStock" v-if="isAdd">
        <el-switch v-model="orderAddUpdateModel.changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="addUpdateDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="isAdd?orderAdd():orderUpdate(); addUpdateDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--订单添加商品对话框-->
  <el-dialog v-model="orderAddGoodsDialogVisible" :title="title" width="40%">
    <el-form inline :model="orderGoods">
      <div v-for="(v,i) in orderGoods" :key="i" style="margin-bottom: 16px">
        <el-form-item :label="'商品'+(i+1)" :rules="[{ required: true, message: '请选择商品',trigger: 'blur'}]">
          <el-select placeholder="请选择" v-model="v.goodsId" style="--el-select-width: 300px;">
            <el-option
                v-for="g in goods"
                :key="g.id"
                :label="g.name"
                :value="g.id"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button v-if="i===0" type="primary" @click="orderGoods.push({goodsId: null, price: null, quantity: null})">添加</el-button>
          <el-button v-else type="danger" @click="orderGoods.splice(i, 1)">删除</el-button>
        </el-form-item>
        <el-form-item label="数量" :rules="[{ required: true, message: '请输入数量',trigger: 'blur'}]">
          <el-input-number v-model="v.quantity" placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="购买单价" :prop="v.price">
          <el-input-number v-model="v.price" :precision="2" :step="0.01" placeholder="请输入" />
        </el-form-item>
      </div>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="orderAddGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="orderAddGoods(); orderAddGoodsDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--删除订单对话框-->
  <el-dialog v-model="deleteDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="deleteDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="deleteOrder(); deleteDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--订单修改商品对话框-->
  <el-dialog v-model="orderUpdateGoodsDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="数量" :rules="[{required: true, message: '请输入数量', trigger: 'blur'}]" >
        <el-input-number v-model="orderUpdateGoodsModel.quantity" placeholder="请输入"/>
      </el-form-item>
      <el-form-item label="单价" :rules="[{required: true, message: '请输入单价', trigger: 'blur'}]">
        <el-input-number v-model="orderUpdateGoodsModel.price" :precision="2" :step="0.01" placeholder="请输入" />
      </el-form-item>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="orderUpdateGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="orderUpdateGoods(); orderUpdateGoodsDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--订单删除商品对话框-->
  <el-dialog v-model="orderDeleteGoodsDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="orderDeleteGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="orderDeleteGoods(); orderDeleteGoodsDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

</template>

<style scoped lang="scss">
.page-container {
  min-height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .search-form {
    .el-select {
      --el-select-width: 180px;
    }
  }
}
</style>