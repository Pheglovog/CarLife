import {reactive} from 'vue'
import {User} from '../types/user'
import { ElMessage } from 'element-plus';
import { useUserStore } from '../store/user'
import router from '../router'
export default function() {
    const  userStore = useUserStore()
    let user = reactive<User>({userID: '', userType: '', password: ''})
    async function Register() {
        const formData = new URLSearchParams();
        formData.append('userID', user.userID)
        formData.append('userType', user.userType)
        formData.append('password', user.password)
        try {
            await userStore.register(formData)
            router.push({name: 'login'})
        } catch (error) {
            ElMessage.error('Register Error: ' + error)
        }
    }

    async function Login() {
        const formData = new URLSearchParams();
        formData.append('userID', user.userID)
        formData.append('password', user.password)
        try {
            await userStore.login(formData)
            router.push({name: 'home'})
        } catch (error) {
            ElMessage.error('Login Error: ' + error)
        }
    }

    function Logout() {
        userStore.logout()
        router.push({name: 'login'})
    }

    return {
        user,
        Register,
        Login,
        Logout
    }
}