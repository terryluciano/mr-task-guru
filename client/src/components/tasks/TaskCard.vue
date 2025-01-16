<script setup lang="ts">
import { computed } from 'vue'
import type { Category, Task } from '../../constants/types'

const props = defineProps<{ task: Task; categories?: Category[] }>()

const category = computed<Category | undefined>(() => {
    if (!props.task.category || !props.categories) {
        return undefined
    } else {
        const category = props?.categories.find((cate) => {
            if (cate.id === props.task.category) {
                return cate
            }
        })
        return category ? category : undefined
    }
})
</script>

<template>
    <div
        class="flex flex-col gap-2 w-full h-auto border rounded-[4px] overflow-hidden p-2 font-Roboto"
    >
        <!-- Title -->
        <p class="text-lg leading-none mb-2">
            {{ task.name }}
        </p>

        <!-- Minutes -->
        <p class="text-sm leading-none">Estimate: {{ task.minutes }} Minutes</p>
        <!-- Category -->
        <div
            v-show="props.task.category && category?.name"
            class="text-sm py-0.5 px-2 w-fit rounded-full leading-none"
            :style="{
                backgroundColor: category?.color ?? '#CBD5E0',
            }"
        >
            <p>
                {{ category?.name }}
            </p>
        </div>

        <!-- Change Status -->

        <!-- Timestamp -->
        <p class="text-xs leading-none">
            {{ new Date(task.timestamp).toLocaleString() }}
        </p>
    </div>
</template>
