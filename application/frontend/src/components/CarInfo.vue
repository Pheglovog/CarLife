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
                <el-col :span="8">{{CarStore.car?.carID}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="8">{{CarStore.car?.owner}}</el-col>
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
                <el-col :span="8">{{CarStore.car?.tires.time}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎宽度：</el-col>
                <el-col :span="8">{{CarStore.car?.tires.width}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎半径：</el-col>
                <el-col :span="8">{{CarStore.car?.tires.radius}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{CarStore.car?.tires.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="8">{{CarStore.car?.tires.txID}}</el-col>
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
                <el-col :span="8">{{CarStore.car?.body.time}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身材料：</el-col>
                <el-col :span="8">{{CarStore.car?.body.material}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身重量：</el-col>
                <el-col :span="8">{{CarStore.car?.body.weight}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身颜色：</el-col>
                <el-col :span="8">{{CarStore.car?.body.color}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{CarStore.car?.body.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="8">{{CarStore.car?.body.txID}}</el-col>
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
                <el-col :span="8">{{CarStore.car?.interior.time}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰颜色：</el-col>
                <el-col :span="8">{{CarStore.car?.interior.color}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰材料：</el-col>
                <el-col :span="8">{{CarStore.car?.interior.material}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{CarStore.car?.interior.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="8">{{CarStore.car?.interior.txID}}</el-col>
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
                <el-col :span="8">{{CarStore.car?.manu.time}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="8">{{CarStore.car?.manu.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="8">{{CarStore.car?.manu.txID}}</el-col>
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
                <el-col :span="8">{{CarStore.car?.store.time}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">门店：</el-col>
                <el-col :span="8">{{CarStore.car?.store.store}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">费用：</el-col>
                <el-col :span="8">{{CarStore.car?.store.cost}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="8">{{CarStore.car?.store.ownerID}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="8">{{CarStore.car?.store.txID}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInsure</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.insure" style="width: 100%">
                <el-table-column prop="name" label="Name" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="beginTime" label="BeginTime" />
                <el-table-column prop="endTime" label="EndTime" />
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>     
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarMaint</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.maint" style="width: 100%">
                <el-table-column prop="part" label="Part" />
                <el-table-column prop="extent" label="Extent" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="time" label="Time" />
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>  
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>TransferRecord</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.record" style="width: 100%">
                <el-table-column prop="OldUser" label="OldUser" />
                <el-table-column prop="NewUser" label="NewUser" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="time" label="Time" />
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>                                                                           
    </div>
</template>

<script lang="ts" setup>
import {useCarStore} from '../store/car';
const CarStore = useCarStore();
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