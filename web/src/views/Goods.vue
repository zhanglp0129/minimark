<script setup>

import {ref} from "vue";
import {categoryDeleteService, categoryListService} from "@/api/category.js";
import {goodsCreateService, goodsDeleteService, goodsPageService, goodsUpdateService} from "@/api/goods.js";
import {Delete, Edit} from "@element-plus/icons-vue";
import {ElMessage, ElMessageBox} from "element-plus";

// 搜索数据
const searchData = ref({
  pageNum: 1,
  pageSize: 10,
  categoryId: null,
  goodsName: null
})


// 获取商品分类列表
const categories = ref([])
const categoryList = async () => {
  let result = await categoryListService()
  categories.value = result.data
}
categoryList()

const total = ref(0)
const goods = ref([]) // 所有展示的商品
// 商品分页查询
const goodsPage = async () => {
  const result = await goodsPageService(searchData.value)
  total.value = result.data.total
  goods.value = result.data.items
}
goodsPage()

//当每页条数发生了变化
const onSizeChange = (size) => {
  searchData.value.pageSize = size
  goodsPage()
}
//当前页码发生变化
const onCurrentChange = (num) => {
  searchData.value.pageNum = num
  goodsPage()
}

// 控制增加修改对话框的显示
const dialogVisible = ref(false)
const title = ref('')

// 商品信息的数据模型
const goodsModel = ref({
  gtin: null,
  name: null,
  categoryId: null,
  price: null,
  chargeType: null,
  stock: null
})
// 参数校验
const rules = ref({
  name: [
    {required: true, message: '请输入商品名称', trigger: 'blur'},
  ],
  price: [
    {required: true, message: '请输入商品价格', trigger: 'blur'},
  ],
  chargeType: [
    {required: true, message: '请选择计价类型', trigger: 'blur'},
  ],
  stock: [
    {required: true, message: '请输入商品库存', trigger: 'blur'},
  ]
})

// 增加商品的对话框
const addDialog = () => {
  dialogVisible.value=true
  title.value='添加商品'
  // 清空数据模型
  goodsModel.value.gtin = null
  goodsModel.value.name = null
  goodsModel.value.categoryId = null
  goodsModel.value.price = null
  goodsModel.value.chargeType = null
  goodsModel.value.stock = null
}
// 新增商品
const goodsAdd = async () => {
  await goodsCreateService(goodsModel.value)
  ElMessage.success("添加成功")
  await goodsPage()
}

// 修改商品的对话框
const goodsId = ref(0)
const updateDialog = (row) => {
  dialogVisible.value=true
  title.value='修改商品'
  // 拷贝数据
  goodsId.value = row.id
  for(let key in goodsModel.value) {
    goodsModel.value[key] = row[key]
  }
}
// 修改商品
const goodsUpdate = async () => {
  await goodsUpdateService(goodsId.value, goodsModel.value)
  ElMessage.success("修改成功")
  await goodsPage()
}

// 删除商品的对话框
const deleteDialog = (row) => {
  ElMessageBox.confirm(
      '确认删除该商品吗？',
      '温馨提示',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async ()=> {
    await goodsDeleteService(row.id)
    ElMessage.success("删除成功")
    await goodsPage()
  })
}

</script>

<template>
  <el-card class="page-container">
    <!--头部-->
    <template #header>
      <div class="header">
        <span>商品管理</span>
        <div class="extra">
          <el-button type="primary" @click="addDialog">添加商品</el-button>
        </div>
      </div>
    </template>

    <!--搜索表单-->
    <el-form :inline="true" class="search-form">
      <el-form-item label="商品分类">
        <el-select placeholder="请选择" v-model="searchData.categoryId" clearable>
          <el-option
              v-for="c in categories"
              :key="c.id"
              :label="c.name"
              :value="c.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="商品名称">
        <el-input v-model="searchData.goodsName" placeholder="请输入" clearable />
      </el-form-item>
      <!--提交表单-->
      <el-form-item>
        <el-button type="primary" @click="goodsPage">搜索</el-button>
        <el-button @click="searchData.categoryId=null;searchData.goodsName=null;goodsPage()">重置</el-button>
      </el-form-item>
    </el-form>

    <!--商品列表-->
    <el-table :data="goods" stripe style="width: 100%">
      <el-table-column prop="id" label="商品id"/>
      <el-table-column prop="gtin" label="GTIN" />
      <el-table-column prop="name" label="商品名称" />
      <el-table-column prop="category.name" label="商品分类" />
      <el-table-column prop="price" label="商品单价" />
      <el-table-column prop="chargeType" label="计价类型" />
      <el-table-column prop="stock" label="库存" />
      <el-table-column label="操作" width="100px">
        <template #default="{ row }">   <!--获取当前行数据-->
          <el-button :icon="Edit" circle plain type="primary" @click="updateDialog(row)"></el-button>
          <el-button :icon="Delete" circle plain type="danger" @click="deleteDialog(row)"></el-button>
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

  <!--对话框-->
  <el-dialog v-model="dialogVisible" :title="title" width="30%">
    <el-form :model="goodsModel" :rules="rules" label-width="120px" style="padding-right: 30px">
      <el-form-item label="GTIN" prop="gtin">
        <el-input v-model="goodsModel.gtin" placeholder="请输入" maxlength="15" clearable></el-input>
      </el-form-item>
      <el-form-item label="商品名称" prop="name">
        <el-input v-model="goodsModel.name" placeholder="请输入" minlength="1" maxlength="50"></el-input>
      </el-form-item>
      <el-form-item label="商品分类">
        <el-select placeholder="请选择" v-model="goodsModel.categoryId" clearable>
          <el-option
              v-for="c in categories"
              :key="c.id"
              :label="c.name"
              :value="c.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="商品单价" prop="price">
        <el-input-number v-model="goodsModel.price" :precision="2" :step="0.01" placeholder="请输入" />
      </el-form-item>
      <el-form-item label="计价类型" prop="chargeType">
        <el-select placeholder="请选择" v-model="goodsModel.chargeType">
          <el-option value="按个数计价"></el-option>
          <el-option value="散装称重"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="商品库存" prop="stock">
        <el-input-number v-model="goodsModel.stock" :precision="3" :step="0.001" placeholder="请输入" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="title==='添加商品'?goodsAdd():goodsUpdate(); dialogVisible = false"> 确认 </el-button>
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
    .el-input {
      --el-input-width: 400px;
    }
  }
}

</style>