import axios from 'axios';
import { API } from './index.api';
import type { Task } from '../constants/types';

export const getAllTasks = async () => {
    try {
        const res = await axios.get(`${API}/task/all`);
        return res.data.data;
    } catch (err) {
        console.error(err);
        return [];
    }
};

export const deleteTask = async (id: number): Promise<boolean> => {
    try {
        const res = await axios.delete(`${API}/task/remove/${id}`);
        if (res.status !== 200) {
            return false;
        }
        return true;
    } catch (err) {
        console.error(err);
        return false;
    }
};

export const putUpdateTask = async (
    id: number,
    task: Task
): Promise<boolean> => {
    try {
        const res = await axios.put(`${API}/task/update/${id}`, task);
        if (res.status !== 200) {
            return false;
        }
        return true;
    } catch (err) {
        console.error(err);
        return false;
    }
};

export const postAddTask = async (
    task: Omit<Task, 'id' | 'timestamp'>
): Promise<boolean> => {
    const res = await axios.post(`${API}/task/add`, task);
    if (res.status !== 200) {
        throw new Error('Failed to add task');
    }
    return true;
};
