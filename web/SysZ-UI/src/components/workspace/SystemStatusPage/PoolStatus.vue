<template>
  <div class="PoolStatusContainer">
    <div class="Left_container">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span class="Pool_info_title">Coroutine-Pool Information</span>
          </div>
        </template>
        <div class="Pool_info_name">Core Threads:
          <div class="Pool_info_value"> {{ props.value.core_num }}</div>
        </div>
        <hr/>
        <div class="Pool_info_name">Max Threads:
          <div class="Pool_info_value">{{ props.value.max_num }}</div>
        </div>
        <hr/>
        <div class="Pool_info_name">Tasks Number:
          <div class="Pool_info_value">{{ props.value.job_num }}</div>
        </div>
        <hr/>
      </el-card>
    </div>
    <div class="Right_container">
      <div id="Pool_Usage_Chart" style="width: 500px"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {DarkTheme} from "@/assets/json/Pool_echart/DarkTheme";
import {LightTheme} from "@/assets/json/Pool_echart/LightTheme";

import {onMounted, watch} from "vue";
import {useGlobalStore} from "@/stores/GlobalStore";
import * as echarts from "echarts";

const props = defineProps(['value'])
let pool_chart: any = null;
const store = useGlobalStore()

//初始化
onMounted(() => {
  if (props.value) {
    Pool_Chart(props.value.activate_num, props.value.max_num);
  } else {
    Pool_Chart(0, 0);
  }
})

//侦测器监听父组件传参
watch(props, (newVal) => {
  pool_chart.dispose();
  Pool_Chart(newVal.value.activate_num, newVal.value.max_num);
});

//监听Theme变化
watch(store.getTheme, () => {
  if (!props.value) {
    return
  }
  pool_chart.dispose();
  Pool_Chart(props.value.activate_num, props.value.max_num);
});

//绘制内存使用率
function Pool_Chart(pool_active: number, pool_max: number) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    pool_chart = echarts.init(document.getElementById("Pool_Usage_Chart")!, 'LightTheme', {height: 300});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    pool_chart = echarts.init(document.getElementById("Pool_Usage_Chart")!, 'DarkTheme', {height: 300});
  }
  const data = [
    {
      name: "Active Coroutine",
      type: 'pie',
      radius: 80,
      center: ['50%', '60%'],
      stillShowZeroSum: false,
      data: [
        {value: pool_active, name: 'Activate Coroutines'},
        {value: pool_max - pool_active, name: 'Free Coroutines'},
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
  pool_chart.setOption({
    title: {
      text: 'Coroutine Pool Activate Number',
      left: 'center',
      top: 15,
    },
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        return params.seriesName + "<br>" + params.marker + params.name + ":&nbsp;&nbsp;&nbsp;&nbsp;" + "<span style='font-weight: bold'>" + params.value + "</span>";
      },
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      top: '50',
    },
    series: data,
  });
  window.onresize = function () {//自适应大小
    pool_chart.resize();
  };
}
</script>

<style scoped>
/*总布局*/
.PoolStatusContainer {
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
.Pool_info_title {
  font-size: 20px;
  text-align: center;
  width: 100%;
  font-weight: bold;
}

/*内存使用率Echart样式*/
#Pool_Usage_Chart {
  transition: all 0.3s;
  border-radius: 2px;
  border: 1px solid #0000001f;
  box-shadow: var(--el-box-shadow-light);
}

#Pool_Usage_Chart:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light #Pool_Usage_Chart:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark #Pool_Usage_Chart:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*CPU信息字段*/
.Pool_info_name {
  font-size: 14px;
  padding: 6px 7px;
  text-align: left;
  line-height: 35px;
  transition: all 0.3s;
  margin-top: 8px;
}

.dark .Pool_info_name:hover {
  background-color: #edede9;
  color: #303133;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.light .Pool_info_name:hover {
  background-color: #caf0f8;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.Pool_info_value {
  font-weight: normal;
  text-align: center;
  display: inline-block;
  width: 150px;
  float: right;
}
</style>