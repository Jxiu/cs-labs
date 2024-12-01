<template>
    <div class="game">
        <h2>游戏</h2>
        <span>游戏{{ name }},版本 :V{{game.version}},价格: {{ price }}</span><br/>
        <button @click="discount">促销</button>
        <button @click="changeGame">换新</button>
        <button @click="newVersion">更新</button>
    </div>

</template>
<script lang="ts" setup  name="Game">
    import { toRefs,toRef } from 'vue';
    import { reactive } from 'vue';

    let  game = reactive({name:'Terraia',price:48, version: 1})
    let {name, price} = toRefs(game)
    let version = toRef(game, 'version')
    function discount(){
        price.value -= 1
    }
    function changeGame(){
        name.value += "~"
        game.version = 1
        // 注意：game对象的值也被改变了
        console.log(game.name)
    }
    function newVersion (){
        version.value +=0.1
    }

</script>

<style scoped>
    .game{
    background-color: rgb(98, 191, 231);
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
    margin: 10px 0px;
}

button {
    margin: 0 5px;

}
</style>