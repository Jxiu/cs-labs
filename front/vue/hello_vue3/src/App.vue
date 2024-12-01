<template>
    <!-- :属性 是js表达式 没有:只是普通的xml标签属性 -->
    <h2 a="1+1" :b="1+1" c="x" :d="9" ref="aa"></h2>
    <!-- <Demo a="aFromParent" b="aFromParent" :list="devloperList"/> -->
    <Demo />
</template> 
<script lang="ts" setup name="App">
import Demo from './components/Demo.vue';
// @ 根目录
import { type PersonList } from '@/types';

import { reactive } from 'vue';
    let x = 9
    let devloperList= reactive<PersonList>([
            {name:"john",age:10,id:"1"},
            {name:"sim",age:11,id:"2"},
            {name:"jojo",age:7,id:"3",sex:"male"}
        ]) 
</script>
