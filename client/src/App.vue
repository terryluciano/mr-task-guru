<script setup lang="ts">
import { onMounted, ref } from 'vue';
import NavBar from './components/NavBar.vue';
import TaskColumn from './components/tasks/TaskColumn.vue';
import { useTaskStore } from './store/task.store';
import addTaskIcon from './assets/add-task.svg';
import categoryIcon from './assets/category.svg';
import { useCategoryStore } from './store/category.store';
import AddTaskModal from './components/tasks/AddTaskModal.vue';
import DeleteTaskModal from './components/tasks/DeleteTaskModal.vue';
import EditTaskModal from './components/tasks/EditTaskModal.vue';

const taskStore = useTaskStore();
const categoryStore = useCategoryStore();

const showAddTaskModal = ref(false);

onMounted(async () => {
    taskStore.fetchTasks();
    categoryStore.fetchCategories();
});
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
                    @click="showAddTaskModal = true"
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
                />
                <TaskColumn
                    title="Active"
                    status="active"
                    :tasks="taskStore.sortedTasks.active"
                />
                <TaskColumn
                    title="Complete"
                    status="complete"
                    :tasks="taskStore.sortedTasks.complete"
                />
                <TaskColumn
                    title="Incomplete"
                    status="incomplete"
                    :tasks="taskStore.sortedTasks.incomplete"
                />
            </div>
        </div>
        <AddTaskModal
            :show="showAddTaskModal"
            @closeAddTaskModel="showAddTaskModal = false"
        />
        <DeleteTaskModal />
        <EditTaskModal />
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

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.slide-up-enter-active,
.slide-up-leave-active {
    transition: all 0.2s ease;
}

.slide-up-enter-from,
.slide-up-leave-to {
    transform: translateY(100%);
}

.slide-up-enter-to,
.slide-up-leave-from {
    transform: translateY(0%);
}
</style>
