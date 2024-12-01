<template>
        <Person/>
</template> 
<script lang="ts">
import Person from './components/person.vue';
    export default {
        name: 'APP', //组件名字
        components:{Person} //注册组件
    }
</script>
