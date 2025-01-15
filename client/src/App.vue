<script setup lang="ts">
import { onMounted } from "vue";
import NavBar from "./components/NavBar.vue";
import TaskColumn from "./components/tasks/TaskColumn.vue";
import axios from "axios";

const getAllTasks = async () => {
    try {
    const res = await axios.get("http://localhost:4000/task/all")
    console.log(res.data)
    }
    catch (err) {
        console.log(err)
    }
}

onMounted(() => {
    getAllTasks()
})
</script>

<template>
    <div class="w-full h-full flex flex-col justify-start items-center gap-10">
        <NavBar />
        <div class="grid grid-cols-4 gap-2">
            <TaskColumn title="Inactive" status="inactive" />
            <TaskColumn title="Active" status="active" />
            <TaskColumn title="Complete" status="complete" />
            <TaskColumn title="Incomplete" status="incomplete" />
        </div>
    </div>
</template>

<style>
#app {
    margin: 0;
    padding: 0;
}
body {
    margin: 0;
    padding: 0;
}
</style>
