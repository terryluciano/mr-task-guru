<script setup lang="ts">
import { computed, onMounted, ref, watchEffect } from 'vue'
import NavBar from './components/NavBar.vue'
import TaskColumn from './components/tasks/TaskColumn.vue'
import type { Category, Task } from './constants/types'
import { getAllTasks } from './api/tasks.api'
import { getAllCategories } from './api/categories.api'
import { useTaskStore } from './store'

const taskStore = useTaskStore()

const categories = ref<Category[]>([])

onMounted(async () => {
    const tasks = await getAllTasks()
    taskStore.setTasks(tasks)

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
                :tasks="taskStore.sortedTasks.inactive"
                :categories="categories"
            />
            <TaskColumn
                title="Active"
                status="active"
                :tasks="taskStore.sortedTasks.active"
                :categories="categories"
            />
            <TaskColumn
                title="Complete"
                status="complete"
                :tasks="taskStore.sortedTasks.complete"
                :categories="categories"
            />
            <TaskColumn
                title="Incomplete"
                status="incomplete"
                :tasks="taskStore.sortedTasks.incomplete"
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
