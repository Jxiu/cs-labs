<template>
  <div class="person">
    <h2>姓名:{{ name }}</h2>
    <h2>年龄:{{ age }}</h2>
    <button @click="changeAge">修改年龄</button>
    <button @click="changeName">修改名字</button>
    <button @click="showTel">查看联系方式</button>
    
  </div>
</template>

<script>
export default {
    name:'Person',
    // vue2--选项式（配置式）写法
    data(){
        return {
            name : '张三',
            age: 11,
            tel: '13733332222'
        }
    },
    methods:{
        showTel(){
            alert(this.tel)
        },
        changeAge(){
            this.age +=1;
        },
        changeName(){
            this.name = 'john'
        }
    }
}
</script>

<style>
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