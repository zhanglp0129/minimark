import request from "@/utils/request.js";

export const procurementCreateService = (data) => {
    return request.post('/procurement', data)
}

export const procurementPageService = (data) => {
    return request.get('/procurement', data)
}

export const procurementFindService = (id) => {
    return request.get(`/procurement/${id}`)
}

export const procurementUpdateService = (id, data) => {
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