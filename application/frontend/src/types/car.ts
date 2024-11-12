export interface Car{
    carID: string
    tires: CarTires
    body: CarBody
    interior: CarInterior
    manu: CarManu
    store: CarStore
    insure: Insures|null
    maint: Maints|null
    record: Records|null
    owner: string
}

export interface Insures{
    insures: CarInsure[]
}

export interface Maints {
    maints: CarMaint[]
}

export interface Records{
    records: TransferRecord[]
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
    owner: string
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
    newUser: string
    oldUser?: string
    cost: number
    time?: Date
    txID?: string 
}