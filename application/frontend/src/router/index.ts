import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";
import CarTires from "../views/CarTires.vue";
import CarBody from "../views/CarBody.vue";
import CarInterior from "../views/CarInterior.vue";
import CarStore from "../views/CarStore.vue";
import CarInsure from "../views/CarInsure.vue";
import CarMaint from "../views/CarMaint.vue";
import CarMenu from "../views/CarMenu.vue";
import CarRecord from "../views/CarRecord.vue";

const routes: RouteRecordRaw[] = [
    {path:'/', redirect: '/home'},
    {path: '/home', name: "home", component: HomeView},
    {path: '/login', name: "login", component: LoginView},
    {path: '/register', name: "register", component: RegisterView},
    {path: '/carTires', name: "carTires", component: CarTires},
    {path: '/carBody', name: "carBody", component: CarBody},
    {path: '/carInterior', name: "carInterior", component: CarInterior},
    {path: '/carStore', name: "carStore", component: CarStore},
    {path: '/carInsure', name: "carInsure", component: CarInsure},
    {path: '/carMaint', name: "carMaint", component: CarMaint},
    {path: '/carMenu', name: "carMenu", component: CarMenu},
    {path: '/carRecord', name: "carRecord", component: CarRecord},
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;