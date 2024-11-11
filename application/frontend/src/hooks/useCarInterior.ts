import { reactive } from "vue";
import { CarBody } from "../types/car";
import axios from '../axios';
import { ElMessage } from 'element-plus';

export default function() {
    let carInterior = reactive<CarBody>({
        carID: '',
        material: '',
        weight: 0,
        color: '',
        workshop: ''
    })
    

    async function SetCarInterior() {
        const formData = new URLSearchParams();
        formData.append('carID', carInterior.carID as string)
        formData.append('material', carInterior.material)
        formData.append('weight', carInterior.weight.toString())
        formData.append('color', carInterior.color)
        formData.append('workshop', carInterior.workshop)
        try {
            const response = await axios.post(
                '/SetCarInterior',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarInterior Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarInterior Error: ' + error)
        }
    }

    return {
        carInterior,
        SetCarInterior
    }
}