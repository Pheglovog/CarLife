import {reactive} from 'vue';
import {TransferRecord} from '../types/car';
import axios from '../axios';
import {ElMessage} from 'element-plus';

export default function(){
    let transferRecord = reactive<TransferRecord>({
        carID: '',
        NewUser: '',
        cost: 0,
    })

    async function SetTransferRecord() {
        const formData = new URLSearchParams();
        formData.append('carID', transferRecord.carID as string)
        formData.append('newOwnerID', transferRecord.NewUser)
        formData.append('cost', transferRecord.cost.toString())
        try {
            const response = await axios.post(
                '/TransferCar',
                formData,
                {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
            )
            ElMessage.success(`TransferCar Success, CarID: ${response.data.carID}, TXID: ${response.data.txid}`)
        } catch (error) {
            throw new Error('TransferCar Error: ' + error)
        }
    }

    return {
        transferRecord,
        SetTransferRecord
    }
}

