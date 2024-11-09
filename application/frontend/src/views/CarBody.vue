<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carBody"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carBody.carID" />
        </el-form-item>
        <el-form-item label="车身材料" prop="material">
          <el-input v-model="carBody.material" />
        </el-form-item>
        <el-form-item label="车身重量" prop="weight">
          <el-input v-model="carBody.weight" />
        </el-form-item>
        <el-form-item label="车身颜色" prop="color">
          <el-input v-model="carBody.color" />
        </el-form-item>
        <el-form-item label="生产车间" prop="workshop">
          <el-input v-model="carBody.workshop" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="SetCarBody" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import type { FormRules } from 'element-plus'
import useCarBody from '../hooks/useCarBody';

let {carBody, SetCarBody} = useCarBody();

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

const rules = reactive<FormRules<typeof carBody>>({
    carID: [{ required: true, validator: checkString, trigger: 'blur' }],
    material: [{ required: true, validator: checkString, trigger: 'blur' }],
    weight: [{ required: true, validator: checkNumber, trigger: 'blur' }],
    color: [{ required: true, validator: checkString, trigger: 'blur' }],
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