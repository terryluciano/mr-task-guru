import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import type { Task } from '../constants/types';
import {
    postAddTask,
    deleteTask,
    putUpdateTask,
    getAllTasks,
} from '../api/tasks.api';

interface SortedTasks {
    inactive: Task[];
    active: Task[];
    complete: Task[];
    incomplete: Task[];
}

export const useTaskStore = defineStore('task', () => {
    // states
    const tasks = ref<Task[]>([]);

    const tempDeleteTask = ref<{
        show: boolean;
        id: number;
        name: string;
    } | null>(null);

    const tempEditTask = ref<{ show?: boolean; id?: number } | null>(null);

    // getters
    const sortedTasks = computed<SortedTasks>(() => {
        const sorted: SortedTasks = {
            inactive: [],
            active: [],
            complete: [],
            incomplete: [],
        };

        if (tasks.value.length == 0) {
            return sorted;
        } else {
            tasks.value.forEach((task) => {
                switch (task.current_status) {
                    case 'active':
                        sorted.active.push(task);
                        break;
                    case 'complete':
                        sorted.complete.push(task);
                        break;
                    case 'incomplete':
                        sorted.incomplete.push(task);
                        break;
                    default:
                        sorted.inactive.push(task);
                        break;
                }
            });
            return sorted;
        }
    });

    // actions
    const setTasks = (incomingTasks: Task[]) => {
        tasks.value = incomingTasks;
    };

    const handleDeleteTask = async (id: number) => {
        const res = await deleteTask(id);
        if (res) {
            tasks.value = tasks.value.filter((task) => task.id !== id);
        }
    };

    const handleUpdateTask = async (id: number, task: Task) => {
        try {
            const res = await putUpdateTask(id, task);
            if (res) {
                tasks.value = tasks.value.map((prevTask) => {
                    if (prevTask.id === id) {
                        return task;
                    }
                    return prevTask;
                });
            }
        } catch (err) {
            console.error(err);
        }
    };

    const handleAddTask = async (
        task: Omit<Task, 'id' | 'timestamp'>
    ): Promise<boolean> => {
        try {
            const res = await postAddTask(task);
            fetchTasks();
            return res;
        } catch (err) {
            console.error(err);
            return false;
        }
    };

    const fetchTasks = async () => {
        const res = await getAllTasks();
        setTasks(res);
    };

    const setTempDeleteTask = (opt: {
        id?: number;
        name?: string;
        reset?: boolean;
    }) => {
        if (opt.reset) {
            tempDeleteTask.value = null;
            return;
        }
        tempDeleteTask.value = {
            show: true,
            id: opt.id ?? 0,
            name: opt.name ?? '',
        };
    };

    return {
        tasks,
        sortedTasks,
        tempDeleteTask,
        tempEditTask,
        setTasks,
        fetchTasks,
        handleDeleteTask,
        handleUpdateTask,
        handleAddTask,
        setTempDeleteTask,
    };
});
