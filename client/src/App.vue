<script setup lang="ts">
import { onMounted, ref } from 'vue'
import NavBar from './components/NavBar.vue'
import TaskColumn from './components/tasks/TaskColumn.vue'
import type { Category } from './constants/types'
import { getAllTasks } from './api/tasks.api'
import { getAllCategories } from './api/categories.api'
import { useTaskStore } from './store'
import addTaskIcon from './assets/add-task.svg'
import categoryIcon from './assets/category.svg'

const taskStore = useTaskStore()

const categories = ref<Category[]>([])

onMounted(async () => {
    const tasks = await getAllTasks()
    taskStore.setTasks(tasks)

    categories.value = await getAllCategories()
})
</script>

<template>
    <div class="w-full h-full flex flex-col justify-start items-center gap-10">
        <NavBar />
        <div class="flex flex-col gap-4">
            <div
                class="flex flex-row gap-2 justify-center items-center font-Roboto text-lg font-medium"
            >
                <button
                    class="h-10 rounded-full border-2 border-black px-4 flex flex-row items-center justify-center gap-1 bg-active/25 hover:bg-active/40 transition-all ease-linear duration-[50ms] active:scale-[0.975]"
                >
                    <img :src="addTaskIcon" alt="" class="size-5" />
                    <p>Add Task</p>
                </button>
                <button
                    class="h-10 rounded-full border-2 border-black px-4 flex flex-row items-center justify-center gap-1 bg-complete/25 hover:bg-complete/40 transition-all ease-linear duration-[50ms] active:scale-[0.975]"
                >
                    <img :src="categoryIcon" alt="" class="size-5" />
                    <p>Add Category</p>
                </button>
            </div>
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
    </div>
</template>

<style>
#app {
    margin: 0;
    padding: 0;
}
html {
    background-color: #fffafb;
}
body {
    margin: 0;
    padding: 0;
}
</style>
