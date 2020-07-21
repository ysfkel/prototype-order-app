import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import VueRouter from 'vue-router'
import { routes } from './routes'
import { store } from './store/store'
import VueMoment from 'vue-moment'


Vue.use(VueResource)
Vue.use(VueRouter)
Vue.use(VueMoment);


const router = new VueRouter({
  routes,
  mode: 'history'
})

Vue.http.options.root = "http://localhost:5000/v1/api"

new Vue({
  el: '#app',
  router: router,
  store,
  render: h => h(App)
})
