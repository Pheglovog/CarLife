<template>
    <el-container class="container">
      <el-form
        style="max-width: 600px"
        :model="carStore"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
        status-icon
      >
        <el-form-item label="carID" prop="carID">
          <el-input v-model="carStore.carID" />
        </el-form-item>
        <el-form-item label="销售点" prop="store">
          <el-input v-model="carStore.store" />
        </el-form-item>
        <el-form-item label="花费" prop="cost">
          <el-input v-model="carStore.cost" />
        </el-form-item>
        <el-form-item label="车主" prop="ownerID">
          <el-input v-model="carStore.ownerID" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="SetCarStore" style="margin-left: 125px;">
            提交
          </el-button>
        </el-form-item>
      </el-form>
    </el-container>
</template>

<script setup lang="ts">
import { reactive} from 'vue'
import useCarStore from '../hooks/useCarStore';
import type { FormRules } from 'element-plus'
let {carStore, SetCarStore} = useCarStore();

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

const rules = reactive<FormRules<typeof carStore>>({
    carID: [{ required: true, validator: checkString, trigger: 'blur' }],
    store: [{ required: true, validator: checkString, trigger: 'blur' }],
    cost: [{ required: true, validator: checkNumber, trigger: 'blur' }],
    ownerID: [{ required: true, validator: checkString, trigger: 'blur' }],
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