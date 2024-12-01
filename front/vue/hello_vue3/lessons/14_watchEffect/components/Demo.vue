<template>
  <div class=".main">
    <h2>水文监测</h2>
    <span>需求：水温大于27度或者水位低于90告警</span>
    <h3>水温：{{ temp }}℃</h3>
    <h3>水位：{{ heigh }}cm</h3>
    <button @click="changeTemp">水温上升</button>
    <button @click="changeHeigh">水位下降</button>

  </div>
</template>

<script lang="ts" setup name="Demo">
    import { ref,watchEffect } from 'vue';
    let temp = ref(25)
    let heigh = ref(100)
    function changeTemp(){
        temp.value += 0.5
    }
    function changeHeigh(){
        heigh.value -= 3
    }

    // watch([temp,heigh],(value)=>{
    //     let [temp,heigh] = value
    //     if(temp>27 || heigh <90){
    //         console.log("告警：：水文情况异常！")
    //     }
    // })

    watchEffect(()=>{
        // 根据用到的属性自动监视
        if(temp.value >27 || heigh.value <90){
            console.log("告警：：水文情况异常！")
        }
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