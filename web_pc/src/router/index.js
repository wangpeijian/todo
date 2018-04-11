import Vue from 'vue'
import Router from 'vue-router'

const Desktop = resolve => require(['../view/Desktop'], resolve);
const Register = resolve => require(['../view/Register'], resolve);
const Login = resolve => require(['../view/Login'], resolve);

Vue.use(Router);

const router = new Router({
    routes: [{
        path: '/',
        component: Desktop
    },{
        path: '/register',
        component: Register
    }, {
        path: '*',
        component: Login
    }
    ]
});

export default router;
