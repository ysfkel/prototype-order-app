import Vue from 'vue'
import Vuex from 'vuex'
import { getOrders } from '../services/order'

Vue.use(Vuex)

let order = {

}
 

export const store = new Vuex.Store({
    state: {
        orders: [],
        totalAmount: 0,
        pager: { 
            currentPage: 1,
            pageCount: 4,
            offset: 0  
        },
        dateRange: {
            startDate: null,
            endDate: null
        }, 
        search: null
    },
    getters: { 
        orders: (state) => {
            return {
                orders: [...state.orders],
                totalAmount: state.totalAmount
            }
        },
        pager: (state) => {
            return {
                ...state.pager
            }
        }, 
        searchFields: (state) => {
            return {
                dateRange: {
                   ...state.dateRange,
                   search:state.search
                }
            }
        }
    }, mutations:{
        setOrders: (state, orders) => { 
 
            let ordersList = orders.Orders || []
            state.orders = [...ordersList]
            state.totalAmount = orders.GrandTotalAmount
        },  
        updateSearchParam: (state, search) => {
      
            state.search = search
        }, 
        updateDateRange: (state, dateRange) => {
 
              state.dateRange = {
                 ...dateRange
             }
        }
    }, actions: {
        getOrders: async ({ commit, state }, payload) => {

        

            const orderListResult = await getOrders({
                startDate: state.dateRange.startDate,
                endDate: state.dateRange.endDate,
                search: state.search,
                offset: state.pager.offset,
                pageCount:state.pager.pageCount
            })  
             
            commit('setOrders',orderListResult)
        },  
        updatePage: async ({ commit, state }, {currentPage}) => {

 
            let offset=(currentPage - 1) *   state.pager.pageCount


            state.pager = {
               pageCount: state.pager.pageCount,
                currentPage: currentPage,
                offset: offset
            }

            const orderListResult = await getOrders({
                startDate: state.dateRange.startDate,
                endDate: state.dateRange.endDate,
                search: state.search,
                offset: state.pager.offset,
                pageCount:state.pager.pageCount
            })  

            commit('setOrders',orderListResult)

       },
    }
})