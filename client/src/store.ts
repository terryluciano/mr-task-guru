import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import type { Task } from './constants/types'
import { deleteTask } from './api/tasks.api'

interface SortedTasks {
    inactive: Task[]
    active: Task[]
    complete: Task[]
    incomplete: Task[]
}

export const useTaskStore = defineStore('task', () => {
    // states
    const tasks = ref<Task[]>([])

    // getters
    const sortedTasks = computed<SortedTasks>(() => {
        const sorted: SortedTasks = {
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

    // actions
    const setTasks = (incomingTasks: Task[]) => {
        tasks.value = incomingTasks
    }

    const handleDeleteTask = async (id: number) => {
        const res = await deleteTask(id)
        if (res) {
            tasks.value = tasks.value.filter((task) => task.id !== id)
        }
    }

    return {
        tasks,
        sortedTasks,
        setTasks,
        handleDeleteTask,
    }
})
