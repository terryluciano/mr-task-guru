import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Category } from '../constants/types'

export const useCategoryStore = defineStore('category', () => {
    // states
    const categories = ref<Category[]>([])

    // actions
    const setCategories = (incomingCategories: Category[]) => {
        categories.value = incomingCategories
    }

    return {
        categories,
        setCategories,
    }
})
