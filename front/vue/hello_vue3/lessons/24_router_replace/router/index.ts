// 创建一个路由器并暴露出去
import { createRouter,createWebHistory,createWebHashHistory } from "vue-router";
import Home from "@/views/Home.vue";
import News from "@/views/News.vue";
import About from "@/views/About.vue";
import Detail from "@/components/Detail.vue";
const router = createRouter({
    // 1.history工作模式：展示的url中不带#，和传统网站相同
    // 1.1 后期项目线上需要配合路径处理，否则会出现404 
    history: createWebHistory(), //路由器工作模式
    // 2.hash工作模式：展示的url中带# 如：http://localhost:5173/news#/news
    // 2.1 SEO搜索相当差
    // history: createWebHashHistory(),
    routes:[
        {
            name:"首页",
            path: '/home',
            component: Home
        },
        {
            name:"xinwen",
            path: '/news',
            component: News,
            children:[
                {
                    name:'xiang',
                    path: 'detail',
                    component: Detail,
                    // 1. 将路由收到params参数做为组件的props传给组件（只支持params类型传参，）
                    // path: 'detail/:id/:title/:content?', // ?非必填
                    // props: true,
                    // 2.函数写法，可以自己决定将什么作为props传给组件（支持query类型传参）
                    props(route){
                        return route.query
                    }
                    // 3. 固定传参，直接使用对象传参
                    // props:{id:1,title:'hello',content:'world'}
                }
            ]
        },
        {
            name:"guanyu",
            path: '/about',
            component: About
        },
    ]
})

export default router