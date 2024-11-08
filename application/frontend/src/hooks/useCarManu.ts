import { reactive } from "vue";
import { CarManu } from "../types/car";
import axios from "axios";
import { ElMessage } from 'element-plus';

export default function(){
    let carManu = reactive<CarManu>({
        carID: '',
        workshop: ''
    })

    async function SetCarManu() {
        const formData = new URLSearchParams();
        formData.append('carID', carManu.carID as string)
        formData.append('workshop', carManu.workshop)
        try {
            const response = await axios.post(
                '/SetCarManu',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarManu Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            ElMessage.error('SetCarManu Error: ' + error)
        }
    }

    return {
        carManu,
        SetCarManu
    }
}