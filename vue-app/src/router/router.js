import vueRouter from "vue-router";
import Vue from "vue";
import Home from "../views/Home";
import Favorites from "../views/Favorites";
import Watch from "../views/Watch";

Vue.use(vueRouter);

const router = new vueRouter({
  mode: "history",
  routes: [
    {path: "/" , component: Home},
    {path: "/home", component: Home},
    {path: "/favorites/:userid", component: Favorites},
    {path: "/watch", component: Watch}
  ]
});

export default router;
