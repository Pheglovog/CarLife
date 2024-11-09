<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carInsure"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carInsure.carID" />
        </el-form-item>
        <el-form-item label="保险名" prop="name">
          <el-input v-model="carInsure.name" />
        </el-form-item>
        <el-form-item label="花费" prop="cost">
          <el-input v-model="carInsure.cost" />
        </el-form-item>
        <el-form-item label="年限" prop="years">
          <el-input v-model="carInsure.years" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="SetCarInsure" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import useCarInsure from '../hooks/useCarInsure';
import type { FormRules } from 'element-plus'
let {carInsure, SetCarInsure} = useCarInsure();

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

const rules = reactive<FormRules<typeof carInsure>>({
    carID: [{ required: true, validator: checkString, trigger: 'blur' }],
    name: [{ required: true, validator: checkString, trigger: 'blur' }],
    cost: [{ required: true, validator: checkNumber, trigger: 'blur' }],
    years: [{ required: true, validator: checkNumber, trigger: 'blur' }],
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