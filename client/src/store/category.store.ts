import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { Category } from '../constants/types';
import { getAllCategories, postAddCategory } from '../api/categories.api';

export const useCategoryStore = defineStore('category', () => {
    // states
    const categories = ref<Category[]>([]);

    // actions
    const setCategories = (incomingCategories: Category[]) => {
        categories.value = incomingCategories;
    };

    const fetchCategories = async () => {
        const res = await getAllCategories();
        setCategories(res);
    };

    const handleAddCategory = async (
        category: Omit<Category, 'id' | 'deleted' | 'timestamp'>
    ): Promise<boolean> => {
        try {
            const res = await postAddCategory(category);
            if (res) {
                categories.value.push(res);
                return true;
            }
            return false;
        } catch (err) {
            console.error(err);
            return false;
        }
    };

    return {
        categories,
        setCategories,
        fetchCategories,
        handleAddCategory,
    };
});
