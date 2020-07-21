import Vue from 'vue'


export const getOrders = async ({search,startDate,endDate,offset,pageCount}) => {
  
        try {

        let query = ''
        if (search) {
            query = addQueryParam('search',search,query )
        }

        if (startDate) {
            query = addQueryParam('start_date',startDate,query )
        }

        if (endDate) {
            query = addQueryParam('end_date',endDate,query )
        } 

 
        query = addQueryParam('off_set',offset,query )
    

        query = addQueryParam('page_count',pageCount,query )
        
         const result = await Vue.http.get('orders'+query)
 
         return result.body.Data
        
       } catch(e) {
           return e
       }

}


let addQueryParam = (paramName, value,query) => {

    if(!query) {
        query='?'+paramName+'='+value
    } else {
        query=query+'&'+paramName+'='+value
    }
    return query
}
