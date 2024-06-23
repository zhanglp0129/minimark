import request from "@/utils/request.js";
import moment from "moment/moment.js";

export const procurementCreateService = (data) => {
    if(data.payTime !== null) {
        data.payTime = moment(data.payTime).format('YYYY-MM-DD HH:mm:ss')
    }
    if(data.purchaseTime !== null) {
        data.purchaseTime = moment(data.purchaseTime).format('YYYY-MM-DD HH:mm:ss')
    }
    return request.post('/procurement', data)
}

export const procurementPageService = (data) => {
    let uri = `/procurement?page_num=${data.pageNum}&page_size=${data.pageSize}`

    if(data.payMethodId && data.payMethodId !== '') {
        uri += `&pay_method_id=${data.payMethodId}`
    }
    if(data.purchaseTimeRange !== null && data.purchaseTimeRange.length > 0) {
        uri += `&min_purchase_time=${moment(data.purchaseTimeRange[0]).format('YYYY-MM-DD HH:mm:ss')}`
        uri += `&max_purchase_time=${moment(data.purchaseTimeRange[1]).format('YYYY-MM-DD HH:mm:ss')}`
    }
    if(data.payTimeRange !== null && data.payTimeRange.length > 0) {
        uri += `&min_pay_time=${moment(data.payTimeRange[0]).format('YYYY-MM-DD HH:mm:ss')}`
        uri += `&max_pay_time=${moment(data.payTimeRange[1]).format('YYYY-MM-DD HH:mm:ss')}`
    }
    if(data.totalPayRange !== null && data.totalPayRange.length > 0) {
        uri += `&min_total_pay=${data.totalPayRange[0]}`
        uri += `&max_total_pay=${data.totalPayRange[1]}`
    }
    return request.get(uri)
}

export const procurementFindService = (id) => {
    return request.get(`/procurement/${id}`)
}

export const procurementUpdateService = (id, data) => {
    if(data.payTime !== null) {
        data.payTime = moment(data.payTime).format('YYYY-MM-DD HH:mm:ss')
    }
    if(data.purchaseTime !== null) {
        data.purchaseTime = moment(data.purchaseTime).format('YYYY-MM-DD HH:mm:ss')
    }
    return request.put(`/procurement/${id}`, data)
}

export const procurementAddGoodsService = (id, data) => {
    return request.post(`/procurement/add-goods/${id}`, data)
}

export const procurementUpdateGoodsService = (id, data) => {
    return request.patch(`/procurement/update-goods/${id}`, data)
}

export const procurementDeleteService = (id, changeStock) => {
    return request.delete(`/procurement?id=${id}&change_stock=${changeStock}`)
}