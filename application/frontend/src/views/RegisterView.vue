<template>
<el-container class="container">
  <el-form
    ref="ruleFormRef"
    style="max-width: 600px"
    :model="user"
    :rules="rules"
    label-width="auto"
    class="demo-ruleForm"
    status-icon
  >
    <el-form-item label="用户名" prop="userID">
      <el-input v-model="user.userID" />
    </el-form-item>
    <el-form-item label="用户类型" prop="userType">
      <el-select v-model="user.userType" placeholder="type">
        <el-option label="组件供应商" value="componentSupplier" />
        <el-option label="制造商" value="manufacturer" />
        <el-option label="销售商" value="store" />
        <el-option label="保险商" value="insurer" />
        <el-option label="维修商" value="maintenancer" />
        <el-option label="消费者" value="consumer" />
      </el-select>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="user.password" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="Register">
        注册
      </el-button>
    </el-form-item>
  </el-form>
</el-container>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import useUser from '../hooks/useUser'

const ruleFormRef = ref<FormInstance>()
let {user, Register} = useUser()

const checkName = (_: any, value: any, callback: any) => {
    if (typeof value !== 'string') {
        return callback(new Error('用户名必须为字符串'))
    }
    if (value === '') {
        return callback(new Error('请输入用户名'))
    }
    if (value.length < 3) {
        return callback(new Error('用户名长度不能小于3'))
    }
    callback()
}
const validatePass = (_: any, value: any, callback: any) => {
  if (value === '') {
    return callback(new Error('请输入密码'))
  } 
  if (value.length < 6 || value.length > 16) {
    return callback(new Error('密码长度不能小于6' + '或大于16'))
  }
  callback()
}
const checkUserType = (_: any, value: any, callback: any) => {
  if (value === '') {
    return callback(new Error('请输入用户类型'))
  }
  callback()
}


const rules = reactive<FormRules<typeof user>>({
    userID: [{ required: true, validator: checkName, trigger: 'blur' }],
    password: [{ required: true, validator: validatePass, trigger: 'blur' }],
    userType: [{ required: true, validator: checkUserType, trigger: 'blur' }]
})
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