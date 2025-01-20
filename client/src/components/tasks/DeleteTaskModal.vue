<script setup lang="ts">
import { useTaskStore } from '../../store/task.store';

const taskStore = useTaskStore();

const emit = defineEmits(['closeDeleteTaskModal']);

const closeDeleteTaskModal = () => {
    taskStore.setTempDeleteTask({ reset: true });
};

const deleteTask = () => {
    if (taskStore.tempDeleteTask == null) {
        return;
    }

    taskStore.handleDeleteTask(taskStore.tempDeleteTask.id).then(() => {
        closeDeleteTaskModal();
    });
};
</script>

<template>
    <Transition name="fade">
        <div
            v-show="taskStore.tempDeleteTask?.show"
            class="fixed bg-black/30 flex justify-center items-center w-screen h-screen z-50 top-0 left-0"
            id="delete-modal"
            @click.prevent.self.stop="closeDeleteTaskModal"
        >
            <Transition name="slide-up">
                <div
                    v-show="taskStore.tempDeleteTask?.show"
                    class="bg-white rounded-md border-2 border-black/30 overflow-hidden w-full max-w-96"
                >
                    <div
                        class="bg-red-400 w-full text-2xl font-FunnelSans h-10 flex justify-center items-center"
                    >
                        Delete Task
                    </div>
                    <div class="flex flex-col gap-4 p-4 leading-none">
                        <p class="text-wrap text-center font-Roboto">
                            Would you like to delete the following task?
                        </p>
                        <p
                            class="text-wrap text-center font-semibold text-lg font-Roboto"
                        >
                            {{ taskStore.tempDeleteTask?.name }}
                        </p>
                        <div
                            class="flex flex-row gap-2 justify-center items-center"
                        >
                            <button
                                class="bg-blue-500 hover:bg-blue-600 text-white font-FunnelSans h-10 rounded w-full"
                                @click="closeDeleteTaskModal"
                            >
                                Cancel
                            </button>
                            <button
                                class="bg-red-500 hover:bg-red-600 text-white font-FunnelSans h-10 rounded w-full"
                                @click="deleteTask"
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                </div>
            </Transition>
        </div>
    </Transition>
</template>
