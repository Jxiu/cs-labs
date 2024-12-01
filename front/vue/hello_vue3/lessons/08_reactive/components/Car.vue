<template>
    <div class="car">
        <h2>汽车信息</h2>
        <span>一辆{{ car.brand }},价格{{car.price}}</span><br/>
        <button @click="changePrice">涨价</button>
        <hr>
        <h2>游戏展示</h2>
        <ul>
            <li v-for=" game in games" :key="game.id">{{game.name}}</li>
        </ul>
        <button @click="changeGame1name">修改第一个游戏名称</button>
        <hr>
        <h2>test deep: {{ obj.a.b.c }}</h2>
        <button @click="changeDeep">改变深层属性</button>
    </div>

</template>
<script lang="ts" setup  name="Car">
    import { reactive } from 'vue';
    // ref 可以定义基本类型响应式数据
    // reactive只能定义对象类型响应式数据
    let car = reactive({brand:'四驱',price:111})
    // ref 基本类型响应式数据

    let games = reactive([ 
        {id:"g1",name:"dota2"},
        {id:"g2",name:"太吾绘卷"},
        {id:"g3",name:"Terraria"}
    ])
    let obj = reactive({
        a:{
            b:{
                c: "10"
            }
        }
    })
    console.log("car: ", car)
    console.log(Proxy)
    // 方法
    function changePrice(){
        car.price += 10
    }
    function changeGame1name(){
        games[0].name = "饥荒"
    }
    function changeDeep(){
        obj.a.b.c = "666"
    }
</script>

<style scoped>
    .car{
    background-color: rgb(206, 202, 202);
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
    margin: 10px 0px;
}

button {
    margin: 0 5px;

}
</style>