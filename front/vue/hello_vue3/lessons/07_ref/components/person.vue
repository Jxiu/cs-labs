<template>
  <div class="person">
    <h2>姓名:{{ name }}</h2>
    <h2>年龄:{{ age }}</h2>
    <button @click="changeAge">修改年龄</button>
    <button @click="changeName">修改名字</button>
    <button @click="showTel">查看联系方式</button>
  </div>
</template>

<script lang="ts" setup name="Person1">
    import { ref } from 'vue';
        // 数据
        // ref 可以定义基本类型响应式数据
        let name = ref('风暴之灵' ) // 响应式
        let age = ref(33)
        let tel = '199998888' // 非相应

        console.log("name:" , name)
        console.log("age:" , age)
        console.log("tel:" , tel)

        // 方法 
        function changeName(){ 
            console.log('click changeName')
            name.value = 'SA'
            console.log("new name: " + name)
        }
        function changeAge(){
            age.value += 1
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