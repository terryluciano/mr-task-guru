<script setup lang="ts">
import { computed, ref } from 'vue';
import type { Category, Task } from '../../constants/types';
import pencilIcon from '../../assets/pencil.svg';
import deleteIcon from '../../assets/delete.svg';
import { useTaskStore } from '../../store/task.store';
import DeleteModal from '../DeleteModal.vue';
import { useCategoryStore } from '../../store/category.store';

const taskStore = useTaskStore();
const categoryStore = useCategoryStore();

const props = defineProps<{ task: Task }>();

const deleteModal = ref(false);

const onAcceptDelete = () => {
    taskStore.handleDeleteTask(props.task.id);
    deleteModal.value = false;
};

const onRejectDelete = () => {
    deleteModal.value = false;
};

const category = computed<Category | undefined>(() => {
    if (!props.task.category || !categoryStore.categories) {
        return undefined;
    } else {
        const category = categoryStore.categories.find((cate) => {
            if (cate.id === props.task.category) {
                return cate;
            }
        });
        return category ? category : undefined;
    }
});
</script>

<template>
    <div
        class="flex flex-col gap-2 w-full h-auto border rounded-[4px] overflow-hidden p-2 font-Roboto group relative"
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

        <!-- Timestamp -->
        <p class="text-xs leading-none">
            {{ new Date(task.timestamp).toLocaleString() }}
        </p>

        <!-- Change Status -->
        <div
            class="flex flex-row justify-center items-center gap-1 w-full h-auto mt-2"
        >
            <button
                v-show="
                    props.task.current_status != 'inactive' &&
                    props.task.current_status != null &&
                    props.task.current_status != undefined
                "
                @click="
                    () => {
                        taskStore.handleUpdateTask(props.task.id, {
                            ...props.task,
                            current_status: 'inactive',
                        });
                    }
                "
                class="bg-inactive w-full h-8 rounded active:scale-95 hover:brightness-105 transition-all ease-linear duration-[50ms]"
            >
                Inactive
            </button>
            <button
                v-show="props.task.current_status != 'active'"
                @click="
                    () => {
                        taskStore.handleUpdateTask(props.task.id, {
                            ...props.task,
                            current_status: 'active',
                        });
                    }
                "
                class="bg-active w-full h-8 rounded active:scale-95 hover:brightness-105 transition-all ease-linear duration-[50ms]"
            >
                Active
            </button>
            <button
                v-show="props.task.current_status != 'complete'"
                @click="
                    () => {
                        taskStore.handleUpdateTask(props.task.id, {
                            ...props.task,
                            current_status: 'complete',
                        });
                    }
                "
                class="bg-complete w-full h-8 rounded active:scale-95 hover:brightness-105 transition-all ease-linear duration-[50ms]"
            >
                Complete
            </button>
            <button
                v-show="props.task.current_status != 'incomplete'"
                @click="
                    () => {
                        taskStore.handleUpdateTask(props.task.id, {
                            ...props.task,
                            current_status: 'incomplete',
                        });
                    }
                "
                class="bg-incomplete w-full h-8 rounded active:scale-95 hover:brightness-105 transition-all ease-linear duration-[50ms]"
            >
                Incomplete
            </button>
        </div>

        <!-- Add pencil button to change details of task-->
        <div
            class="absolute top-2 right-2 flex flex-col gap-1 opacity-0 group-hover:opacity-100 transition-all ease-in-out duration-200"
        >
            <button
                class="bg-slate-200 hover:bg-slate-300 transition-all duration-100 border rounded size-8 flex items-center justify-center"
            >
                <img :src="pencilIcon" alt="pencil icon" class="size-6" />
            </button>
            <button
                class="bg-red-200 hover:bg-red-300 transition-all duration-100 border rounded size-8 flex items-center justify-center"
                @click="
                    () => {
                        deleteModal = true;
                    }
                "
            >
                <img :src="deleteIcon" alt="pencil icon" class="size-6" />
            </button>
        </div>
        <DeleteModal
            title="Delete Task"
            :message="`Are you sure you want to delete this task: ${props.task.name}`"
            :show="deleteModal"
            @accept="onAcceptDelete"
            @reject="onRejectDelete"
        />
    </div>
</template>
