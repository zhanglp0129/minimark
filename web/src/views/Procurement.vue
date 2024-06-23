<script setup>

import {ref} from "vue";
import {
  procurementAddGoodsService,
  procurementCreateService,
  procurementDeleteService,
  procurementPageService, procurementUpdateGoodsService,
  procurementUpdateService
} from "@/api/procurement.js";
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

// 进货分页查询
const procurements = ref([])
const total = ref(0)
const searchData = ref({
  pageNum: 1,
  pageSize: 10,
  payMethodId: null,
  payTimeRange: null,
  purchaseTimeRange: null,
  totalPayRange: null
})
const clearSearchData = () => {
  searchData.value.payMethodId = null
  searchData.value.payTimeRange = null
  searchData.value.purchaseTimeRange = null
  searchData.value.totalPayRange = null
}
const procurementPage = async () => {
  const result = await procurementPageService(searchData.value)
  procurements.value = result.data.items
  total.value = result.data.total
}
procurementPage()

//当每页条数发生了变化
const onSizeChange = (size) => {
  searchData.value.pageSize = size
  procurementPage()
}
//当前页码发生变化
const onCurrentChange = (num) => {
  searchData.value.pageNum = num
  procurementPage()
}

// 查看进货中的商品详情
const goodsInfo = ref([])
const procurementGoodsInfo = async (row) => {
  goodsInfo.value = []
  let goodsIds = []
  for (let item of row.goods) {
    goodsInfo.value.push({
      id: item.goodsId,
      price: item.price,
      quantity: item.quantity,
      procurementId: item.procurementId
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

// 添加进货和编辑进货基本信息
const procurementAddUpdateModel = ref({
  totalPaid: null,
  payMethodId: null,
  purchaseTime: null,
  payTime: null,
  goods: [],
  changeStock: true
})
const procurementAddUpdateRules = ref({
  goods: [
    {required: true, message: '请输入商品信息', trigger: 'blur'},
  ]
})
const clearProcurementAddUpdateModel = () => {
  procurementAddUpdateModel.value.totalPaid = null
  procurementAddUpdateModel.value.payMethodId = null
  procurementAddUpdateModel.value.payTime = null
  procurementAddUpdateModel.value.purchaseTime = null
  procurementAddUpdateModel.value.goods = []
  procurementAddUpdateModel.value.changeStock = true
}
const isAdd = ref(false) // 判断是否正在添加进货
const title = ref('')
const addUpdateDialogVisible = ref(false) // 控制对话框的显示
// 添加进货对话框
const addDialog = () => {
  isAdd.value = true
  title.value = '添加进货记录'
  addUpdateDialogVisible.value = true
  clearProcurementAddUpdateModel()
  procurementAddUpdateModel.value.goods = [{
    goodsId: null,
    price: null,
    quantity: null,
  }]
}
const procurementAdd = async () => {
  await procurementCreateService(procurementAddUpdateModel.value)
  ElMessage.success("添加成功")
  await procurementPage()
}
// 修改进货对话框
const procurementId = ref(null)
const updateDialog = (row) => {
  isAdd.value = false
  title.value = '修改进货记录'
  addUpdateDialogVisible.value = true
  // 拷贝数据
  procurementId.value = row.id
  for(let key in procurementAddUpdateModel.value) {
    procurementAddUpdateModel.value[key] = row[key]
  }
  procurementAddUpdateModel.value.goods = null
}
const procurementUpdate = async () => {
  await procurementUpdateService(procurementId.value, procurementAddUpdateModel.value)
  ElMessage.success("修改成功")
  await procurementPage()
}

// 进货添加商品
const procurementAddGoodsDialogVisible = ref(false)
const procurementGoods = ref([])
const addProcurementGoodsDialog = (row)=>{
  procurementAddGoodsDialogVisible.value = true
  procurementGoods.value = [{
    goodsId: null,
    price: null,
    quantity: null
  }]
  title.value = '进货记录添加商品'
  procurementId.value = row.id
}
const procurementAddGoods = async () => {
  await procurementAddGoodsService(procurementId.value, procurementGoods.value)
  ElMessage.success("添加成功")
  await procurementPage()
}

// 进货修改商品
const procurementUpdateGoodsDialogVisible = ref(false)
const procurementUpdateGoodsModel = ref({
  procurementId: null,
  goodsId: null,
  price: null,
  quantity: null,
})
const updateProcurementGoodsDialog = (row) => {
  procurementUpdateGoodsDialogVisible.value = true
  title.value = '修改进货记录中的商品'
  procurementUpdateGoodsModel.value.procurementId = row.procurementId
  procurementUpdateGoodsModel.value.goodsId = row.id
  procurementUpdateGoodsModel.value.price = row.price
  procurementUpdateGoodsModel.value.quantity = row.quantity
}
const procurementUpdateGoods = async () => {
  await procurementUpdateGoodsService(procurementUpdateGoodsModel.value.procurementId, [procurementUpdateGoodsModel.value])
  ElMessage.success("修改成功")
  await procurementPage()
}

// 进货删除商品
const goodsId = ref(null)
const procurementDeleteGoodsDialogVisible = ref(false)
const deleteProcurementGoodsDialog = (row) => {
  procurementDeleteGoodsDialogVisible.value = true
  title.value = '删除进货记录中的商品'
  changeStock.value = true
  procurementUpdateGoodsModel.value.procurementId = row.procurementId
  procurementUpdateGoodsModel.value.goodsId = row.id
  procurementUpdateGoodsModel.value.price = 0
  procurementUpdateGoodsModel.value.quantity = 0
}
const procurementDeleteGoods = async () => {
  await procurementUpdateGoodsService(procurementUpdateGoodsModel.value.procurementId, [procurementUpdateGoodsModel.value])
  ElMessage.success("删除成功")
  await procurementPage()
}

// 删除进货
const deleteDialogVisible = ref(false)
const changeStock = ref(true)
const deleteDialog = (row) => {
  deleteDialogVisible.value = true
  changeStock.value = true
  title.value = '删除进货记录'
  procurementId.value = row.id
}
const deleteProcurement = async ()=>{
  await procurementDeleteService(procurementId.value, changeStock.value)
  ElMessage.success("删除成功")
  await procurementPage()
}


</script>

<template>
  <el-card class="page-container">
    <!--头部-->
    <template #header>
      <div class="header">
        <span>进货记录</span>
        <div class="extra">
          <el-button type="primary" @click="addDialog">添加进货记录</el-button>
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
      <el-form-item label="进货时间">
        <el-date-picker
            v-model="searchData.purchaseTimeRange"
            type="datetimerange"
            unlink-panels
            range-separator="-"
            start-placeholder="开始"
            end-placeholder="结束"
            clearable
            format="YYYY-MM-DD HH:mm:ss"
        />
      </el-form-item>
      <el-form-item label="付款时间">
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
        <el-button type="primary" @click="procurementPage">搜索</el-button>
        <el-button @click="clearSearchData();procurementPage()">重置</el-button>
      </el-form-item>
    </el-form>

    <!--进货列表-->
    <el-table :data="procurements" stripe style="width: 100%">
      <el-table-column prop="id" label="进货记录id"/>
      <el-table-column prop="totalPaid" label="实付金额"/>
      <el-table-column label="商品信息">
        <template #default="{row}">
          <el-popover effect="light" trigger="click" placement="right-start" width="auto">
            <template #default>
              <el-table :data="goodsInfo" stripe style="width: 100%;font-size: smaller">
                <el-table-column prop="name" label="商品名称"/>
                <el-table-column prop="quantity" label="进货数量"/>
                <el-table-column prop="price" label="进货单价"/>
                <el-table-column label="操作" width="100">
                  <template #default="{row}">
                    <el-button :icon="Edit" circle plain type="primary" size="small" @click="updateProcurementGoodsDialog(row)"/>
                    <el-button :icon="Delete" circle plain type="danger" size="small" @click="deleteProcurementGoodsDialog(row)"/>
                  </template>
                </el-table-column>
              </el-table>
            </template>
            <template #reference>
              <el-tag style="cursor: pointer;" @click="procurementGoodsInfo(row)">共{{ row.goods.length }}件</el-tag>
            </template>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column prop="purchaseTime" label="进货时间"/>
      <el-table-column prop="payTime" label="付款时间"/>
      <el-table-column prop="payMethod.name" label="支付方式"/>

      <el-table-column label="操作" width="200px">
        <template #default="{ row }">   <!--获取当前行数据-->
          <el-button :icon="Plus" circle plain type="success" @click="addProcurementGoodsDialog(row)"/>
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

  <!--添加和修改进货对话框-->
  <el-dialog v-model="addUpdateDialogVisible" :title="title" width="45%">
    <el-form :model="procurementAddUpdateModel" :rules="procurementAddUpdateRules" label-width="120px" style="padding-right: 30px">
      <el-form-item label="实付金额" prop="totalPaid">
        <el-input-number v-model="procurementAddUpdateModel.totalPaid" :precision="2" :step="0.01" placeholder="请输入" />
      </el-form-item>
      <el-form-item label="支付方式">
        <el-select placeholder="请选择" v-model="procurementAddUpdateModel.payMethodId" clearable>
          <el-option
              v-for="p in payMethods"
              :key="p.id"
              :label="p.name"
              :value="p.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="进货时间" prop="purchaseTime">
        <el-date-picker
            v-model="procurementAddUpdateModel.purchaseTime"
            type="datetime"
            placeholder="请选择时间"
            clearable
        />
      </el-form-item>
      <el-form-item label="付款时间" prop="payTime">
        <el-date-picker
            v-model="procurementAddUpdateModel.payTime"
            type="datetime"
            placeholder="请选择时间"
            clearable
        />
      </el-form-item>
      <el-form-item label="进货商品" prop="goods" v-if="isAdd">
        <el-form inline :model="procurementAddUpdateModel.goods">
          <div v-for="(v,i) in procurementAddUpdateModel.goods" :key="i" style="margin-bottom: 16px">
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
              <el-button v-if="i===0" type="primary" @click="procurementAddUpdateModel.goods.push({goodsId: null, price: null, quantity: null})">添加</el-button>
              <el-button v-else type="danger" @click="procurementAddUpdateModel.goods.splice(i, 1)">删除</el-button>
            </el-form-item>
            <el-form-item label="数量" :rules="[{ required: true, message: '请输入数量',trigger: 'blur'}]">
              <el-input-number v-model="v.quantity" placeholder="请输入"/>
            </el-form-item>
            <el-form-item label="进货单价" :prop="v.price" :rules="[{ required: true, message: '请输入进货单价',trigger: 'blur'}]">
              <el-input-number v-model="v.price" :precision="2" :step="0.01" placeholder="请输入" />
            </el-form-item>
          </div>
        </el-form>
      </el-form-item>
      <el-form-item label="修改库存" prop="changeStock" v-if="isAdd">
        <el-switch v-model="procurementAddUpdateModel.changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="addUpdateDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="isAdd?procurementAdd():procurementUpdate(); addUpdateDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--进货添加商品对话框-->
  <el-dialog v-model="procurementAddGoodsDialogVisible" :title="title" width="40%">
    <el-form inline :model="procurementGoods">
      <div v-for="(v,i) in procurementGoods" :key="i" style="margin-bottom: 16px">
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
          <el-button v-if="i===0" type="primary" @click="procurementGoods.push({goodsId: null, price: null, quantity: null})">添加</el-button>
          <el-button v-else type="danger" @click="procurementGoods.splice(i, 1)">删除</el-button>
        </el-form-item>
        <el-form-item label="数量" :rules="[{ required: true, message: '请输入数量',trigger: 'blur'}]">
          <el-input-number v-model="v.quantity" placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="进货单价" :prop="v.price" :rules="[{ required: true, message: '请输入进货单价',trigger: 'blur'}]">
          <el-input-number v-model="v.price" :precision="2" :step="0.01" placeholder="请输入" />
        </el-form-item>
      </div>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="procurementAddGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="procurementAddGoods(); procurementAddGoodsDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--删除进货对话框-->
  <el-dialog v-model="deleteDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="deleteDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="deleteProcurement(); deleteDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--进货修改商品对话框-->
  <el-dialog v-model="procurementUpdateGoodsDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="数量" :rules="[{required: true, message: '请输入数量', trigger: 'blur'}]" >
        <el-input-number v-model="procurementUpdateGoodsModel.quantity" placeholder="请输入"/>
      </el-form-item>
      <el-form-item label="单价" :rules="[{required: true, message: '请输入单价', trigger: 'blur'}]">
        <el-input-number v-model="procurementUpdateGoodsModel.price" :precision="2" :step="0.01" placeholder="请输入" />
      </el-form-item>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="procurementUpdateGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="procurementUpdateGoods(); procurementUpdateGoodsDialogVisible = false"> 确认 </el-button>
        </span>
    </template>
  </el-dialog>

  <!--进货删除商品对话框-->
  <el-dialog v-model="procurementDeleteGoodsDialogVisible" :title="title" width="25%">
    <el-form>
      <el-form-item label="修改库存">
        <el-switch v-model="changeStock" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="procurementDeleteGoodsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="procurementDeleteGoods(); procurementDeleteGoodsDialogVisible = false"> 确认 </el-button>
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