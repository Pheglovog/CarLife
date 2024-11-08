<template>
    <div class="card-wrapper">
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarID</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">车辆编号：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>                    
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarTire</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">制造时间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎宽度：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎半径：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarBody</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">制造时间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身材料：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身重量：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身颜色：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInterior</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">制造时间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰颜色：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰材料：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
        </el-card>   
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarManu</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">拼装时间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarStore</span>
            </div>
            </template>
            <el-row class="row-bg">
                <el-col :span="8">销售时间：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">门店：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">费用：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="8">{{data}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInsure</span>
            </div>
            </template>
            <el-table style="width: 100%">
                <el-table-column prop="Name" label="Name" />
                <el-table-column prop="Cost" label="Cost" />
                <el-table-column prop="BeginTime" label="BeginTime" />
                <el-table-column prop="EndTime" label="EndTime" />
            </el-table>
        </el-card>     
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInsure</span>
            </div>
            </template>
            <el-table style="width: 100%">
                <el-table-column prop="Part" label="Part" />
                <el-table-column prop="Extent" label="Extent" />
                <el-table-column prop="Cost" label="Cost" />
            </el-table>
        </el-card>  
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInsure</span>
            </div>
            </template>
            <el-table style="width: 100%">
                <el-table-column prop="OldUser" label="OldUser" />
                <el-table-column prop="NewUser" label="NewUser" />
                <el-table-column prop="Cost" label="Cost" />
                <el-table-column prop="Time" label="Time" />
            </el-table>
        </el-card>                                                                           
    </div>
</template>

<script lang="ts" setup>

const data: string="请设置相关值"
</script>

<style scoped>
.card-wrapper {
    display: flex;
    flex-direction: column;
    /* justify-content: center; */
    gap: 20px;
    width: 100%;
    max-width: 720px;
}

.card-header {
  font-weight: bold;
}

.row-bg {
  margin-bottom: 10px; /* 每行之间添加间距 */
}

.el-card {
  padding: 20px; /* 给el-card组件内添加一些内边距 */
  background-color: #f9f9f9; /* 卡片的背景色 */
  border-radius: 10px; /* 设置圆角 */
  margin-top: 20px;
}

.el-col {
  padding: 10px; /* 给每个列添加内边距 */
}
</style>