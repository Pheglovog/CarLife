import { reactive } from "vue";
import { CarMaint } from "../types/car";
import axios from '../axios';
import { ElMessage } from 'element-plus';

export default function(){
    let carMaint = reactive<CarMaint>({
        carID: '',
        part: '',
        extent: '',
        cost: 0,
    })

    async function SetCarMaint() {
        const formData = new URLSearchParams();
        formData.append('carID', carMaint.carID as string)
        formData.append('part', carMaint.part)
        formData.append('extent', carMaint.extent)
        formData.append('cost', carMaint.cost.toString())
        try {
            const response = await axios.post(
                '/SetCarMaint',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarMaint Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarMaint Error: ' + error)
        }
    }

    return {
        carMaint,
        SetCarMaint
    }
}