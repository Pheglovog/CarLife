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
                <el-col :span="16" class="break-word">{{CarStore.car?.carID}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.owner}}</el-col>
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
                <el-col :span="16" class="break-word">{{formatTime(CarStore.car?.tires.time)}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎宽度：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.tires.width}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">轮胎半径：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.tires.radius}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.tires.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.tires.txID}}</el-col>
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
                <el-col :span="16" class="break-word">{{formatTime(CarStore.car?.body.time)}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身材料：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.body.material}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身重量：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.body.weight}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车身颜色：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.body.color}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.body.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.body.txID}}</el-col>
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
                <el-col :span="16" class="break-word">{{formatTime(CarStore.car?.interior.time)}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰颜色：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.interior.color}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">内饰材料：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.interior.material}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.interior.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.interior.txID}}</el-col>
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
                <el-col :span="16" class="break-word">{{formatTime(CarStore.car?.manu.time)}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">生产车间：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.manu.workshop}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.manu.txID}}</el-col>
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
                <el-col :span="16" class="break-word">{{formatTime(CarStore.car?.store.time)}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">门店：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.store.store}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">费用：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.store.cost}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">车主：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.store.owner}}</el-col>
            </el-row>
            <el-row class="row-bg">
                <el-col :span="8">TXID：</el-col>
                <el-col :span="16" class="break-word">{{CarStore.car?.store.txID}}</el-col>
            </el-row>
        </el-card>
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarInsure</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.insure?.insures" style="width: 100%">
                <el-table-column prop="name" label="Name" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="beginTime" label="BeginTime" :formatter="formatDate" />
                <el-table-column prop="endTime" label="EndTime" :formatter="formatDate"/>
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>     
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>CarMaint</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.maint?.maints" style="width: 100%">
                <el-table-column prop="part" label="Part" />
                <el-table-column prop="extent" label="Extent" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="time" label="Time" :formatter="formatDate"/>
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>  
        <el-card style="max-width: 700px">
            <template #header>
            <div class="card-header">
                <span>TransferRecord</span>
            </div>
            </template>
            <el-table :data="CarStore.car?.record?.records" style="width: 100%">
                <el-table-column prop="oldUser" label="OldUser" />
                <el-table-column prop="newUser" label="NewUser" />
                <el-table-column prop="cost" label="Cost" />
                <el-table-column prop="time" label="Time" :formatter="formatDate"/>
                <el-table-column prop="txID" label="TXID" />
            </el-table>
        </el-card>                                                                           
    </div>
</template>

<script lang="ts" setup>
import {useCarStore} from '../store/car';
import dayjs from 'dayjs';
const CarStore = useCarStore();

function formatTime(time: Date | undefined): string {
  if (time !== undefined) {
    const date = new Date(time);
    // 检查是否为 '0001-01-01T00:00:00Z'
    if (date.getFullYear() === 1 && date.getMonth() === 0 && date.getDate() === 1) {
    return '';
  }
    // 格式化时间
    return dayjs(time).format('YYYY-MM-DD HH:mm');
  }
  return '';
}

const formatDate = (_: any, __: any, cellValue: any) => {
      if (!cellValue) return '';  // 如果时间为空，返回空字符串
      return dayjs(cellValue).format('YYYY-MM-DD HH:mm');  // 格式化时间为 'YYYY-MM-DD HH:mm:ss'
};
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

.break-word {
  word-break: break-all;
  white-space: normal;
}
</style>