import axios from "axios";
import {useGlobalStore} from "@/stores/GlobalStore";
import {Base64} from 'js-base64';

const globalStore = useGlobalStore();

const service = axios.create({
    baseURL: (process.env.VITE_NODE_MODE === 'production' ? process.env.VITE_THEMISE_URL + "" : "/api") + "/v1",
    timeout: 5000
});

// Request interceptors
service.interceptors.request.use(config => {
    const message = globalStore.getRootMessage();
    const username_enc = Base64.encode(message.username);
    const password_enc = Base64.encode(message.password);
    config.headers = {
        "Content-Type": "application/json",
        "account": username_enc,
        "password": password_enc,
        "token": message.token
    };
    return config;
}, error => {
    return Promise.reject(error);
});

// Response interceptors
service.interceptors.response.use(response => {
    const res = response.data;
    if (response.status != 200) {
        return Promise.reject(new Error(res.message || "network error"));
    }
    if (res.code != 20010) {
        return Promise.reject(new Error(res.message || "logic error"));
    }
    return res;
}, error => {
    return Promise.reject(error);
});

export default service;