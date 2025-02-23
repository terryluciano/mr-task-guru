import axios from 'axios';
import { API } from './index.api';
import type { Category } from '../constants/types';

export const getAllCategories = async () => {
    try {
        const res = await axios.get(`${API}/category/all`);
        return res.data.data;
    } catch (err) {
        console.error(err);
        return [];
    }
};

export const postAddCategory = async (
    category: Omit<Category, 'id' | 'deleted' | 'timestamp'>
): Promise<Category | false> => {
    try {
        const res = await axios.post(`${API}/category/add`, {
            name: category.name,
            ...(category.color !== undefined ||
            category.color !== null ||
            category.color !== ''
                ? { color: category.color }
                : {}),
        });

        return res.data.data;
    } catch (err) {
        console.error(err);
        return false;
    }
};
