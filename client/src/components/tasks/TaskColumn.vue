<script setup lang="ts">
import type { Category, Task } from '../../constants/types'
import TaskCard from './TaskCard.vue'

type statusType = 'inactive' | 'active' | 'complete' | 'incomplete'

interface Props {
    title: string
    status: statusType
    tasks?: Task[]
    categories?: Category[]
}

const statusColorMap: Record<statusType, string> = {
    inactive: '#CBD5E0', // A softer gray
    active: '#63B3ED', // A bright but pleasing blue
    complete: '#68D391', // A fresh green
    incomplete: '#FC8181', // A warmer red
}

const props = defineProps<Props>()
</script>

<template>
    <div
        class="flex flex-col gap-0 min-w-80 border border-black/15 rounded-md h-full"
    >
        <div
            class="w-full h-12 text-xl flex items-center flex-row px-4 rounded-t-[5px] font-Roboto"
            :style="{ backgroundColor: statusColorMap[props.status] }"
        >
            {{ props.title }}
        </div>
        <div class="flex flex-col w-full h-full gap-2 p-2">
            <TaskCard
                v-for="task in props.tasks"
                :key="task.id"
                :task="task"
                :categories="props.categories"
            />
        </div>
    </div>
</template>
