<template>
  <div class="FreshDataButton">
    <button @click="FreshData()">Fresh Data</button>
  </div>
</template>

<script lang="ts" setup>
import {GetSchedulerInfo} from "@/network/Manager";
import {ElMessage, ElNotification} from "element-plus";
import {SetupServersStore} from '@/stores/SetupServersStore'

let store = SetupServersStore()
let timer: any = null

//刷新数据
function FreshData() {
  if (timer) {
    clearTimeout(timer)
  }
  //设置定时器防止刷新过快
  timer = setTimeout(() => {
    ElNotification({
      type: 'success',
      title: 'Fresh Data',
      message: 'success',
      duration: 2000,
    })
    GetSchedulerInfo().then((res) => {
      const status = new Map()
      status.set('host_info', res.data.host_info)
      status.set('cpu_info', res.data.cpu_info)
      status.set('mem_info', res.data.mem_info)
      status.set('disk_info', res.data.disk_info)
      status.set('net_info', res.data.net_info)
      status.set('pool_info', {
        core_num: res.data.pool_core_num,
        max_num: res.data.pool_max_num,
        activate_num: res.data.pool_activate_num,
        job_num: res.data.pool_job_num,
      })
      store.SetSystemStatus(status)
    }, (err) => {
      ElMessage({
        message: 'loading data error: ' + err,
        type: 'error',
        duration: 1000,
      })
    })
  }, 1000)
}
</script>

<style scoped>
button {
  padding: 0.8em 1.8em;
  border: 2px solid #48cae4;
  position: relative;
  overflow: hidden;
  background-color: transparent;
  text-align: center;
  text-transform: uppercase;
  font-size: 16px;
  transition: .3s;
  z-index: 1;
  font-family: inherit;
  color: #48cae4;
}

button::before {
  content: '';
  width: 0;
  height: 300%;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) rotate(45deg);
  background: #48cae4;
  transition: .5s ease;
  display: block;
  z-index: -1;
}

button:hover::before {
  width: 105%;
}

button:hover {
  color: #111;
  cursor: pointer;
}

.FreshDataButton {
  scale: 0.8;
}

button:active {
  transform: scale(0.9);
  transition: all 0.2s ease-in-out;
}

</style>