import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import axios from '../axios';
import { reactive } from 'vue';
import { ElMessage } from 'element-plus';

export const useUserStore = defineStore('user', ()=>{
    const userType = ref<string>('');
    let carList = reactive<string[]>([])
    const token = ref<string | null>(localStorage.getItem('token'));

    const isAuthenticated = computed(() => !!token.value);

    const login = async (formData: URLSearchParams) => {
        try {
            const response = await axios.post(
                '/auth/login',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            token.value = response.data.token;
            localStorage.setItem('token', response.data.token);
            userType.value = response.data.type;
            carList = JSON.parse(response.data.list) as string[];
            ElMessage.success('Login Success, TXID:' + response.data.txid)
        } catch (error) {
            throw new Error(`Login failed! ${error}`);
        }
    }
    const register = async (formData: URLSearchParams) => {
        try {
            const response = await axios.post(
                '/auth/register',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success('Register Success, TXID:' + response.data.txid)
        } catch (error) {
            throw new Error(`Register failed: ${error}`);
        }
    }
    const logout = () => {
        token.value = null;
        userType.value='';
        carList=[];
        localStorage.removeItem('token');
    }

    const flashCarList = async () => {
        try {
            if (isAuthenticated.value) {
                const response = await axios.get('/CarList')
                carList = JSON.parse(response.data) as string[];
            }
        } catch (error) {
            ElMessage.error(`Flash car list failed: ${error}`);
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