<template>

</template>

<script lang="ts" setup name="Demo">
import { type Person,type PersonArray,type PersonList } from '@/types';
    let devloper:Person = {name:"john",age:10,id:"1"}
    let devlopers:Array<Person>= [
        {name:"john",age:10,id:"1"},
        {name:"sim",age:11,id:"2"},
        {name:"jojo",age:7,id:"3"}
    ] 
    let devloperArray : PersonArray= [
        {name:"john",age:10,id:"1"},
        {name:"sim",age:11,id:"2"},
        {name:"jojo",age:7,id:"3"}
    ] 
    let devloperList: PersonList= [
        {name:"john",age:10,id:"1"},
        {name:"sim",age:11,id:"2"},
        {name:"jojo",age:7,id:"3"}
    ] 
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

</style>