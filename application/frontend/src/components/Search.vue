<template>
<el-col :span="6">
    <div class="search-container">
        <el-button type="primary" round @click="search">搜索</el-button>
        <!-- <el-button type="primary" round @click="carStore.setCar">搜索</el-button> -->
        <el-input placeholder="请输入搜索内容" class="search-input" v-model="carStore.carID"/>
    </div>
    <el-card style="max-width: 700px" >
        <template #header>
        <div class="card-header">
            <el-button type="primary" round @click="flashCarList">刷新</el-button>
            <span>有关车辆</span>
        </div>
        </template>
        <el-row v-for="(carId, index) in userStore.carList" :key="index">{{ carId }}</el-row>
    </el-card>
</el-col>
</template>

<script setup lang="ts">
import { useCarStore } from '../store/car';
import { useUserStore } from '../store/user';
let carStore = useCarStore();
let userStore = useUserStore();

const flashCarList = () => {
    if (userStore.isAuthenticated) {
    userStore.flashCarList();
  }
}
const search = () => {
    carStore.getCar();
}
</script>

<style>
.search-container {
  display: flex; /* 使用 flex 布局 */
  align-items: center; /* 垂直居中对齐 */
  margin-bottom: 20px;
}
.search-input {
    margin-left: 10px;
}
</style>