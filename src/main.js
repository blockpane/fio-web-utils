import Vue from 'vue'
import App from './App.vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootswatch/dist/litera/bootstrap.min.css';
import 'bootstrap-vue/dist/bootstrap-vue.css'

import { store } from './store/store.js'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  store,
}).$mount('#app')
