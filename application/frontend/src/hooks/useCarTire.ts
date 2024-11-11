import { reactive } from "vue";
import { CarTires } from "../types/car";
import axios from "../axios";
import { ElMessage } from 'element-plus';


export default function() {
    let carTires = reactive<CarTires>(
        {
            carID: '',
            width: 0,
            radius: 0,
            workshop: ''
        }
    )

    async function SetCarTires() {
        const formData = new URLSearchParams();
        formData.append('carID', carTires.carID as string)
        formData.append('width', carTires.width.toString())
        formData.append('radius', carTires.radius.toString())
        formData.append('workshop', carTires.workshop)
        try {
            const response = await axios.post(
                '/SetCarTires',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`SetCarTires Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('SetCarTires Error: ' + error)
        }
    }

    return {
        carTires,
        SetCarTires
    }
}