import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/', // 根路径
        component: () => import('../views/Home.vue'), // 默认显示 Home 页面
    },
    {
        path: '/product/:id', // 动态路由，用于显示商品详情
        name: 'ProductDetail',
        component: () => import('../views/ProductDetail.vue'), // 懒加载 ProductDetail 组件
    },
    {
        path:'/search-results',
        name: 'SearchResults',
        component: () => import('../views/SearchResult.vue'), // 懒加载 ProductDetail 组件
    },
    {
        path:'/login',
        name: 'login',
        component: () => import('@/views/Login.vue'), // 懒加载 ProductDetail 组件
    },
    {
        path:'/register',
        name: 'register',
        component: () => import('@/views/Register.vue'), // 懒加载 ProductDetail 组件
    },
    {
        path:'/user',
        name: 'user',
        component: () => import('@/views/User.vue'), // 懒加载 ProductDetail 组件
    },
];

const router = createRouter({
    history: createWebHistory(), // 使用 HTML5 History API
    routes,
});

export default router;