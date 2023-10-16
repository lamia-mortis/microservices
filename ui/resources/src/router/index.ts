import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import MainPage from "../views/MainPage.vue";
import AboutPage from "../views/AboutPage.vue";
import CaruselComponent from "../components/CaruselComponent.vue";
import BlogPage from "../views/BlogPage.vue";
import LocationPage from "../views/LocationPage.vue";
import TransactionsPage from "../views/TransactionsPage.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: MainPage,

    children: [
      { path: "/", component: CaruselComponent },
      {
        path: "aboutpage",
        component: AboutPage,
      },
      { path: "transactions", component: TransactionsPage },
      { path: "blog", component: BlogPage },
      { path: "location", component: LocationPage },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
