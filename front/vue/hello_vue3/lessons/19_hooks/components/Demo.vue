<template>
    <div class="main">
        <h2>求和：{{ sum }}</h2>
        <h2>求和*10：{{ bigSum }}</h2>
        <button @click="add">add</button>
        <hr>
        <img v-for="(dog,idx) in dogList"   :src="dog" :key="idx" class="dog"><br>
        <button @click="moreDog">再来一只小狗</button>
    </div>
</template>

<script lang="ts" setup name="Demo">
    import useSum from '@/hooks/useSum'
    import useDog from '@/hooks/useDog'
    // 使用hooks 封装模块
    const {sum, bigSum, add} = useSum()
    const {dogList, moreDog} = useDog()
</script>
<style  scoped>
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
.dog {
    height: 100px;
}
</style>