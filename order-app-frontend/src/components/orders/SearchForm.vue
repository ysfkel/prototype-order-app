<template>
 
   <form  >
       <div class="search-form field-collection">
         <button @click.stop.prevent="onSubmit" class="btn btn-light form-item">Search</button>   
           <input class="form-control mb-4 form-item"  type="text" v-on:keyup="updateSearchParam" placeholder="Search" v-model="search">
       </div>
  
 
       <div class="field-collection">
       
        <date-range-picker
                ref="picker"
                :opens="center"
                :locale-data="{ firstDay: 1, format: 'DD-MM-YYYY HH:mm:ss' }"
                :ranges= "false"
                :showDropdowns="true"
                :autoApply="true"
                v-model="dateRange"
                @update="updateDateRange"
                :linkedCalendars="false">
   

          <template v-slot:input="picker" style="min-width: 350px;">
               <div  style="min-width: 350px;"> {{ picker.startDate | moment("dddd, MMMM Do YYYY")  }} - {{ picker.endDate | moment("dddd, MMMM Do YYYY")  }}</div>
            </template>
           </date-range-picker>     <button @click.prevent="clearDate" class="btn btn-secondary">clear</button>   
         </div>
     
 
    <div>

    </div>
</form> 
 
</template>

<script>

    import { mapActions } from 'vuex'
    import DateRangePicker from 'vue2-daterange-picker' 
    import 'vue2-daterange-picker/dist/vue2-daterange-picker.css'

    export default {
 
         data: function(){

              return {
                  search:'',
                  dateRange : {
                      startDate: null,
                      endDate: null
                  }
              }
             
         },
         components: { DateRangePicker },
          methods: {
            ...mapActions(['getOrders']),

            onSubmit () { 
                    this.getOrders()
              },

              updateDateRange(values) {

                   let startDate = values.startDate.toISOString()
                   let endDate = values.endDate.toISOString()
                    this.$store.commit('updateDateRange',{
                           startDate,
                           endDate
                    })
        
              }, updateSearchParam(values) {
 
                    this.$store.commit('updateSearchParam',this.search)
        
              },

              clearDate() {
                    this.dateRange.startDate = null 
                    this.dateRange.endDate = null
              },
              getOrders() {
                     this.$store.dispatch('getOrders', {
                         search: this.search,
                        ...this.dateRange,
                  })
              }
              
          }, created() {
             this.getOrders()
          }
    }
</script>

<style scoped>
  .search-form .form-item {
      display: inline-block;
   }


   .search-form input { 
    width: 90%;
   }

   .field-collection {
       margin: 8px 0;
   }
  
</style>
