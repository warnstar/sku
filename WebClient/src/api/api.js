import axios from 'axios';
import { Message } from 'element-ui';
import { route } from '../main';

axios.defaults.baseURL = "http://" + window.location.hostname + ":810";

axios.interceptors.request.use(
    config => {
        config.headers.Authorization = "Bearer " + sessionStorage.getItem('Authorization');

        return config;
    },
    err => {
        return Promise.reject(err);
    });

axios.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 401:
                    // 返回 401 清除token信息并跳转到登录页面
                    Message.error('无权限访问');

                    route.push("/login");
            }
        }
        return Promise.reject(error)   // 返回接口返回的错误信息
    });


export const requestLogin = params => { return axios.post(`/admin/admin/login`, params)};

export const getAppRelease = params => { return axios.get(`/admin/app-release`, { params: params }); };

export const getTsiStatus = params => { return axios.get(`/control/tsi-check`, { params: params }); };
