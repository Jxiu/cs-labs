<template>
    <div class="app">
        <Header></Header>
        <!-- 导航区 -->
        <div class="navigate">
            <!-- 默认浏览器历史记录保存模式是push，可以前后跳转，replace模式是replace，只能前进，不能后退。 -->
            <RouterLink replace to="/home" active-class="active">首页</RouterLink>
            <RouterLink replace :to="{name:'xinwen'}" active-class="active">新闻</RouterLink>
            <RouterLink replace :to="{path:'/about'}" active-class="active">关于</RouterLink>
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
    import Header from './components/Header.vue';
    // 路由组件放在views或者pages目录下（路由文件内引用的文件）
    // 标签组件放在components目录下（当标签用的组件 ）
</script>
<style scoped>
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
    height: 300px;
    margin-top: 10px;
    border: 1px solid;
}
</style>
