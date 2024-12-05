<template>
    <div class = "main">
        <h1>这是Home</h1>
    </div>
</template>

<script setup lang="ts" namne="Home">
import { onMounted,onUnmounted } from 'vue';

onMounted(()=>{
    console.log("Home 视图挂载完成！")
})
onUnmounted(()=>{
    console.log("Home 视图卸载完成！")
})
</script>

<style>

</style>