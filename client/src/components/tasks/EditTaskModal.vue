<script setup lang="ts">
import { capitalize, watch, computed, ref } from 'vue';
import { useTaskStore } from '../../store/task.store';
import { getTask } from '../../api/tasks.api';
import closeIcon from '../../assets/close.svg';
import { useCategoryStore } from '../../store/category.store';
import { Statuses } from '../../constants/status';
import type { Task } from '../../constants/types';

const taskStore = useTaskStore();
const categoryStore = useCategoryStore();
const fetchedTask = ref<Task | null>(null);

const taskName = defineModel<string>('taskName');
const taskCategory = defineModel<number | null>('taskCategory');
const taskStatus = defineModel<string>('taskStatus');
const taskMinutes = defineModel<number>('taskMinutes');

const closeEditTaskModal = () => {
    taskStore.tempEditTask = null;
};

const editTask = async () => {
    if (!taskName.value) {
        alert('Please enter a task name');
        return;
    }
    if (fetchedTask.value && taskStore.tempEditTask?.id) {
        taskStore.handleUpdateTask(taskStore.tempEditTask.id, {
            id: fetchedTask.value.id,
            name: taskName.value,
            category: taskCategory.value ?? undefined,
            minutes: taskMinutes.value ?? undefined,
            current_status: taskStatus.value ?? undefined,
            timestamp: fetchedTask.value.timestamp,
        });
        closeEditTaskModal();
    }
};

const fetchTask = async (id: number) => {
    const res = await getTask(id);
    if (res) {
        fetchedTask.value = res;
        taskName.value = res.name;
        taskCategory.value = res.category;
        taskStatus.value = res.current_status;
        taskMinutes.value = res.minutes;
    }
};

watch(
    () => taskStore.tempEditTask,
    (newVal) => {
        if (newVal?.id) {
            fetchTask(newVal?.id);
        }
    }
);

const categoryOptions = computed(() => {
    return [
        { id: null, name: 'None' },
        ...categoryStore.categories.map((cate) => ({
            id: cate.id,
            name: cate.name,
        })),
    ];
});
</script>

<template>
    <Transition name="fade">
        <div
            v-show="taskStore.tempEditTask?.show"
            class="fixed bg-black/30 flex justify-center items-center w-screen h-screen z-50 top-0 left-0"
            @click.prevent.self.stop="closeEditTaskModal"
        >
            <Transition name="slide-up">
                <div
                    v-show="taskStore.tempEditTask?.show"
                    class="relative flex flex-col justify-center items-center bg-white rounded-md border-2 border-black/30 overflow-hidden w-full max-w-96 font-Roboto"
                >
                    <div
                        class="absolute top-2 right-2 bg-slate-200 hover:bg-slate-300 transition-all duration-100 border-black border-2 rounded size-8 flex items-center justify-center cursor-pointer"
                        @click="closeEditTaskModal"
                    >
                        <img :src="closeIcon" alt="X" class="size-6" />
                    </div>
                    <h2
                        class="font-FunnelSans text-2xl font-medium w-full text-left h-12 px-2 flex flex-row items-center bg-active/25"
                    >
                        Edit Task
                    </h2>
                    <div
                        class="w-full p-4 h-auto flex flex-col justify-start items-start gap-2"
                    >
                        <div class="flex flex-col gap-1 w-full">
                            <label
                                for="taskName"
                                class="font-Roboto text-sm text-black/75"
                                >Name</label
                            >
                            <input
                                name="taskName"
                                id="taskName"
                                v-model="taskName"
                                type="text"
                                class="bg-slate-100 hover:bg-slate-200 transition-all duration-100 border rounded w-full px-1"
                            />
                        </div>
                        <div class="flex flex-col gap-1 w-full">
                            <label
                                for="taskCategory"
                                class="font-Roboto text-sm text-black/75"
                            >
                                Category
                            </label>
                            <select name="taskCategory" v-model="taskCategory">
                                <option
                                    v-for="cate in categoryOptions"
                                    :value="cate.id"
                                >
                                    {{ cate.name }}
                                </option>
                            </select>
                        </div>
                        <div class="flex flex-col gap-1 w-full">
                            <label
                                for="taskStatus"
                                class="font-Roboto text-sm text-black/75"
                            >
                                Status
                            </label>
                            <select name="taskStatus" v-model="taskStatus">
                                <option
                                    v-for="status in Statuses"
                                    :value="status"
                                >
                                    {{ capitalize(status) }}
                                </option>
                            </select>
                        </div>
                        <div class="flex flex-col gap-1 w-full">
                            <label
                                for="taskMinutes"
                                class="font-Roboto text-sm text-black/75"
                                >Minutes</label
                            >
                            <input
                                name="taskMinutes"
                                id="taskMinutes"
                                v-model="taskMinutes"
                                type="number"
                                class="bg-slate-100 hover:bg-slate-200 transition-all duration-100 border rounded w-full px-1 text-base"
                            />
                        </div>
                        <button
                            class="mt-2 w-full h-8 text-center bg-active/25 hover:bg-active/40 transition-all ease-linear duration-[50ms] active:scale-[0.975] rounded-full border-2 border-black px-4"
                            @click="editTask"
                        >
                            Edit Task
                        </button>
                    </div>
                </div>
            </Transition>
        </div>
    </Transition>
</template>
