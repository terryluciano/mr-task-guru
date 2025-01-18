import axios from 'axios'
import { API } from '.'
import type { Task } from '../constants/types'

export const getAllTasks = async () => {
    try {
        const res = await axios.get(`${API}/task/all`)
        return res.data.data
    } catch (err) {
        console.error(err)
        return []
    }
}

export const deleteTask = async (id: number): Promise<boolean> => {
    try {
        const res = await axios.delete(`${API}/task/remove/${id}`)
        if (res.status !== 200) {
            return false
        }
        return true
    } catch (err) {
        console.error(err)
        return false
    }
}

export const updateTask = async (id: number, task: Task): Promise<boolean> => {
    try {
        const res = await axios.put(`${API}/task/update/${id}`, task)
        if (res.status !== 200) {
            return false
        }
        return true
    } catch (err) {
        console.error(err)
        return false
    }
}
