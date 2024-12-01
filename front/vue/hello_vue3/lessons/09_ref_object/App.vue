<template>
        <Car/>
        <Person/>
</template> 
<script lang="ts">
import Person from './components/person.vue';
import Car from './components/Car.vue';
    export default {
        name: 'App', //组件名字
        components:{Person,Car} //注册组件
    }
</script>
