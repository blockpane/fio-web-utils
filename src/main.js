import Vue from 'vue'
import App from './App.vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import VueClipboard from 'vue-clipboard2'
import './plugins/table.js'

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootswatch/dist/litera/bootstrap.min.css';
import 'bootstrap-vue/dist/bootstrap-vue.css'

import { store } from './store/store.js'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(VueClipboard)

import router from './router'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
