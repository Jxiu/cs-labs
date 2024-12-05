<template>
    <div class = "main">
        <h1>这是About</h1>
    </div>
</template>

<script setup lang="ts" namne="About">
import { onMounted,onUnmounted } from 'vue';

onMounted(()=>{
    console.log("About 视图挂载完成！")
})
onUnmounted(()=>{
    console.log("About 视图卸载完成！")
})
</script>

<style>

</style>