<template>
 <div>
  <span>Total Amount: {{orders.totalAmount}}</span>
  <table class="table table-bordered  ">
     <thead>
        <tr>
            <td>Order Name</td>
           <td>Company Name</td>
            <td>Customer Name</td>
            <td>Order Date</td>
            <td>Delivered Amount</td>
            <td>Total Amount</td>
        </tr>
    </thead>
    <tbody id="myTable">
     
      <tr v-for="order in orders.orders" :key="order.id">
            <td>{{order.order_name}}</td>
            <td>{{order.company_name}}</td>
            <td>{{order.customer_name}}</td>
            <td>{{order.created_at}}</td>
            <td>{{order.delivered_amount}}</td>
            <td>{{order.total_amount}}</td> 
        </tr>
    </tbody>
  </table>
  <Pager></Pager>

 
 </div>
</template>
<script>
    import { mapGetters ,mapActions} from 'vuex'
    import Pager  from './Pager.vue' 

    export default { 

        components: { Pager },
        computed: {
            ...mapGetters([
            'orders', 
          ]) 

        }
    }
</script>
 