<script setup lang="ts">
import { computed, onMounted, ref, watchEffect } from 'vue'
import NavBar from './components/NavBar.vue'
import TaskColumn from './components/tasks/TaskColumn.vue'
import type { Category, Task } from './constants/types'
import { getAllTasks } from './api/tasks.api'
import { getAllCategories } from './api/categories.api'

const tasks = ref<Task[]>([])
const categories = ref<Category[]>([])

interface SortTasks {
    inactive: Task[]
    active: Task[]
    complete: Task[]
    incomplete: Task[]
}

const sortedTasks = computed<SortTasks>(() => {
    const sorted: SortTasks = {
        inactive: [],
        active: [],
        complete: [],
        incomplete: [],
    }

    if (tasks.value.length == 0) {
        return sorted
    } else {
        tasks.value.forEach((task) => {
            switch (task.current_status) {
                case 'active':
                    sorted.active.push(task)
                    break
                case 'complete':
                    sorted.complete.push(task)
                    break
                case 'incomplete':
                    sorted.incomplete.push(task)
                    break
                default:
                    sorted.inactive.push(task)
                    break
            }
        })
        return sorted
    }
})

onMounted(async () => {
    tasks.value = await getAllTasks()
    categories.value = await getAllCategories()
})

watchEffect(() => {
    console.log(categories.value)
})
</script>

<template>
    <div class="w-full h-full flex flex-col justify-start items-center gap-10">
        <NavBar />
        <div class="grid grid-cols-4 gap-2">
            <TaskColumn
                title="Inactive"
                status="inactive"
                :tasks="sortedTasks.inactive"
                :categories="categories"
            />
            <TaskColumn
                title="Active"
                status="active"
                :tasks="sortedTasks.active"
                :categories="categories"
            />
            <TaskColumn
                title="Complete"
                status="complete"
                :tasks="sortedTasks.complete"
                :categories="categories"
            />
            <TaskColumn
                title="Incomplete"
                status="incomplete"
                :tasks="sortedTasks.incomplete"
                :categories="categories"
            />
        </div>
    </div>
</template>

<style>
#app {
    margin: 0;
    padding: 0;
}
body {
    margin: 0;
    padding: 0;
}
</style>
