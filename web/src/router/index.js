import Login from "@/views/Login.vue";
import {createRouter, createWebHistory} from "vue-router";
import Admin from "@/views/Admin.vue";
import Category from "@/views/Category.vue";
import PayMethod from "@/views/PayMethod.vue";
import Dashboard from "@/views/Dashboard.vue";
import Goods from "@/views/Goods.vue";
import Order from "@/views/Order.vue";
import Procurement from "@/views/Procurement.vue";

const routes = [
    {path: '/login', component: Login},
    {path: '/', redirect: '/admin'},
    {
        path: '/admin',
        component: Admin,
        redirect: '/admin/dashboard',
        children: [
            {path: 'dashboard', component: Dashboard},
            {path: 'category', component: Category},
            {path: 'pay-method', component: PayMethod},
            {path: 'goods', component: Goods},
            {path: 'order', component: Order},
            {path: 'procurement', component: Procurement}
        ]
    }
]

const router = createRouter(
    {
        routes: routes,
        history: createWebHistory()
    }
)

export default router