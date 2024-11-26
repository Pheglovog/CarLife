<template>
  <div class="common-layout">
    <el-container>
    <el-aside width="200px">
        <Menu></Menu>
    </el-aside>
    <el-main class="main-container">
        <el-row class="main-row">
        <!-- 左侧区域 -->
            <el-col :span="16" class="left-column">
                <!-- <CarInfo v-if="userStore.isAuthenticated && carStore.carID != ''"></CarInfo>
                <Welcome v-if="!userStore.isAuthenticated || carStore.carID == ''"></Welcome> -->
                <CarInfo></CarInfo>
            </el-col>   
        <!-- 右侧区域（搜索框） -->
            <Search></Search>
        </el-row>
    </el-main>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import Menu from '../components/Menu.vue';
import CarInfo from '../components/CarInfo.vue';
import Search from '../components/Search.vue';
// import { useUserStore } from '../store/user';
// import { useCarStore } from '../store/car';

// const userStore = useUserStore();
// const carStore = useCarStore();
// onMounted(() => {
//   if (userStore.isAuthenticated) {
//     userStore.flashCarList();
//   }
// });
</script>

<style scoped>
/* 让 el-row 的高度撑满整个屏幕 */
.full-height-row {
  height: 100vh; /* 让容器高度占满整个视口 */
  margin: 0;
  padding: 0;
}

/* 让菜单项垂直居中 */
.el-menu-item {
  display: flex;
  align-items: center;
}

.row-bg {
    margin-top: 10px;
}

.main-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  height: 100vh;
}

.main-row {
  width: 100%;
}

.left-column {
    margin-left: auto;
  /* padding-left: 300px; 向右移动内容 */
}
</style>