import {acceptHMRUpdate, defineStore} from 'pinia'
import {ref} from 'vue';

export const SetupServersStore = defineStore('SetupServersStore', () => {

    //系统状态
    const SystemStatus = ref(new Map());

    //获取系统状态
    function GetSystemStatus() {
        return SystemStatus.value
    }

    //设置系统状态
    function SetSystemStatus(status: Map<string, any>) {
        SystemStatus.value = status
    }

    return {
        GetSystemStatus, SetSystemStatus,
    }
})

// 热更新
if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(SetupServersStore, import.meta.hot))
}