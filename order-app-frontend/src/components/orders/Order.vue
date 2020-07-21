<template>
    <div class="container"> 
       <div class="row">
             <SearchForm></SearchForm>
        </div> 
        <div class="row">
             <ordersList></ordersList>
        </div> 
    </div>
</template>

<script>
    import OrdersList  from './OrdersList.vue' 
    import SearchForm  from './SearchForm.vue' 

    export default { 
         components: {
            ordersList:OrdersList,
            SearchForm
        }
    }
</script>

<style lang="css">
 
.row {
  margin: 30px 0; 
}
</style>