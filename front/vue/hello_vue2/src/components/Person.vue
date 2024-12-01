<template>
    <div class="person">
        <h2>当前和：{{ sum }}</h2>
        <button @click="add">+1</button>
    </div>
  </template>
  
  <script>
  export default {
    name: "Person",
    data(){
        return{
            sum:0
        }
    },
    methods:{
        add(){
            this.sum +=1
        }
    }
  }
  </script>
  
  <style scoped>
    .person {
        background-color: skyblue;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0 0 10px;
    }
  
  </style>
  