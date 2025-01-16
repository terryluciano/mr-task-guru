import axios from 'axios'
import { API } from '.'

export const getAllCategories = async () => {
    try {
        const res = await axios.get(`${API}/category/all`)
        return res.data.data
    } catch (err) {
        console.error(err)
        return []
    }
}
