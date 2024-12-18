import { baseRequest, request } from '@/utils/http'
// 默认导出

export default {
    upload: (data) => request.post('/upload', data, { needToken: true }),
    updateUser: (data) => request.put('/userinfo', data, { needToken: true }),
    getUser: () => request.get('/user', { needToken: true }),

    loginService:(data) => baseRequest.post('/login', data),
    registerService:(data) => baseRequest.post('/register', data),
    logout:()=> baseRequest.post('/logout'),

    searchProduct: (params={}) => baseRequest.get('/products',  { params }),
    getPageList: (data) => baseRequest.get('/page-list', data),
    getProductRe: (params={}) => baseRequest.get(`/product`,{ params } ),
    getProductDetail: (params={}) => baseRequest.get(`/product-details`,{ params }),
    getProductImages: (params={}) => baseRequest.get(`/product-images`,{ params }),
};
