<template>

  <paginate
  :page-count="pager.pageCount"
  :click-handler="changePage"
  :prev-text="'Prev'"
  :next-text="'Next'"
  :container-class="'pagination'">
</paginate>
 
</template>
<script>
    import { mapGetters ,mapActions} from 'vuex'
    import Paginate from 'vuejs-paginate'

    export default { 

        components: { Paginate },
        computed: {
            ...mapGetters([
            'pager'
          ]) 

        }, methods: {
 
            changePage(currentPage) {
          
                this.$store.dispatch('updatePage', {
                        currentPage
                 })
            }, 
        }
    }
</script>

<style lang="css">
 
.pagination {
  display: inline-block;
}

.pagination a {
  color: black;
  float: left;
  padding: 8px 16px;
  text-decoration: none;
}
</style>