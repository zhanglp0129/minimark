<script type="javascript" setup>
import {ref} from "vue";
import {
  Edit,
  Delete
} from '@element-plus/icons-vue'
import {
  payMethodCreateService,
  payMethodDeleteService,
  payMethodListService,
  payMethodUpdateService
} from '@/api/pay-method.js'
import {ElMessage, ElMessageBox} from "element-plus";

// 获取商品分类信息
const payMethods = ref([])
const payMethodList = async () => {
  let result = await payMethodListService()
  payMethods.value = result.data
}
payMethodList()

// 用于控制对话框
const dialogVisible = ref(false)

// 添加数据模型
const payMethodModel = ref({
  name: ''
})
const payMethodId = ref(0)
// 输入校验
const rules = {
  name: [
    {required: true, message: '请输入支付方式名称', trigger: 'blur'},
  ]
}

// 对话框标题
const title = ref('')

// 打开新增分类的对话框
const addDialog = () => {
  dialogVisible.value=true
  title.value='添加支付方式'
  payMethodModel.value.name = ''
}
// 新增商品分类
const payMethodAdd = async () => {
  await payMethodCreateService(payMethodModel.value)
  ElMessage.success("添加成功")
  // 刷新分类列表
  await payMethodList()
}

// 打开更新分类的对话框
const updateDialog = (row) => {
  dialogVisible.value=true
  title.value='更新分类'
  payMethodModel.value.name = row.name
  payMethodId.value = row.id
}
// 更新分类
const payMethodUpdate = async () => {
  await payMethodUpdateService(payMethodId.value, payMethodModel.value)
  ElMessage.success("更新成功")
  await payMethodList()
}

// 打开删除分类的对话框
const deleteDialog = (row) => {
  ElMessageBox.confirm(
      '确认删除该分类吗？',
      '温馨提示',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async ()=> {
    await payMethodDeleteService(row.id)
    ElMessage.success("删除成功")
    await payMethodList()
  })
}

</script>

<template>
  <el-card class="page-container">
    <template #header>
      <div class="header">
        <span>支付方式</span>
        <div class="extra">
          <el-button type="primary" @click="addDialog">添加支付方式</el-button>
        </div>
      </div>
    </template>
    <el-table :data="payMethods" stripe style="width: 100%">
      <el-table-column prop="id" label="支付方式id"/>
      <el-table-column prop="name" label="支付方式名称" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">   <!--获取当前行数据-->
          <el-button :icon="Edit" circle plain type="primary" @click="updateDialog(row)"></el-button>
          <el-button :icon="Delete" circle plain type="danger" @click="deleteDialog(row)"></el-button>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty description="没有数据" />
      </template>
    </el-table>
  </el-card>

  <!--新增分类或更新分类弹窗-->
  <el-dialog v-model="dialogVisible" :title="title" width="30%">
    <el-form :model="payMethodModel" :rules="rules" label-width="100px" style="padding-right: 30px">
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="payMethodModel.name" minlength="1" maxlength="20"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="title==='添加分类'?payMethodAdd():payMethodUpdate(); dialogVisible = false"> 确认 </el-button>
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
}
</style>