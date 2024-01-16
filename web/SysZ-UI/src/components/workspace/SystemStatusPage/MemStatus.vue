<template>
  <div class="MemStatusContainer">
    <div class="Left_container">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span class="Mem_info_title">Memory Information</span>
          </div>
        </template>
        <div class="Mem_info_name">Memory Total Size:
          <div class="Mem_info_value">{{ props.value.mem_total }} GB</div>
        </div>
        <hr/>
        <div class="Mem_info_name">Memory Used Size:
          <div class="Mem_info_value">{{ props.value.mem_used }} GB</div>
        </div>
        <hr/>
        <div class="Mem_info_name">Memory Available Size:
          <div class="Mem_info_value">{{ props.value.mem_avi }} GB</div>
        </div>
        <hr/>
      </el-card>
    </div>
    <div class="Right_container">
      <div id="Mem_Usage_Chart" style="width: 500px"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {DarkTheme} from "@/assets/json/Mem_echart/DarkTheme";
import {LightTheme} from "@/assets/json/Mem_echart/LightTheme";

import {onMounted, watch} from "vue";
import {useGlobalStore} from "@/stores/GlobalStore";
import * as echarts from "echarts";

const props = defineProps(['value'])
let mem_chart: any = null;
const store = useGlobalStore()

//初始化
onMounted(() => {
  if (props.value) {
    Mem_Chart(props.value.mem_usage);
  } else {
    Mem_Chart(0);
  }
})

//侦测器监听父组件传参
watch(props, (newVal) => {
  mem_chart.dispose();
  Mem_Chart(newVal.value.mem_usage)
});

//监听Theme变化
watch(store.getTheme, () => {
  if (!props.value) {
    return
  }
  mem_chart.dispose();
  Mem_Chart(props.value.mem_usage);
});

//绘制内存使用率
function Mem_Chart(mem_usage: number) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    mem_chart = echarts.init(document.getElementById("Mem_Usage_Chart")!, 'LightTheme', {height: 300});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    mem_chart = echarts.init(document.getElementById("Mem_Usage_Chart")!, 'DarkTheme', {height: 300});
  }
  let mem_avi = 100 - mem_usage;
  mem_usage = 100 - mem_avi;
  const data = [
    {
      name: "Memory Usage",
      type: 'pie',
      radius: 80,
      center: ['50%', '60%'],
      stillShowZeroSum: false,
      data: [
        {value: mem_usage, name: 'Mem Usage'},
        {value: mem_avi, name: 'Mem Idle'},
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
  mem_chart.setOption({
    title: {
      text: 'Memory Usage',
      left: 'center',
      top: 15,
    },
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        return params.seriesName + "<br>" + params.marker + params.name + ":&nbsp;&nbsp;&nbsp;&nbsp;" + "<span style='font-weight: bold'>" + params.value + "</span>" + " %";
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
    mem_chart.resize();
  };
}
</script>

<style scoped>
/*总布局*/
.MemStatusContainer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  width: 1050px;
  margin: auto;
}

/*卡片布局*/
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 20px;
}

.box-card {
  width: 500px;
  height: 300px;
  border-radius: 6px;
  transition: all 0.3s;
}

.box-card:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light .box-card:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark .box-card:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*左右布局*/
.Left_container .Right_container {
  height: 100%;
}


/*名称*/
.Mem_info_title {
  font-size: 20px;
  text-align: center;
  width: 100%;
  font-weight: bold;
}

/*内存使用率Echart样式*/
#Mem_Usage_Chart {
  transition: all 0.3s;
  border-radius: 2px;
  border: 1px solid #0000001f;
  box-shadow: var(--el-box-shadow-light);
}

#Mem_Usage_Chart:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light #Mem_Usage_Chart:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark #Mem_Usage_Chart:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*CPU信息字段*/
.Mem_info_name {
  font-size: 14px;
  padding: 6px 7px;
  text-align: left;
  line-height: 35px;
  transition: all 0.3s;
  margin-top: 8px;
}

.dark .Mem_info_name:hover {
  background-color: #edede9;
  color: #303133;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.light .Mem_info_name:hover {
  background-color: #caf0f8;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.Mem_info_value {
  font-weight: normal;
  text-align: center;
  display: inline-block;
  width: 150px;
  float: right;
}
</style>