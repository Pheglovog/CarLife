import { reactive } from "vue";
import { CarStore } from "../types/car";
import axios from '../axios';
import { ElMessage } from 'element-plus';

export default function(){
    let carStore = reactive<CarStore>({
        carID: '',
        store: '',
        cost: 0,
        ownerID: ''
    })
    
    async function SetCarStore() {
        const formData = new URLSearchParams();
        formData.append('carID', carStore.carID as string)
        formData.append('store', carStore.store)
        formData.append('cost', carStore.cost.toString())
        formData.append('ownerID', carStore.ownerID)
        try {
            const response = await axios.post(
                '/SetCarStore',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarStore Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarStore Error: ' + error)
        }
    }
    
    return {
        carStore,
        SetCarStore
    }
}