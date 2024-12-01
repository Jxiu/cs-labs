<template>
  <div class=".main">
    <h2>情况一：监视【ref】定义的【基本类型】数据</h2>
    <h3>当前求和：{{ sum }}</h3>
    <button @click="add1">+1</button>
    <hr>
    <h2>情况二：监视【ref】定义的【对象类型】数据</h2>
    <span>姓名:{{ person.name }}</span><br>
    <span>年龄:{{ person.age }}</span><br>
    <button @click="changeName">改名</button>
    <button @click="changePerson">变身</button>
    <hr>
    <h2>情况三：监视【reactive】定义的【对象】数据</h2>
    <span>游戏名称：{{ game.name }}</span><br>
    <span>售价：{{ game.price }}</span><br>
    <span>PC发行商：{{ game.publisher.pc }}</span><br>
    <span>Xbox发行商 {{ game.publisher.xbox }}</span><br>
    <button @click="discount">促销</button>
    <button @click="newGame">上新</button>
    <h2>情况四：函数返回值（getter函数）</h2>
    <button @click="changePC">PC更换发行商</button>
    <button @click="changeXbox">xbox更换发行商</button>
    <button @click="changeAllPublisher">更换所有发行商</button>

  </div>
</template>

<script lang="ts" setup name="Demo">
    import { ref,reactive,watch } from 'vue';
    // watch 只能监视以下四种情况
    // 1.ref定义的数据
    // 2.reactive定义的数据
    // 3.函数返回值（getter函数）
    // 4.一个包含上述内容的数组
    let sum = ref(0)
    let person = ref({name:"Terraia",age:33})
    let game = reactive({name:"Terraia",
    price:48,
    publisher:{
        pc:"steam",
        xbox:"microSoft",
    }})
    // 方法
    function add1(){
        sum.value +=1
    }
    // 监视器，监视的是代理对象
    const watchRef = watch(sum, (newVal,oldVal)=>{
        console.log("sum 监视器触发！", oldVal, newVal)
        if(newVal > 5){
            // 停止监视
            watchRef()
        }
    })

    function changeName(){
        person.value.name += "~" 
    }
    function changePerson(){
        person.value = {name:"lisa",age:18}
    }
    // 2.1 监视的是对象地址变化。监视属性变化需要开启手动监视
    // immediate 立即监视，启动就监视
    watch(person,(newVal,oldVal)=>{
        console.log("person 监视器触发！", oldVal, newVal)
        // 深度监视时候，地址没变时候，old=new
    },{deep: true,immediate: true})

    function discount(){
        game.price -= 1
    }
    function newGame(){
        Object.assign(game,{name:"don't strave together", price:55})
    }
    // 3.监视reactive默认是开启深度监视的
    watch(game,(newVal,oldVal)=>{
        console.log("game 监视器触发！", oldVal, newVal)
    })
    // 4.监视对象的某个属性
    watch(()=> game.price ,(newVal,oldVal)=>{
        console.log("game price 监视器触发！", oldVal, newVal)
    })
    function changePC(){
        game.publisher.pc += "+"
    }
    function changeXbox(){
        alert("not support")
    }
    function changeAllPublisher(){
        game.publisher = {pc:"Epic",xbox:"no one"}
    }
    // 4.1 更换整个对象地址不触发,而且地址变化后无法继续触发监视
    watch(game.publisher ,(newVal,oldVal)=>{
        console.log("game publisher 监视器触发！", oldVal, newVal)
    })
    // 4.1 监视内嵌对象的地址变化, deep 开启属性变化也能监视
    watch(()=> game.publisher ,(newVal,oldVal)=>{
        console.log("game publisher seter 监视器触发！", oldVal, newVal)
    },{deep: true})
    // 5 监视多个对象的数组
    watch([sum,()=>game.price],()=>{
        console.log("数字变化监视器触发！")
    })

</script>

<style>

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