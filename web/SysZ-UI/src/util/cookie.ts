import { Base64 } from 'js-base64'

class Account {
    constructor(account?: string, password?: string) {
        if (account && password) {
            this.account = account
            this.password = password
            this.BOOL = true
        } else {
            this.BOOL = false
        }
    }
    public account?: string;
    public password?: string;
    public BOOL: boolean;
}

const cookie_username = "themis-username";
const cookie_password = "themis-password";

/**cookie操作 */
function setCookie(account: string, password: string, days: number) {
    if (days === 0) {
        window.document.cookie =
            Base64.encode(cookie_username) + "=" + Base64.encode(account) + ";max-age=" + "0";
        window.document.cookie =
            Base64.encode(cookie_password) + "=" + Base64.encode(password) + ";max-age=" + "0";
        return
    }
    const date = new Date();
    date.setTime(date.getTime() + 24 * 60 * 60 * 1000 * days);
    window.document.cookie =
        Base64.encode(cookie_username) + "=" + Base64.encode(account) + ";max-age=" + date.toString();
    window.document.cookie =
        Base64.encode(cookie_password) + "=" + Base64.encode(password) + ";max-age=" + date.toString();
}

// 读取cookie 将用户名和密码回显到input框中
function getCookie(cookie: string): Account {
    let account: string = "";
    let password: string = "";
    if (cookie.length > 0) {
        const arr = cookie.split("; ");
        for (let i = 0; i < arr.length; i++) {
            const arr2 = arr[i].split("=");
            if (arr2[0] === Base64.encode(cookie_username)) {
                account = Base64.decode(arr2[1]);// base64解密
            }
            if (arr2[0] === Base64.encode(cookie_password)) {
                password = Base64.decode(arr2[1]);// base64解密
            }
        }
    }
    if (account != "" && password != "") {
        return new Account(account, password);
    }
    return new Account();
}

export { setCookie, getCookie };