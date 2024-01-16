<template>
  <div class="CpuStatus_container">
    <div class="Left_container">
      <el-card v-for="(cpu_info, index) in props.value" :key="index" class="box-card">
        <template #header>
          <div class="card-header">
            <span class="CPU_info_title">{{ cpu_info.cpu_name }}</span>
          </div>
        </template>
        <div class="CPU_info_name">CPU Cores Num:
          <div class="CPU_info_value">{{ cpu_info.cpu_core_num }}</div>
        </div>
        <hr/>
        <div class="CPU_info_name">CPU Frequency:
          <div class="CPU_info_value">{{ cpu_info.cpu_frequency }} GHz</div>
        </div>
        <hr/>
        <div class="CPU_info_name">CPU Vendor ID:
          <div class="CPU_info_value">{{ cpu_info.cpu_vendor_id }}</div>
        </div>
        <hr/>
        <div class="CPU_info_name">CPU Physical ID:
          <div class="CPU_info_value">{{ cpu_info.cpu_physical_id }}</div>
        </div>
        <hr/>
      </el-card>
    </div>
    <div class="Right_container">
      <div id="CPU_Usage_Chart" style="width: 500px"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, watch} from "vue";
import * as echarts from "echarts";
import {DarkTheme} from "@/assets/json/CPU_echart/DarkTheme";
import {LightTheme} from "@/assets/json/CPU_echart/LightTheme";
import {useGlobalStore} from "@/stores/GlobalStore";

const props = defineProps(['value'])
let CPU_Usage_Chart: any = null;
const store = useGlobalStore()

//初始化
onMounted(() => {
  if (props.value) {
    let [CPU_usage, CPU_name] = Create_Data()
    CPU_Usage(CPU_usage, CPU_name);
  } else {
    CPU_Usage([0], ['Not Found']);
  }
});

//侦测器监听父组件传参
watch(props, (newVal) => {
  if (!newVal.value) {
    return
  }
  CPU_Usage_Chart.dispose();
  let [CPU_usage, CPU_name] = Create_Data()
  CPU_Usage(CPU_usage, CPU_name);
});

//监听Theme变化
watch(store.getTheme, () => {
  if (!props.value) {
    return
  }
  CPU_Usage_Chart.dispose();
  let [CPU_usage, CPU_name] = Create_Data()
  CPU_Usage(CPU_usage, CPU_name);
});

//绘制CPU使用率
const CPU_Usage = function (usage: Array<number>, name: Array<string>) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    CPU_Usage_Chart = echarts.init(document.getElementById("CPU_Usage_Chart")!, 'LightTheme', {height: 322 * usage.length - 22});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    CPU_Usage_Chart = echarts.init(document.getElementById("CPU_Usage_Chart")!, 'DarkTheme', {height: 322 * usage.length - 22});
  }
  //CPU使用率
  let data = []
  for (let i = 0; i < usage.length; i++) {
    data.push({
      name: 'CPU' + (i + 1),
      type: 'pie',
      radius: 80,
      center: ['50%', 200 + i * 322],
      stillShowZeroSum: false,
      data: [
        {value: usage[i], name: 'CPU Usage'},
        {value: 100 - usage[i], name: 'CPU Idle'},
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
  //CPU名称
  let title = []
  title.push({
    text: 'CPU Usage',
    left: 'center',
    top: 15,
  })
  for (let i = 0; i < usage.length; i++) {
    title.push({
      subtext: 'CPU Usage: ' + name[i],
      left: '50%',
      top: 50 + 322 * i,
      textAlign: 'center'
    })
  }

  CPU_Usage_Chart.setOption({
    title: title,
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
    CPU_Usage_Chart.resize();
  };
}

//构造数据
function Create_Data() {
  let CPU_usage = []
  let CPU_name = []
  for (let i = 0; i < props.value.length; i++) {
    CPU_usage.push(props.value[i].cpu_usage)
    CPU_name.push(props.value[i].cpu_name)
  }
  return [CPU_usage, CPU_name]
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
  margin-bottom: 20px;
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
.CPU_info_title {
  font-size: 20px;
  text-align: center;
  width: 100%;
  font-weight: bold;
}

/*CPU使用率Echart样式*/
#CPU_Usage_Chart {
  transition: all 0.3s;
  border-radius: 2px;
  border: 1px solid #0000001f;
  box-shadow: var(--el-box-shadow-light);
  margin-bottom: 20px;
}

#CPU_Usage_Chart:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light #CPU_Usage_Chart:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark #CPU_Usage_Chart:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*CPU信息字段*/
.CPU_info_name {
  font-size: 14px;
  padding: 6px 7px;
  text-align: left;
  line-height: 35px;
  transition: all 0.3s;
}

.dark .CPU_info_name:hover {
  background-color: #edede9;
  color: #303133;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.light .CPU_info_name:hover {
  background-color: #caf0f8;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.CPU_info_value {
  font-weight: normal;
  text-align: center;
  display: inline-block;
  width: 150px;
  float: right;
}
</style>