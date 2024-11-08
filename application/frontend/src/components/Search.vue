<template>
<el-col :span="6">
    <div class="search-container">
        <el-button type="primary" round>搜索</el-button>
        <el-input placeholder="请输入搜索内容" class="search-input" />
    </div>
    <el-card style="max-width: 700px" >
        <template #header>
        <div class="card-header">
            <span>有关车辆</span>
        </div>
        </template>
        <el-row>车辆id列表</el-row>
    </el-card>
</el-col>
</template>

<script setup lang="ts">

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