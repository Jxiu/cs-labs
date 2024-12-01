<template>
    <div class="car">
        <h2>汽车信息 reactive</h2>
        <span>一辆{{ car.brand }},价格: {{car.price}}</span><br/>
        <button @click="changePrice">涨价</button>
        <button @click="changeCar">换车</button>
        <h2>游戏 ref</h2>
        <span>游戏{{ game.name }},价格: {{game.price}}</span><br/>
        <button @click="discount">促销</button>
        <button @click="changeGame">换新</button>
    </div>

</template>
<script lang="ts" setup  name="Car">
    import { ref } from 'vue';
    import { reactive } from 'vue';
    // ref和reactive区别
    // 1.reactive 只能代理对象,ref 可以作用在对象和基本类型
    // 2.ref必须使用value，才能改变值
    //   reactive重新分配一个新对象，会失去响应式（对象的引用从代理对象，改变到新对象了），可以使用Object.assign 将新对象的值都copy到代理对象中
    // 最佳实践
    // 1. 响应式基本类型使用ref 2.深层对象使用reactive
    let car = reactive({brand:'四驱',price:111})
    let  game = ref({name:'Terraia',price:48})
    // 方法
    function changePrice(){
        car.price += 10
    }
    function changeCar(){
        // 将新的对象的所有属性copy到代理对象中
        Object.assign(car, {brand:'火箭',price:999})
    }
    function discount(){
        game.value.price -= 1
    }
    function changeGame(){
        game.value = {name:'socroll of taiwu',price:68}
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