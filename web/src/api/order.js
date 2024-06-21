import request from "@/utils/request.js";

export const orderCreateService = (data) => {
    return request.post('/order', data)
}

export const orderPageService = (data) => {
    return request.get('/order', data)
}

export const orderFindService = (id) => {
    return request.get(`/order/${id}`)
}

export const orderUpdateService = (id, data) => {
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