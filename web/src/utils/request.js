import {ElMessage} from 'element-plus'
//定制请求的实例

//导入axios  npm install axios
import axios from 'axios';
import {useTokenStore} from '@/stores/token.js'
import router from '@/router/index.js'
//定义一个变量,记录公共的前缀  ,  baseURL
const baseURL = 'http://localhost:5173/api';
const instance = axios.create({baseURL})

// 添加请求拦截器
instance.interceptors.request.use(
    (config)=>{
        // 请求前的回调
        const tokenStore = useTokenStore()
        if(tokenStore.token && tokenStore.token!=='') {
            config.headers.Authorization = tokenStore.token
        }
        return config
    },
    (err)=>{
        // 请求错误时的回调
        return Promise.reject(err)
    }
)

//添加响应拦截器
instance.interceptors.response.use(
    result=>{
        return result;
    },
    err=>{
        // 判断响应状态码，如果为401，则未登录，跳转到登录页面
        if(err.response.status === 401) {
            ElMessage.error('请先登录')
            router.push('/login')
        } else {
            if (!err.response.data || err.response.data === "") {
                err.response.data = "未知错误"
            }
            ElMessage.error(err.response.data)
        }
        return Promise.reject(err);//异步的状态转化成失败的状态
    }
)

export default instance;