import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { Category } from '../constants/types';
import { getAllCategories } from '../api/categories.api';

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

    return {
        categories,
        setCategories,
        fetchCategories,
    };
});
