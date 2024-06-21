import Login from "@/views/Login.vue";
import {createRouter, createWebHistory} from "vue-router";
import Admin from "@/views/Admin.vue";
import Category from "@/views/Category.vue";
import PayMethod from "@/views/PayMethod.vue";
import Dashboard from "@/views/Dashboard.vue";

const routes = [
    {path: '/login', component: Login},
    {path: '/', redirect: '/admin/dashboard'},
    {
        path: '/admin',
        component: Admin,
        redirect: '/admin/dashboard',
        children: [
            {path: 'dashboard', component: Dashboard},
            {path: 'category', component: Category},
            {path: 'pay-method', component: PayMethod},
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