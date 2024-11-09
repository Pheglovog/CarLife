<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carManu"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carManu.carID" />
        </el-form-item>
        <el-form-item label="车间" prop="workshop">
          <el-input v-model="carManu.workshop" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="SetCarManu" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import useCarManu from '../hooks/useCarManu';
import type { FormRules } from 'element-plus'
let {carManu, SetCarManu} = useCarManu();

const checkString = (rule: any, value: any, callback: any) => {
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

const checkNumber = (rule: any, value: any, callback: any) => {
    if (typeof value !== 'number') {
        return callback(new Error('请输入数字'))
    }
    if (value <= 0) {
        return callback(new Error('请输入大于0的数字'))
    }
    callback()
}

const rules = reactive<FormRules<typeof carManu>>({
    carID: [{ required: true, validator: checkString, trigger: 'blur' }],
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
</style>