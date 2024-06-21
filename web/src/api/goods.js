import request from "@/utils/request.js";
import {id} from "element-plus/es/locale/index";

export const goodsCreateService = (data) => {
    // json格式
    return request.post('/goods', data)
}

export const goodsPageService = () => {
    return request.get('/goods')
}

export const goodsFindService = (ids) => {
    if(ids.length === 0) {
        // 参数个数为0，则返回空数组
        return request.get(`/goods/`)
    }
    let idsString = ''
    for (const id of ids) {
        idsString += id+','
    }
    idsString.substring(0, idsString.length - 1)
    return request.get(`/goods/${idsString}`)
}

export const goodsUpdateService = (id, data) => {
    return request.put(`/goods/${id}`, data)
}

export const goodsDeleteService = (id) => {
    return request.delete(`/goods/?id=${id}`)
}