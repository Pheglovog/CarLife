import { defineStore } from 'pinia';
import { ref} from 'vue';
import { Car } from '../types/car';
import axios from 'axios';
import { ElMessage } from 'element-plus';

export const useCarStore = defineStore('car', ()=>{
    const car = ref<Car | null>(null);
    const carID = ref<string>('');

    const getCar = async () => {
        const params = {
            carID: carID.value
        }
        try {
            const response = await axios.get('/Car', {params});
            car.value = JSON.parse(response.data) as Car;
            ElMessage.success('获取车辆信息成功');
        } catch (error) {
            ElMessage.error('获取车辆信息失败');
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