<template>
  <div class=".main">
    <h1>中国</h1>
    <!-- <h2 id="h2">湖北</h2>  -->
    <h2 ref="h2">湖北</h2>
    <h3>武汉</h3>
    <button @click="logH2">输出h2</button>
  </div>
</template>

<script lang="ts" setup name="Demo">
    import { ref , defineExpose} from 'vue';
    let h2 = ref()
    function logH2(){
        // <style  scoped> 会输出局部样式 
        console.log(h2.value)
    }
    // function logH2(){
    //     // 使用id 在父子组件中可能出现冲突不建议使用
    //     console.log(document.getElementById('h2'))
    // }
    let a = ref(0)
    let b = ref(1)
    let c = ref(2) // c没有暴露对父组件不可以见到
    defineExpose({a,b})
</script>


<style  scoped>
/* 局部样式 选择器只会选择当前组件的样式 环境隔离 */
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