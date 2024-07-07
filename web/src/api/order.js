import request from "@/utils/request.js";
import moment from 'moment'

export const orderCreateService = (data) => {
    if(data.payTime !== null) {
        data.payTime = moment(data.payTime).format('YYYY-MM-DD HH:mm:ss')
    }
    return request.post('/order', data)
}

export const orderPageService = (data) => {
    let uri = `/order?page_num=${data.pageNum}&page_size=${data.pageSize}`

    if(data.payMethodId && data.payMethodId !== '') {
        uri += `&pay_method_id=${data.payMethodId}`
    }
    if(data.payTimeRange !== undefined && data.payTimeRange !== null && data.payTimeRange.length > 0) {
        uri += `&min_pay_time=${moment(data.payTimeRange[0]).format('YYYY-MM-DD HH:mm:ss')}`
        uri += `&max_pay_time=${moment(data.payTimeRange[1]).format('YYYY-MM-DD HH:mm:ss')}`
    }
    if(data.totalPayRange !== undefined && data.totalPayRange !== null && data.totalPayRange.length > 0) {
        uri += `&min_total_pay=${data.totalPayRange[0]}`
        uri += `&max_total_pay=${data.totalPayRange[1]}`
    }
    return request.get(uri)
}

export const orderFindService = (id) => {
    return request.get(`/order/${id}`)
}

export const orderUpdateService = (id, data) => {
    if(data.payTime !== null) {
        data.payTime = moment(data.payTime).format('YYYY-MM-DD HH:mm:ss')
    }
    return request.put(`/order/${id}`, data)
}

export const orderAddGoodsService = (id, data) => {
    return request.post(`/order/add-goods/${id}`, data)
}

export const orderUpdateGoodsService = (id, data) => {
    return request.patch(`/order/update-goods/${id}`, data)
}

export const orderDeleteService = (id, changeStock) => {
    return request.delete(`/order?id=${id}&change_stock=${changeStock}`)
}