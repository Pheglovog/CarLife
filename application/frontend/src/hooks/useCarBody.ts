import { reactive } from "vue"
import { CarBody } from "../types/car"
import axios from '../axios';
import { ElMessage } from 'element-plus';

export default function(){
    const carBody = reactive<CarBody>({
        carID: "",
        material: "",
        weight: 0,
        color: "",
        workshop: ""
    })

    async function SetCarBody() {
        const formData = new URLSearchParams();
        formData.append('carID', carBody.carID as string)
        formData.append('material', carBody.material)
        formData.append('weight', carBody.weight.toString())
        formData.append('color', carBody.color)
        formData.append('workshop', carBody.workshop)
        try {
            const response = await axios.post(
                '/SetCarBody',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarBody Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarBody Error: ' + error)
        }
    }

    return {
        carBody,
        SetCarBody
    }
}