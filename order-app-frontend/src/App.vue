<template>
    <div class="container">
       <Header></Header> 
        <router-view></router-view>
        <Footer></Footer>
    </div>
</template>

<script>
    import OrdersList  from './components/orders/OrdersList.vue'
    import Header  from './components/shared/Header.vue'
    import Footer  from './components/shared/Footer.vue'
    import { mapMutations } from 'vuex'
    import { mapActions } from 'vuex'

    export default {
  
         components: { 
            Header,
            Footer
        }
    }
</script>

<style lang="css">
 
.container {
   margin-top: 20px;
  padding: 2px 16px; 
}
</style>
