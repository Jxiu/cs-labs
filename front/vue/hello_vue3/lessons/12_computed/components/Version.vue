<template>
    <div class="version">
        <h2>版本</h2>
        主版本：<input type="text" v-model="main"/>
        副版本：<input type="text" v-model="minor"/>
        计算版本号：<span>{{ version }}</span>
        计算版本号：<span>{{ version }}</span><br>
        <button @click="changeVersion">只读</button><br>
        可修改计算版本号：<span>{{ resetVersionVal }}</span><br>
        <button @click="restVersion">重置版本</button><br>
        <!-- vue2计算版本号：<span>{{ version2 }}</span> -->
        版本号：<span>V{{ main }}.{{ minor }}</span>
        <hr>
        方法版本号：<span>{{ getVersion() }}</span>
        方法版本号：<span>{{ getVersion() }}</span>

    </div>

</template>
<!-- 
<script lang="ts">
export default {
    computed:{
        version2(){
            console.log(this)
            // 获取不到 main的值
            return "V" + this.main + "." + this.minor
        }
    }
}
</script>
 -->

<script lang="ts" setup  name="Version">
    import { reactive,toRefs,ref,computed } from 'vue';
    // let version = {main:1 ,minor:4}  TODO 为啥绑定不了对象
    // let {main,minor} = toRefs(version);

    let main = ref("1")
    let minor = ref("4")
    // version变量是 只读的
    let version = computed(()=>{
        console.log("computed version execute")
        return "V" + main.value + "." + minor.value
    });
    let resetVersionVal = computed({
        get(){
            return "V" + main.value + "." + minor.value
        },
        set(val){
            console.log("resetVersion set ", val)
            if(val.startsWith("V") || val.startsWith("v") ){
                val = val.slice(1)
            }
            const [str1,str2] = val.split(".")
            main.value = str1
            minor.value = str2
        }
    });
    // 方法会多次执行，计算属性有缓存
    function getVersion(){
        console.log("getVersion() execute")
        return "V" + main.value + "." + minor.value
    }
    function changeVersion(){
        // console 告警 属性是只读的
        version.value = "随机版本"
    }
    function restVersion(){
        // console 告警 属性是只读的
        resetVersion.value = "V1.0"
    }

</script>

<style scoped>
    .version{
    background-color: rgb(98, 191, 231);
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
    margin: 10px 0px;

}

input {
    margin: 0 5px;
    display: block;
}
</style>