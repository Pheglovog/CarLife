<template>
<el-col class="full-left-col">
    <el-menu
    default-active="2"
    class="el-menu-vertical-demo  full-width-menu"
    @open="handleOpen"
    @close="handleClose"
    >
    <el-sub-menu index="1">
        <template #title>
        <el-icon><Ship /></el-icon>
        <span>配件供应</span>
        </template>
        <el-menu-item index="CarTire">车轮</el-menu-item>
        <el-menu-item index="CarBody">车身</el-menu-item>
        <el-menu-item index="CarInterior">内饰</el-menu-item>
    </el-sub-menu>
    <el-menu-item index="Manufacture">
        <el-icon><Van /></el-icon>
        <span>制造</span>
    </el-menu-item>
    <el-menu-item index="Store">
        <el-icon><School /></el-icon>
        <span>销售</span>
    </el-menu-item>
    <el-menu-item index="Maintenancer">
        <el-icon><Setting /></el-icon>
        <span>维修</span>
    </el-menu-item>
    <el-menu-item index="Insure">
        <el-icon><Postcard /></el-icon>
        <span>保险</span>
    </el-menu-item>
    <el-menu-item index="Transfer">
        <el-icon><Switch /></el-icon>
        <span>转卖</span>
    </el-menu-item>       
    </el-menu>
</el-col>
</template>

<script setup lang="ts">
import {Postcard,School,Setting,Ship,Van, Switch} from '@element-plus/icons-vue'
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style>
/* 将 el-col 固定在屏幕的最左侧并让它竖向铺满整个屏幕 */
.full-left-col {
  height: 100%; /* 让 el-col 的高度铺满整个屏幕 */
  background-color: #f5f5f5; /* 可选：设置背景颜色 */
  padding-top: 10px; /* 上方内边距 */
  padding-bottom: 10px; /* 下方内边距 */
}

/* 让菜单占满整个 el-col 容器宽度 */
.full-width-menu {
  width: 100%;
}
</style>