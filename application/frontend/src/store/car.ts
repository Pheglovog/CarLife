import { defineStore } from 'pinia';
import { ref} from 'vue';
import { Car } from '../types/car';
import axios from '../axios';
import { ElMessage, ElLoading } from 'element-plus';

export const useCarStore = defineStore('car', ()=>{
    const car = ref<Car | null>(null);
    const carID = ref<string>('');
    let loadingInstance: any = null;

    const getCar = async () => {
        loadingInstance = ElLoading.service({
            lock: true,
            text: '搜索中...',
            background: 'rgba(0, 0, 0, 0.7)', // 设置背景透明度
          })
        try {
            const response = await axios.get('/Car', {
                params:{
                    carID: carID.value
                }
            });
            console.log(response.data);
            car.value = JSON.parse(response.data.data);
            console.log(car.value);
            ElMessage.success('获取车辆信息成功');
        } catch (error) {
            ElMessage.error('获取车辆信息失败');
        } finally {
            // 提交完成后，移除loading
            if (loadingInstance) {
              loadingInstance.close()
            }
        }
    }

    const cleanCar = () => {
        car.value = null;
        carID.value = '';
    }

    const setCar = () => {
        const jsonString = `{
            "carID": "car1",
            "tires": {
                "time": "2024-11-11T16:50:02+08:00",
                "width": 0.2,
                "radius": 0.5,
                "workshop": "tire1",
                "txID": "eacfc43988e8e088c042a97242c045c1821dc1b06e3da10bfa458455ed3d07af"
            },
            "body": {
                "time": "0001-01-01T00:00:00Z",
                "material": "",
                "weitght": 0,
                "color": "",
                "workshop": "",
                "txID": ""
            },
            "interior": {
                "time": "0001-01-01T00:00:00Z",
                "material": "",
                "weitght": 0,
                "color": "",
                "workshop": "",
                "txID": ""
            },
            "manu": {
                "time": "0001-01-01T00:00:00Z",
                "workshop": "",
                "txID": ""
            },
            "store": {
                "time": "0001-01-01T00:00:00Z",
                "store": "",
                "cost": 0,
                "owner": "",
                "txID": ""
            },
            "insure": null,
            "maint": null,
            "owner": "",
            "record": null
        }`;

        car.value = JSON.parse(jsonString);
    }
    return {
        car,
        carID,
        getCar,
        cleanCar,
        setCar
    };
})