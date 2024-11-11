import { reactive } from "vue";
import { CarInsure } from "../types/car";
import axios from '../axios';
import { ElMessage } from 'element-plus';

export default function(){
    let carInsure = reactive<CarInsure>({
        carID: '',
        name: '',
        cost: 0,
        years: 0,
    })
    
    async function SetCarInsure() {
        const formData = new URLSearchParams();
        formData.append('carID', carInsure.carID as string)
        formData.append('name', carInsure.name)
        formData.append('cost', carInsure.cost.toString())
        formData.append('years', (carInsure.years as number).toString())
        try {
            const response = await axios.post(
                '/SetCarInsure',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarInsure Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarInsure Error: ' + error)
        }
    }

    return {
        carInsure,
        SetCarInsure
    }

}