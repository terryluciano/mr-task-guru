export interface Task {
    id: number
    name: string
    category?: number
    minutes?: number
    current_status?: string
    timestamp: string | Date
}

export interface Category {
    id: number
    name: string
    deleted?: boolean
    timestamp: string | Date
    color?: string
}
