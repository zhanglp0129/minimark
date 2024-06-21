import request from "@/utils/request.js";
import {id} from "element-plus/es/locale/index";

export const payMethodCreateService = (data) => {
    // json格式
    return request.post('/pay_method', data)
}

export const payMethodListService = () => {
    return request.get('/pay_method')
}

export const payMethodFindService = (id) => {
    return request.get(`/pay_method/${id}`)
}

export const payMethodUpdateService = (id, data) => {
    return request.put(`/pay_method/${id}`, data)
}

export const payMethodDeleteService = (id) => {
    return request.delete(`/pay_method?id=${id}`)
}