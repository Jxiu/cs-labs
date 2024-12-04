<template>
    <Demo v-if="isShow"/>
</template> 
<script lang="ts" setup name="App">
import Demo from './components/Demo.vue';
import { ref,onMounted} from 'vue';
    let isShow = ref(true)
    onMounted(()=>{
        console.log("父组件挂在完毕")
    })
</script>
