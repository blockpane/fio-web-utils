import Vue from 'vue'
import Router from 'vue-router'
import Names from "@/components/Names";
import FeeVote from "@/components/Fees";
import Tables from "@/components/Tables";

Vue.use(Router)

export default new Router({
    //mode: 'history',
    //base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'home',
            component: Names
        },
        {
            path: '/names',
            name: 'names',
            component: Names
        },
        {
            path: '/fees',
            name: 'fees',
            component: FeeVote
        },
        {
            path: '/tables',
            name: 'tables',
            component: Tables
        },
    ]
})
