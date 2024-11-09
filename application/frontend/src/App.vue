<template>
    <el-container>
      <div class="header">
        <h1>CarLife</h1>
      </div>
      <el-header>
      <el-menu
        class="el-menu-demo custom-menu"
        mode="horizontal"
      >
        <el-menu-item index="Home">
          <RouterLink :to="{name: 'home'}" class="no-underline">CarLife</RouterLink>
        </el-menu-item>
        <el-menu-item index="Explorer">
          <a href="http://127.0.0.1:3030" class="no-underline">Explorer</a> 
        </el-menu-item>
        <el-sub-menu index="auth">
          <template #title>用户管理</template>
          <RouterLink :to="{name: 'login'}" class="no-underline" v-if="!userStore.isAuthenticated">
            <el-menu-item index="Login">登录</el-menu-item>
          </RouterLink>
          <RouterLink :to="{name: 'register'}" class="no-underline" v-if="!userStore.isAuthenticated">
            <el-menu-item index="Register">注册</el-menu-item>
          </RouterLink>
          <el-menu-item index="Logout" @click="Logout" v-if="userStore.isAuthenticated">退出</el-menu-item>
        </el-sub-menu>
      </el-menu>
      </el-header>
      <el-main>
        <RouterView></RouterView>
      </el-main>
    </el-container>

</template>

<script setup lang="ts">
  import { RouterLink, RouterView } from 'vue-router';
  import useUser from './hooks/useUser';
  import { useUserStore } from './store/user';
  const userStore = useUserStore();
  let {Logout} = useUser();
</script>

<style>
.custom-menu {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  width: 100%;
}

.el-menu-demo {
  width: 100%;
}

.el-menu-item {
  white-space: nowrap;
}

.header {
  text-align: center; /* 水平居中 */
  margin-top: 20px;    /* 可选：设置顶部距离 */
  padding-right: 20px;  /* 向左移动，增加一些左边距 */
}

h1 {
  margin: 0; /* 去掉默认的 h1 外边距，防止其被浏览器的默认样式影响 */
  background-color: rgb(9, 134, 236); /* 设置背景颜色为黑色 */
  color: white;          /* 设置文本颜色为白色 */
}
.no-underline{
  text-decoration: none;
}
</style>
