<template>
  <ul class="news-detail">
    <li>编号：{{ query.id }}</li>
    <li>标题：{{ query.title }}</li>
    <li>内容：{{ query.content }}</li>
  </ul>
</template>

<script lang="ts" setup name="Detail">
// 路由不能用defineProps传递参数
//  let x = defineProps(["news"])
    import { toRefs } from 'vue';
import { useRoute } from 'vue-router';
    const route = useRoute()
    console.log('@',route)
    let {query} = toRefs(route)
</script>

<style scoped>


</style>