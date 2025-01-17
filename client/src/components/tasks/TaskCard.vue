<script setup lang="ts">
import { computed, ref, watchEffect } from 'vue'
import type { Category, Task } from '../../constants/types'
import pencilIcon from '../../assets/pencil.svg'
import deleteIcon from '../../assets/delete.svg'
import { useTaskStore } from '../../store'
import DeleteModal from '../DeleteModal.vue'

const taskStore = useTaskStore()

const props = defineProps<{ task: Task; categories?: Category[] }>()

const deleteModal = ref(false)

const onAcceptDelete = () => {
    taskStore.handleDeleteTask(props.task.id)
    deleteModal.value = false
}

const onRejectDelete = () => {
    deleteModal.value = false
}

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

watchEffect(() => {
    console.log(deleteModal.value)
})
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

        <!-- Change Status -->
        <!-- Add row of  buttons to change status of task-->

        <!-- Add pencil button to change details of task-->
        <div
            class="absolute top-2 right-2 flex flex-col gap-1 opacity-0 group-hover:opacity-100 transition-all ease-in-out duration-200"
        >
            <button
                class="bg-slate-200 border rounded size-8 flex items-center justify-center"
            >
                <img :src="pencilIcon" alt="pencil icon" class="size-6" />
            </button>
            <button
                class="bg-red-200 border rounded size-8 flex items-center justify-center"
                @click="
                    () => {
                        deleteModal = true
                    }
                "
            >
                <img :src="deleteIcon" alt="pencil icon" class="size-6" />
            </button>
        </div>

        <!-- Timestamp -->
        <p class="text-xs leading-none">
            {{ new Date(task.timestamp).toLocaleString() }}
        </p>

        <DeleteModal
            title="Delete Task"
            :message="`Are you sure you want to delete this task: ${props.task.name}`"
            :show="deleteModal"
            @accept="onAcceptDelete"
            @reject="onRejectDelete"
        />
    </div>
</template>
