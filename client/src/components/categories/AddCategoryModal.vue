<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import closeIcon from '../../assets/close.svg';
import { useCategoryStore } from '../../store/category.store';

const categoryStore = useCategoryStore();

const categoryName = defineModel<string>('categoryName');

const colorOne = defineModel<string>('colorOne');
const colorTwo = defineModel<string>('colorTwo');
const colorThree = defineModel<string>('colorThree');

colorOne.value = '0';
colorTwo.value = '0';
colorThree.value = '0';

const categoryColor = computed(() => {
    return `rgb(${colorOne.value}, ${colorTwo.value}, ${colorThree.value})`;
});

const props = defineProps<{
    show: boolean;
}>();

const closeEmit = defineEmits(['closeAddCategoryModel']);

const closeAddCategoryModal = () => {
    closeEmit('closeAddCategoryModel');
};

const addCategory = async () => {
    if (!categoryName.value || categoryName.value.length === 0) {
        return;
    }
    const res = await categoryStore.handleAddCategory({
        name: categoryName.value,
        color: categoryColor.value,
    });

    if (res) {
        closeAddCategoryModal();
    } else {
        console.error('Failed to add category');
    }
};

watch(
    () => props.show,
    (newVal) => {
        if (newVal === false) {
            categoryName.value = '';
            colorOne.value = '0';
            colorTwo.value = '0';
            colorThree.value = '0';
        }
    }
);
</script>

<template>
    <Transition name="fade">
        <div
            v-show="props.show"
            class="fixed bg-black/30 flex justify-center items-center w-screen h-screen z-50 top-0 left-0"
            @click.prevent.self.stop="closeAddCategoryModal"
        >
            <Transition name="slide-up">
                <div
                    v-show="props.show"
                    class="relative flex flex-col justify-center items-center bg-white rounded-md border-2 border-black/30 overflow-hidden w-full max-w-96 font-Roboto"
                >
                    <div
                        class="absolute top-2 right-2 bg-slate-200 hover:bg-slate-300 transition-all duration-100 border-black border-2 rounded size-8 flex items-center justify-center cursor-pointer"
                        @click="closeAddCategoryModal"
                    >
                        <img :src="closeIcon" alt="X" class="size-6" />
                    </div>
                    <h2
                        class="font-FunnelSans text-2xl font-medium w-full text-left h-12 px-2 flex flex-row items-center bg-active/25"
                    >
                        Add a New Category
                    </h2>
                    <div
                        class="w-full p-4 h-auto flex flex-col justify-start items-start gap-3"
                    >
                        <div class="flex flex-col gap-1 w-full">
                            <label
                                for="categoryName"
                                class="font-Roboto text-sm text-black/75"
                                >Name</label
                            >
                            <input
                                name="categoryName"
                                id="categoryName"
                                v-model="categoryName"
                                type="text"
                                class="bg-slate-100 hover:bg-slate-200 transition-all duration-100 border rounded w-full px-1"
                            />
                        </div>
                        <div class="flex flex-col gap-1 w-full">
                            <p class="font-Roboto text-sm text-black/75">
                                Color
                            </p>
                            <div class="flex flex-row gap-2 w-full">
                                <div
                                    class="flex flex-col justify-between items-start w-full"
                                >
                                    <div class="flex flex-row gap-1 w-full">
                                        <p
                                            class="font-Roboto text-sm text-black/75"
                                        >
                                            R
                                        </p>
                                        <input
                                            type="range"
                                            min="0"
                                            max="255"
                                            step="1"
                                            v-model="colorOne"
                                            class="w-full"
                                        />
                                    </div>
                                    <div class="flex flex-row gap-1 w-full">
                                        <p
                                            class="font-Roboto text-sm text-black/75"
                                        >
                                            G
                                        </p>
                                        <input
                                            type="range"
                                            min="0"
                                            max="255"
                                            step="1"
                                            v-model="colorTwo"
                                            class="w-full"
                                        />
                                    </div>
                                    <div class="flex flex-row gap-1 w-full">
                                        <p
                                            class="font-Roboto text-sm text-black/75"
                                        >
                                            B
                                        </p>
                                        <input
                                            type="range"
                                            min="0"
                                            max="255"
                                            step="1"
                                            v-model="colorThree"
                                            class="w-full"
                                        />
                                    </div>
                                    <p class="font-Roboto text-base text-black">
                                        {{ categoryColor }}
                                    </p>
                                </div>
                                <div
                                    class="size-24 flex-shrink-0 rounded-md"
                                    :style="{
                                        backgroundColor: `${categoryColor}`,
                                    }"
                                ></div>
                            </div>
                        </div>
                        <button
                            class="mt-2 w-full h-8 text-center bg-active/25 hover:bg-active/40 transition-all ease-linear duration-[50ms] active:scale-[0.975] rounded-full border-2 border-black px-4"
                            @click="addCategory"
                        >
                            Add Category
                        </button>
                    </div>
                </div>
            </Transition>
        </div>
    </Transition>
</template>
