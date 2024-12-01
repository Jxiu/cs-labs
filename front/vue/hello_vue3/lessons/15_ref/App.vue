<template>
    <h2 ref="h2"> 你好！</h2>
    <button @click="logH2">输出父组件h2</button>
    <Demo ref="ddd"/>
</template> 
<script lang="ts" setup name="App">
import Demo from './components/Demo.vue';
import { ref } from 'vue';
    // Vue2
    // export default {
    //     name: 'App', //组件名字
    //     components:{Demo} //注册组件
    // }
    let h2 = ref()
    let ddd = ref()
    function logH2(){
        console.log(h2.value)
        console.log(ddd.value.a)
    }
</script>
