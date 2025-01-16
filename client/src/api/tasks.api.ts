import axios from 'axios'
import { API } from '.'

export const getAllTasks = async () => {
    try {
        const res = await axios.get(`${API}/task/all`)
        return res.data.data
    } catch (err) {
        console.error(err)
        return []
    }
}
