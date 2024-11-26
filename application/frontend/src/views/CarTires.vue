<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carTires"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carTires.carID" />
        </el-form-item>
        <el-form-item label="轮胎宽度" prop="width">
          <el-input v-model="carTires.width" />
        </el-form-item>
        <el-form-item label="轮胎半径" prop="radius">
          <el-input v-model="carTires.radius" />
        </el-form-item>
        <el-form-item label="车间" prop="workshop">
          <el-input v-model="carTires.workshop" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import useCarTire from '../hooks/useCarTire';
import { ElLoading ,ElMessage, type FormRules } from 'element-plus'

let {carTires, SetCarTires} = useCarTire();

let loadingInstance: any = null;
const handleSubmit = async () => {
  // 显示全屏loading
  loadingInstance = ElLoading.service({
    lock: true,
    text: '提交中...',
    background: 'rgba(0, 0, 0, 0.7)', // 设置背景透明度
  })

  try {
    await SetCarTires()  // 执行实际的提交操作
  } catch (error) {
    ElMessage.error(error + '')
  } finally {
    // 提交完成后，移除loading
    if (loadingInstance) {
      loadingInstance.close()
    }
  }
}

const checkString = (_: any, value: any, callback: any) => {
    if (typeof value !== 'string') {
        return callback(new Error('请输入string'))
    }
    if (value === '') {
        return callback(new Error('请输入string'))
    }
    if (value.length < 3) {
        return callback(new Error('string长度不能小于3'))
    }
    callback()
}

const checkNumber = (_: any, value: any, callback: any) => {
    const numValue = Number(value)
    if (isNaN(numValue)) {
        return callback(new Error('请输入数字'))
    }
    if (numValue <= 0) {
        return callback(new Error('请输入大于0的数字'))
    }
    callback()
}

const rules = reactive<FormRules<typeof carTires>>({
    carID: [{ required: true, validator: checkString, trigger: 'blur' }],
    width: [{ required: true, validator: checkNumber, trigger: 'blur' }],
    radius: [{ required: true, validator: checkNumber, trigger: 'blur' }],
    workshop: [{ required: true, validator: checkString, trigger: 'blur' }]
});

</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60vh;
  background-color: #f0f2f5;

}

.loading-content {
  text-align: center;
  padding: 30px;
}

.loading-icon {
  font-size: 50px;
  margin-bottom: 20px;
}

.el-dialog {
  width: 200px;
  padding: 20px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
}

.el-dialog .el-dialog__body {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>