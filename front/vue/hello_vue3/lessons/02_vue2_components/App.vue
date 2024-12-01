<template>
    <div class="app">
        <h1>hello world！</h1>
        <Person/>
    </div>
</template> 
<script lang="ts">
import Person from './components/person.vue';
    export default {
        name: 'APP', //组件名字
        components:{Person} //注册组件
    }
</script>
<style>
 .app {
    background-color: #ddd;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
 }
</style>