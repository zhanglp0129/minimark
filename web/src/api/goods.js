import request from "@/utils/request.js";
import {id} from "element-plus/es/locale/index";

export const goodsCreateService = (data) => {
    // json格式
    return request.post('/goods', data)
}

export const goodsPageService = (data) => {
    let uri = `/goods?page_num=${data.pageNum}&page_size=${data.pageSize}`
    if(data.categoryId && data.categoryId !== '') {
        uri += `&category_id=${data.categoryId}`
    }
    if(data.goodsName && data.goodsName !== '') {
        uri += `&goods_name=${data.goodsName}`
    }
    return request.get(uri)
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
    idsString = idsString.substring(0, idsString.length - 1)
    return request.get(`/goods/${idsString}`)
}

export const goodsUpdateService = (id, data) => {

    return request.put(`/goods/${id}`, data)
}

export const goodsDeleteService = (id) => {
    return request.delete(`/goods?id=${id}`)
}