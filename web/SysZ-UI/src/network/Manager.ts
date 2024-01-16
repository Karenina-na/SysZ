import service from "@/network/request";
import {Base64} from "js-base64";

//登录
export function LoginNetwork(account: string, password: string) {
    const username_enc = Base64.encode(account);
    const password_enc = Base64.encode(password);
    const config = {
        method: 'post',
        url: '/login',
        data: {
            "account": username_enc,
            "password": password_enc
        }
    }
    return service(config);
}

//获取系统信息
export function GetSchedulerInfo() {
    const config = {
        method: 'get',
        url: '/operator/getAll'
    }
    return service(config)
}