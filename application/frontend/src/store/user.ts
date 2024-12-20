import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import axios from '../axios';
import { ElMessage, ElLoading } from 'element-plus';
import {useCarStore} from '../store/car';

export const useUserStore = defineStore('user', ()=>{
    const carStore = useCarStore()
    const userType = ref<string>('');
    let carList = ref<string[]>([])
    const token = ref<string | null>(localStorage.getItem('token'));

    const isAuthenticated = computed(() => !!token.value);
    let loadingInstance: any = null;
    const login = async (formData: URLSearchParams) => {
        loadingInstance = ElLoading.service({
            lock: true,
            text: '登录中...',
            background: 'rgba(0, 0, 0, 0.7)', // 设置背景透明度
        })
        try {
            const response = await axios.post(
                '/auth/login',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            token.value = response.data.token;
            localStorage.setItem('token', response.data.token);
            userType.value = response.data.type;
            carList.value = response.data.list;
            console.log(carList)
            ElMessage.success('Login Success')
        } catch (error) {
            throw new Error(`Login failed! ${error}`);
        } finally {
            // 提交完成后，移除loading
            if (loadingInstance) {
              loadingInstance.close()
            }
        }
    }
    const register = async (formData: URLSearchParams) => {
        loadingInstance = ElLoading.service({
            lock: true,
            text: '注册中...',
            background: 'rgba(0, 0, 0, 0.7)', // 设置背景透明度
        })
        try {
            const response = await axios.post(
                '/auth/register',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success('Register Success, TXID:' + response.data.txid)
        } catch (error) {
            throw new Error(`Register failed: ${error}`);
        } finally {
            if (loadingInstance) {
                loadingInstance.close()
            }
        }   
    }
    const logout = () => {
        token.value = null;
        userType.value='';
        carList.value=[];
        localStorage.removeItem('token');
        carStore.cleanCar();
    }

    const flashCarList = async () => {
        loadingInstance = ElLoading.service({
            lock: true,
            text: '刷新中...',
            background: 'rgba(0, 0, 0, 0.7)', // 设置背景透明度
        })
        try {
            if (isAuthenticated.value) {
                const response = await axios.get('/CarList')
                carList.value = response.data.data;
            }
        } catch (error) {
            ElMessage.error(`Flash car list failed: ${error}`);
        } finally {
            // 提交完成后，移除loading
            if (loadingInstance) {
              loadingInstance.close()
            }
        }

    }

    return {
        token,
        isAuthenticated,
        login,
        register,
        logout,
        flashCarList,
        carList,
        userType
    };
}
)