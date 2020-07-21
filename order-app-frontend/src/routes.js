import Home from './components/home/Home.vue'
import Order from './components/orders/Order.vue'

export const routes = [
    {path: '', component: Home},
    {path: '/orders', component: Order}
]