import axios from "axios";
import {useGlobalStore} from "@/stores/GlobalStore";
import {Base64} from 'js-base64';

const globalStore = useGlobalStore();

const service = axios.create({
    // baseURL: (process.env.NODE_ENV === 'production' ? process.env.VITE_SYSZ_URL : "/api") + "/v1",
    baseURL: import.meta.env.VITE_SYSZ_URL + "/v1",
    timeout: 5000
});

// Request interceptors
service.interceptors.request.use(config => {
    const message = globalStore.getRootMessage();
    const username_enc = Base64.encode(message.username);
    const password_enc = Base64.encode(message.password);

    config.headers.set("Content-Type", "application/json")
    config.headers.set("account", username_enc)
    config.headers.set("password", password_enc)
    config.headers.set("token", message.token)

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