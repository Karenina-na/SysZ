<template>
  <div class="CpuStatus_container">
    <div class="Left_container">
      <div id="Bytes_Chart" style="width: 500px"/>
    </div>
    <div class="Right_container">
      <div id="Packets_Chart" style="width: 500px"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, watch} from "vue";
import * as echarts from "echarts";
import {DarkTheme} from "@/assets/json/Net_echart/DarkTheme";
import {LightTheme} from "@/assets/json/Net_echart/LightTheme";
import {useGlobalStore} from "@/stores/GlobalStore";

const props = defineProps(['value'])
let Bytes_Chart: any = null;
let Packets_Chart: any = null;
const store = useGlobalStore()

//初始化
onMounted(() => {
  if (props.value) {
    let [bytes_sent, bytes_recv, packets_sent, packets_recv, net_name] = Create_Data()
    Bytes_chart(bytes_sent, bytes_recv, net_name)
    Packets_chart(packets_sent, packets_recv, net_name)
  } else {
    Bytes_chart([0], [0], ['Not Found'])
    Packets_chart([0], [0], ['Not Found'])
  }
});

//侦测器监听父组件传参
watch(props, () => {
  if (!props.value) {
    return
  }
  Bytes_Chart.dispose();
  Packets_Chart.dispose();
  let [bytes_sent, bytes_recv, packets_sent, packets_recv, net_name] = Create_Data()
  Bytes_chart(bytes_sent, bytes_recv, net_name)
  Packets_chart(packets_sent, packets_recv, net_name)
});

//监听Theme变化
watch(store.getTheme, () => {
  if (!props.value) {
    return
  }
  Bytes_Chart.dispose();
  Packets_Chart.dispose();
  let [bytes_sent, bytes_recv, packets_sent, packets_recv, net_name] = Create_Data()
  Bytes_chart(bytes_sent, bytes_recv, net_name)
  Packets_chart(packets_sent, packets_recv, net_name)
});

//绘制Bytes信息
const Bytes_chart = function (bytes_sent: Array<number>, bytes_recv: Array<number>, name: Array<string>) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    Bytes_Chart = echarts.init(document.getElementById("Bytes_Chart")!, 'LightTheme', {height: 322 * name.length - 22});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    Bytes_Chart = echarts.init(document.getElementById("Bytes_Chart")!, 'DarkTheme', {height: 322 * name.length - 22});
  }
  //bytes 信息
  let data = []
  for (let i = 0; i < name.length; i++) {
    data.push({
      name: 'bytes: ' + name[i],
      type: 'pie',
      radius: 80,
      center: ['50%', 200 + i * 322],
      stillShowZeroSum: true,
      data: [
        {value: bytes_sent[i], name: 'bytes sent'},
        {value: bytes_recv[i], name: 'bytes received'},
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    })
  }
  //Net名称
  let title = []
  title.push({
    text: 'bytes information',
    left: 'center',
    top: 15,
  })
  for (let i = 0; i < name.length; i++) {
    title.push({
      subtext: name[i],
      left: '50%',
      top: 50 + 322 * i,
      textAlign: 'center'
    })
  }
  Bytes_Chart.setOption({
    title: title,
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        return params.seriesName + "<br>" + params.marker + params.name + ":&nbsp;&nbsp;&nbsp;&nbsp;" + "<span style='font-weight: bold'>" + params.value + "</span>" + " MB";
      },
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      top: '30',
    },
    series: data,
  });
  window.onresize = function () {//自适应大小
    Bytes_Chart.resize();
  };
}

//绘制绘制Packets信息
const Packets_chart = function (packets_sent: Array<number>, packets_recv: Array<number>, name: Array<string>) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    Packets_Chart = echarts.init(document.getElementById("Packets_Chart")!, 'LightTheme', {height: 322 * name.length - 22});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    Packets_Chart = echarts.init(document.getElementById("Packets_Chart")!, 'DarkTheme', {height: 322 * name.length - 22});
  }
  //packet 信息
  let data = []
  for (let i = 0; i < name.length; i++) {
    data.push({
      name: 'packets: ' + name[i],
      type: 'pie',
      radius: 80,
      center: ['50%', 200 + i * 322],
      stillShowZeroSum: true,
      data: [
        {value: packets_sent[i], name: 'packets sent'},
        {value: packets_recv[i], name: 'packets received'},
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    })
  }
  //Net名称
  let title = []
  title.push({
    text: 'packets information',
    left: 'center',
    top: 15,
  })
  for (let i = 0; i < name.length; i++) {
    title.push({
      subtext: name[i],
      left: '50%',
      top: 50 + 322 * i,
      textAlign: 'center'
    })
  }
  Packets_Chart.setOption({
    title: title,
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        return params.seriesName + "<br>" + params.marker + params.name + ":&nbsp;&nbsp;&nbsp;&nbsp;" + "<span style='font-weight: bold'>" + params.value + "</span>" + " MB";
      },
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      top: '30',
    },
    series: data,
  });
  window.onresize = function () {//自适应大小
    Packets_Chart.resize();
  };
}

//构造数据
function Create_Data() {
  let bytes_sent = []
  let bytes_recv = []
  let packets_sent = []
  let packets_recv = []
  let net_name = []
  for (let i = 0; i < props.value.length; i++) {
    if (props.value[i].bytes_sent != 0 && props.value[i].bytes_sent != 0 &&
        props.value[i].packets_sent != 0 && props.value[i].packets_received != 0) {
      bytes_sent.push(props.value[i].bytes_sent)
      bytes_recv.push(props.value[i].bytes_sent)
      packets_sent.push(props.value[i].packets_sent)
      packets_recv.push(props.value[i].packets_received)
      net_name.push(props.value[i].net_name)
    }
  }
  return [bytes_sent, bytes_recv, packets_sent, packets_recv, net_name]
}
</script>

<style scoped>

/*总布局*/
.CpuStatus_container {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  width: 1050px;
  margin: auto;
}


/*左右布局*/
.Left_container .Right_container {
  height: 100%;
}

/*Net使用率Echart样式*/
#Bytes_Chart, #Packets_Chart {
  transition: all 0.3s;
  border-radius: 2px;
  border: 1px solid #0000001f;
  box-shadow: var(--el-box-shadow-light);
  margin-bottom: 20px;
}

#Bytes_Chart:hover, #Packets_Chart:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light #Bytes_Chart:hover, .light #Packets_Chart:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark #Bytes_Chart:hover, .dark #Packets_Chart:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}
</style>