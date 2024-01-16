import { ElLoading } from 'element-plus'

export function useLoading() {
    const loading = ElLoading.service({
        lock: true,
        text: 'Loading',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
    })

    const endLoading = () => {
        loading.close()
    }
    return {
        endLoading
    }
}