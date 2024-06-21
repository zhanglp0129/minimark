import request from "@/utils/request.js";

export const userLoginService = (loginData) => {
    // x-www-form-urlencoded 编码格式
    const params = new URLSearchParams()
    for (let key in loginData) {
        params.append(key, loginData[key])
    }

    return request.post('/user/login', params)
}