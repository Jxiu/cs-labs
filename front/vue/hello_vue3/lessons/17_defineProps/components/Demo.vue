<template>
    <div class="main">
        <!-- <h2>{{ a }}</h2>
        <h2>{{ x.b }}</h2> -->
        <h2>开发人员</h2>
        <ul>
            <li v-for="item in list" :key="item.id">
                {{ item.name }}--{{ item.age }}
            </li>
        </ul>

    </div>
</template>

<script lang="ts" setup name="Demo">
    //vue define 开头函数是宏函数不需要引入
    import { defineProps,withDefaults } from 'vue';
    import { type PersonList } from '@/types';

    // 只接受list，没参数校验可能导致错误
    // let x = defineProps(['a','b','list']);
    // console.log(x)
    // 不能同时存在两个defineProps函数
    // [plugin:vite:setup-name-support] [@vue/compiler-sfc] duplicate defineProps() call
    
    // 接受list + 限制类型
    // defineProps<{list : PersonList}>();

    // 接收list + 限制类型 + 指定默认值
    withDefaults(defineProps<{list?:PersonList}>(),
    {
        list:()=>[{name:"default",age:10,id:"1"}]
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