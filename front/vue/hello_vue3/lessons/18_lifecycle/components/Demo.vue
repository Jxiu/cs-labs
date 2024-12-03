<template>
    <div class="main">
        <h2>当前求和：{{ sum }}</h2>
        <button @click="add">求和</button>
    </div>
</template>

<script lang="ts" setup name="Demo">
    import { ref,onBeforeMount,onMounted,onBeforeUpdate,onUpdated,onBeforeUnmount,onUnmounted } from 'vue';
    let sum = ref(0)
    function add (){
        sum.value += 1
    }
    // 子组件先挂在完成后父组件才挂载
    onBeforeMount(()=>{
        console.log("onBeforeMount","挂载前")
    })
    onMounted(()=>{
        console.log("onMounted","挂载后")
    })
    onBeforeUpdate(()=>{
        console.log("onBeforeUpdate","更新前")
    })
    onUpdated(()=>{
        console.log("onUpdated","更新后")
    })
    onBeforeUnmount(()=>{
        console.log("onBeforeUnmount","卸载前")
    })
    onUnmounted(()=>{
        console.log("onUnmounted","卸载后")
    })
</script>
<style  scoped>
.main{
    background-color: rgb(98, 191, 231);
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
    margin: 10px 0px;
}
button {
    margin: 0 10px;
}
</style>