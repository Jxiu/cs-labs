<template>
    <div class = "main">
        <!-- 导航区 -->
        <ul class=".left">
            <li v-for="news in newsList" :key="news.id">
                <!-- 模板字符串语法 ${} 嵌入值-->
                <!-- <RouterLink :to="`/news/detail?id=${news.id}&title=${news.title}&content=${news.content}`" :news="news">{{news.title}}</RouterLink> -->
                <!-- <RouterLink :to="`/news/detail/${news.id}/${news.title}/${news.content}`" :news="news">{{news.title}}</RouterLink> -->
                <RouterLink :to="{
                    name:'xiang', //注意：只能使用name，不能使用path
                    params:{
                        id:news.id,
                        title:news.title,
                        content:news.content,
                    }
                }">{{news.title}}</RouterLink>
            </li>
        </ul>
        <!-- 展示区 -->
        <div class="right">
            <RouterView></RouterView>
        </div>
    </div>
</template>

<script setup lang="ts" namne="News">
import { onMounted,onUnmounted, reactive } from 'vue';
import { RouterView,RouterLink } from 'vue-router';

const newsList = reactive([
    {id:"1",title:"减脂食物",content:"西蓝花"},
    {id:"2",title:"万万没想到",content:"今天是周五！！"},
    {id:"3",title:"老毕登周末",content:"没有员工加班，就裁员！！"},
    {id:"4",title:"珠海航展",content:"歼十一上天，红箭19首现！！"},
])

onMounted(()=>{
    console.log("News 视图挂载完成！")
})
onUnmounted(()=>{
    console.log("News 视图卸载完成！")
})
</script>

<style scoped>
.main{
    height: 90%;
    padding-top: 10px;
    padding-right: 20px;
    display: grid;
    grid-template-columns: 1fr 3fr; /* 创建两个等宽的列 */
}
ul {
    margin-top: 30px;
} 
 ul li::marker{
    color: green;
 }
 ul li>a{
    color: green;
    line-height: 40px;
    font-size: 18px;
 }

.right {
  border-radius: 5px;
  border: solid;
  height: 100%;
}
</style>