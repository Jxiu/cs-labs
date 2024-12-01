<template>
  <div class="person">
    <h2>姓名:{{ name }}</h2>
    <h2>年龄:{{ age }}</h2>
    <button @click="changeAge">修改年龄</button>
    <button @click="changeName">修改名字</button>
    <button @click="showTel">查看联系方式</button>
    <!-- 分割线 hr -->
    <hr>
    <h2>test : {{ data1 }}</h2>
    <h2>test2: {{ data2 }}</h2>
    <button @click="testMethodsExec">Methods执行</button>
  </div>
</template>

<script>
export default {
    // 猜测原理是将 setup中暴露出来的对象和方法，和data返回的对象整合，methods返回的方法整合 到这个导出的对象中，方法的执行顺序决定了能否读取到属性
    name:'Person',
    // data(){
    //     return {
    //         name : '张三',
    //         age: 11,
    //         tel: '13733332222'
    //     }
    // },
    // methods:{
    //     showTel(){
    //         alert(this.tel)
    //     },
    //     changeAge(){
    //         this.age +=1;
    //     },
    //     changeName(){
    //         this.name = 'john'
    //     }
    // },
    beforeCreate(){
        console.log("beforeCreate event")
    },
    data(){
        //三个函数执行顺序setup beforeCreate data
        console.log("data method execute")
        return {
            data1 : 111,
            data2 : this.name
        }
    },
    methods:{
        testMethodsExec(){
            console.log("methods execute~")
        }
    },
    setup(){
        console.log("setup event")
        // setup中的this 是undefined,vue3弱化this
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
        // 将内部属性封装到对象中暴露出去 
        return {name,age,tel,changeAge,changeName,showTel}
    }
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