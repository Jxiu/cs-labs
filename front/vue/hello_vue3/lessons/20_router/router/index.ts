// 创建一个路由器并暴露出去
import { createRouter,createWebHistory } from "vue-router";
import Home from "@/components/Home.vue";
import News from "@/components/News.vue";
import About from "@/components/About.vue";
const router = createRouter({
    history: createWebHistory(), //路由器工作模式
    routes:[
        {
            path: '/home',
            component: Home
        },
        {
            path: '/news',
            component: News
        },
        {
            path: '/about',
            component: About
        },
    ]
})

export default router