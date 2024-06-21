import request from "@/utils/request.js";

export const categoryCreateService = (data) => {
    // json格式
    return request.post('/category', data)
}

export const categoryListService = () => {
    return request.get('/category')
}

export const categoryFindService = (id) => {
    return request.get(`/category/${id}`)
}

export const categoryUpdateService = (id, data) => {
    return request.put(`/category/${id}`, data)
}

export const categoryDeleteService = (id) => {
    return request.delete(`/category?id=${id}`, id)
}