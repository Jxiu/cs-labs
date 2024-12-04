<template>
    <div class="app">
        <h2 class="title">vue路由测试</h2>
        <!-- 导航区 -->
        <div class="navigate">
            <RouterLink to="/home" active-class="active">首页</RouterLink>
            <RouterLink to="/news" active-class="active">新闻</RouterLink>
            <RouterLink to="/about" active-class="active">关于</RouterLink>
        </div>
        <!-- 展示区 -->
        <div class="main-content">
            <!-- 此处以后可以根据路由展示多种组件 -->
             <RouterView></RouterView>
        </div>
    </div>
</template> 
<script lang="ts" setup name="App">
    import { RouterView,RouterLink } from 'vue-router';
</script>
<style scoped>
.title {
    background-color: skyblue;
    height: 70px;
    border-radius: 10px;
    text-align: center;
}
.navigate{
    /* display: flex;
    justify-content: space-between; */

    display: grid;
    grid-template-columns: repeat(3, 1fr); /* 创建三个等宽的列 */
    grid-gap: 10px; /* 设置网格间距 */
}
.navigate a{
    /* flex-grow: 1; 使每个 <a> 元素占据相同的空间 */
    text-align: center; /* 文本居中 */
    margin: auto;
    background-color: grey;
    color: white;
    height: 40px;
    width: 90px;
    font-size: 30px;
}
 .active{
    border: 3px solid;
    border-color: red;
    color: coral;
    background-color: lightseagreen;
}
.main-content{
    height: 600px;
    margin-top: 10px;
    border: 1px solid;
}
</style>
