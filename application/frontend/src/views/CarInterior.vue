<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carInterior"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carInterior.carID" />
        </el-form-item>
        <el-form-item label="内饰材料" prop="material">
          <el-input v-model="carInterior.material" />
        </el-form-item>
        <el-form-item label="内饰重量" prop="weight">
          <el-input v-model="carInterior.weight" />
        </el-form-item>
        <el-form-item label="内饰颜色" prop="color">
          <el-input v-model="carInterior.color" />
        </el-form-item>
        <el-form-item label="车间" prop="workshop">
          <el-input v-model="carInterior.workshop" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="SetCarInterior" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import useCarInterior from '../hooks/useCarInterior';
import type { FormRules } from 'element-plus'
let {carInterior, SetCarInterior} = useCarInterior();

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

const rules = reactive<FormRules<typeof carInterior>>({
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