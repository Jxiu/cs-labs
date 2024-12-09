<template>
    <div class = "main">
        <h1>这是Home</h1>
    </div>
</template>

<script setup lang="ts" namne="Home">
import { onMounted,onUnmounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
onMounted(()=>{
    console.log("Home 视图挂载完成！")
    setTimeout(()=>{
        router.replace("/news")
    },2000)
})
onUnmounted(()=>{
    console.log("Home 视图卸载完成！")
})
</script>

<style>

</style>