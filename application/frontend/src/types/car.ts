export interface Car{
    carID: string
    tires: CarTires
    body: CarBody
    interior: CarInterior
    manu: CarManu
    store: CarStore
    insure: CarInsure[]
    maint: CarMaint[]
    record: TransferRecord[]
    owner: string
}


export interface CarTires{
    carID?: string
    width: number
    radius: number
    workshop: string
    time?: Date
    txID?: string
}

export interface CarBody{
    carID?: string
    material: string
    weight: number
    color: string
    workshop: string
    time?: Date
    txID?: string    
}

export interface CarInterior{
    carID?: string
    material: string
    color: string
    workshop: string
    time?: Date
    txID?: string    
}

export interface CarManu{
    carID?: string
    workshop: string
    time?: Date
    txID?: string
}

export interface CarStore{
    carID?: string
    store: string
    cost: number
    ownerID: string
    time?: Date
    txID?: string    
}

export interface CarInsure{
    carID?: string
    name: string
    cost: number
    years?: number
    beginTime?: Date
    endTime?: Date
    txID?: string
}

export interface CarMaint{
    carID?: string
    part: string
    extent: string
    cost: number
    time?: Date
    txID?: string    
}

export interface TransferRecord{
    carID?: string
    NewUser: string
    OldUser?: string
    cost: number
    time?: Date
    txID?: string 
}