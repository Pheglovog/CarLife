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
            car.value = response.data.data;
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
    return {
        car,
        carID,
        getCar,
        cleanCar
    };
})