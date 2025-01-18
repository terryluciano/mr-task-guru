<script setup lang="ts">
const props = defineProps<{
    show: boolean
    title: string
    message: string
    onAccept: () => void
    onReject: () => void
}>()
</script>

<template>
    <div
        v-show="props.show"
        class="fixed bg-black/30 flex justify-center items-center w-screen h-screen z-50 top-0 left-0"
        id="delete-modal"
        @click.prevent.self.stop="props.onReject"
    >
        <div
            class="bg-white rounded-md border-2 border-black/30 overflow-hidden w-full max-w-96"
        >
            <div
                class="bg-red-400 w-full text-2xl font-FunnelSans h-10 flex justify-center items-center"
            >
                {{ props.title }}
            </div>
            <div class="flex flex-col gap-2 p-4">
                <p class="text-wrap text-left font-Roboto">
                    {{ props.message }}
                </p>
                <div class="flex flex-row gap-2 justify-center items-center">
                    <button
                        class="bg-blue-500 hover:bg-blue-600 text-white font-FunnelSans h-10 rounded w-full"
                        @click="props.onReject"
                    >
                        Cancel
                    </button>
                    <button
                        class="bg-red-500 hover:bg-red-600 text-white font-FunnelSans h-10 rounded w-full"
                        @click="props.onAccept"
                    >
                        Delete
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
