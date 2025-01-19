<script setup lang="ts">
import type { Category, Task } from '../../constants/types'
import TaskCard from './TaskCard.vue'
import gsap from 'gsap'

type statusType = 'inactive' | 'active' | 'complete' | 'incomplete'

interface Props {
    title: string
    status: statusType
    tasks?: Task[]
}

const statusColorMap: Record<statusType, string> = {
    inactive: 'bg-inactive',
    active: 'bg-active',
    complete: 'bg-complete',
    incomplete: 'bg-incomplete',
}

const props = defineProps<Props>()

const onBeforeEnter = (el: Element) => {
    const htmlEl = el as HTMLElement
    htmlEl.style.opacity = '0'
    htmlEl.style.height = '0'
    htmlEl.style.transform = 'translateX(10%)'
}

function onEnter(el: Element, done: () => void) {
    const htmlEl = el as HTMLElement
    gsap.to(htmlEl, {
        opacity: 1,
        height: 'auto',
        transform: 'translateX(0)',
        delay: parseFloat(htmlEl.dataset.index ?? '0') * 0.15,
        onComplete: done,
    })
}

function onLeave(el: Element, done: () => void) {
    const htmlEl = el as HTMLElement
    gsap.to(htmlEl, {
        opacity: 0,
        height: 0,
        transform: 'translateX(10%)',
        delay: parseFloat(htmlEl.dataset.index ?? '0') * 0.15,
        onComplete: done,
    })
}
</script>

<template>
    <div
        class="flex flex-col gap-0 min-w-80 border border-black/15 rounded-md h-full"
    >
        <div
            :class="`w-full h-12 text-xl flex items-center flex-row px-4 rounded-t-[5px] font-Roboto ${
                statusColorMap[props.status]
            }`"
        >
            {{ props.title }}
        </div>
        <TransitionGroup
            tag="div"
            class="flex flex-col w-full h-full gap-2 p-2"
            @before-enter="onBeforeEnter"
            @enter="onEnter"
            @leave="onLeave"
        >
            <TaskCard
                v-for="(task, index) in props.tasks"
                :key="task.id"
                :task="task"
                :data-index="index"
            />
        </TransitionGroup>
    </div>
</template>
