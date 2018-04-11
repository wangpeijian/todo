// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App'
import router from './router'

Vue.config.productionTip = false
Vue.use(ElementUI);

/*---------------------------------------------------*/
//判断是否是生产环境，生产环境下，将console方法重写
if(process.env.NODE_ENV === 'production'){
    window.console.log = ()=>{};
    window.console.info = ()=>{};
    window.console.error = ()=>{};
    window.console.group = ()=>{};
    window.console.groupEnd = ()=>{};
    Vue.config.silent = true;
}

//全局安装promise
import {install} from 'promise-es6'
install();

//全局安装fetch
import 'whatwg-fetch'

//安装自定义组件
import Plugin from './script/Plugin'
Vue.use(new Plugin());

//应用全局样式
import './styles/public.scss';
/*---------------------------------------------------*/

new Vue({
    el: '#app',
    router,
    template: '<App/>',
    components: {App}
});
