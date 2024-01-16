import {acceptHMRUpdate, defineStore} from 'pinia'
import {ref} from 'vue'
//全局可用信息
export const useGlobalStore = defineStore('GlobalStore', () => {

    const token = ref("")
    const username = ref('')
    const password = ref('')

    //更新token
    function updateToken(t: string) {
        token.value = t
    }

    //更新用户名
    function updateUsername(u: string) {
        username.value = u
    }

    //更新密码
    function updatePassword(p: string) {
        password.value = p
    }

    //查找所有数据
    function getRootMessage() {
        return {
            token: token.value,
            username: username.value,
            password: password.value
        }
    }

    const Theme = ref('')

    //设置全局主题
    function setTheme(val: string) {
        Theme.value = val;
    }

    //获取全局主题
    function getTheme() {
        return Theme.value;
    }

    return {
        updateToken, updateUsername, updatePassword, getRootMessage,
        setTheme, getTheme
    }
})

// 热更新
if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useGlobalStore, import.meta.hot))
}