<template>
  <div class="person">
    <h2>姓名:{{ name }}</h2>
    <h2>年龄:{{ age }}</h2>
    <button @click="changeAge">修改年龄</button>
    <button @click="changeName">修改名字</button>
    <button @click="showTel">查看联系方式</button>
  </div>
</template>

<!-- <script lang="ts">
// 默认名字是文件名的前缀，如果组件名和文件一致可以不写这个script标签
// 安装插件也可以直接在setup中设置name
export default {
    name:'Person'
}
</script> -->
<!-- 
name插件安装

1.安装插件 npm i vite-plugin-vue-setup-extend -D
2.<script lang="ts" setup name="Person"></script>
3.修改vite.config.ts
  import VueSetupExtend from 'vite-plugin-vue-setup-extend'
  plugins 中添加调用 VueSetupExtend(),

 -->
<script lang="ts" setup name="Person1">
// setup 语法糖, 脚本的语言必须和导出对象的脚本一致,缺省语言设置默认是js
// 不一致会出现以下错误：[@vue/compiler-sfc] <script> and <script setup> must have the same language type.
        // 数据
        let name = '风暴之灵' //注意此时name不是响应式的
        let age = 33
        let tel = '199998888'
        // 方法 
        function changeAge(){
            age += 1
        }
        function changeName(){ // 注意这样修改name是没有变化的
            console.log('click changeName')
            name = 'SA'
            console.log("new name: " + name)
        }
        function showTel(){
            alert(tel)
        }
</script>

<style scoped>

.person {
    background-color: skyblue;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
}

button {
    margin: 0 5px;

}

</style>