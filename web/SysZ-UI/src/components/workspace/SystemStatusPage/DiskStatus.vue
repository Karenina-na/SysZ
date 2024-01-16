<template>
  <div class="DiskStatus_container">
    <div class="Left_container">
      <el-card v-for="(Disk_info, index) in props.value" :key="index" class="box-card">
        <template #header>
          <div class="card-header">
            <span class="Disk_info_title">Disk {{ Disk_info.disk_name }}</span>
          </div>
        </template>
        <div class="Disk_info_name">Disk Size:
          <div class="Disk_info_value">{{ Disk_info.disk_size }} GB</div>
        </div>
        <hr/>
        <div class="Disk_info_name">Disk Used:
          <div class="Disk_info_value">{{ Disk_info.disk_used }} GB</div>
        </div>
        <hr/>
        <div class="Disk_info_name">Disk Available:
          <div class="Disk_info_value">{{ Disk_info.disk_avi }} GB</div>
        </div>
        <hr/>
        <div class="Disk_info_name">Disk File System Type:
          <div class="Disk_info_value">{{ Disk_info.fs_type }}</div>
        </div>
        <hr/>
        <div class="Disk_info_name">Disk Options Partition Type:
          <div class="Disk_info_value">{{ Disk_info.opts }}</div>
        </div>
        <hr/>
      </el-card>
    </div>
    <div class="Right_container">
      <div id="Disk_Usage_Chart" style="width: 500px"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, watch} from "vue";
import * as echarts from "echarts";
import {DarkTheme} from "@/assets/json/Disk_echart/DarkTheme";
import {LightTheme} from "@/assets/json/Disk_echart/LightTheme";
import {useGlobalStore} from "@/stores/GlobalStore";

const props = defineProps(['value'])
let Disk_Usage_Chart: any = null;
const store = useGlobalStore()

//初始化
onMounted(() => {
  if (props.value) {
    let [Disk_usage, Disk_name] = Create_Data()
    Disk_Usage(Disk_usage, Disk_name);
  } else {
    Disk_Usage([0], ['Not Found']);
  }
});

//侦测器监听父组件传参
watch(props, () => {
  Disk_Usage_Chart.dispose();
  let [Disk_usage, Disk_name] = Create_Data()
  Disk_Usage(Disk_usage, Disk_name);
});

//监听Theme变化
watch(store.getTheme, () => {
  if (!props.value) {
    return
  }
  Disk_Usage_Chart.dispose();
  let [Disk_usage, Disk_name] = Create_Data()
  Disk_Usage(Disk_usage, Disk_name);
});

//绘制Disk使用率
const Disk_Usage = function (usage: Array<number>, name: Array<string>) {
  if (store.getTheme() == "light") {
    echarts.registerTheme('LightTheme', LightTheme)
    Disk_Usage_Chart = echarts.init(document.getElementById("Disk_Usage_Chart")!, 'LightTheme', {height: 372 * usage.length - 22});
  } else {
    echarts.registerTheme('DarkTheme', DarkTheme)
    Disk_Usage_Chart = echarts.init(document.getElementById("Disk_Usage_Chart")!, 'DarkTheme', {height: 372 * usage.length - 22});
  }
  let data = []
  //磁盘使用率
  for (let i = 0; i < usage.length; i++) {
    data.push({
      name: 'Disk' + (i + 1),
      type: 'pie',
      radius: 80,
      center: ['50%', 200 + i * 372],
      stillShowZeroSum: false,
      data: [
        {value: usage[i], name: name[i] + ' Disk Usage'},
        {value: 100 - usage[i], name: name[i] + ' Disk Idle'},
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
  //磁盘使用率图标题
  let title = []
  title.push({
    text: 'Disk Usage',
    left: 'center',
    top: 15,
  })
  for (let i = 0; i < usage.length; i++) {
    title.push({
      subtext: 'Disk Usage: ' + name[i],
      left: '50%',
      top: 50 + i * 372,
      textAlign: 'center'
    })
  }
  Disk_Usage_Chart.setOption({
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
    Disk_Usage_Chart.resize();
  };
}

//构造数据
function Create_Data() {
  let Disk_usage = []
  let Disk_name = []
  for (let i = 0; i < props.value.length; i++) {
    Disk_usage.push(props.value[i].disk_usage)
    Disk_name.push(props.value[i].disk_name)
  }
  return [Disk_usage, Disk_name]
}

</script>

<style scoped>

/*总布局*/
.DiskStatus_container {
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
  height: 350px;
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
.Disk_info_title {
  font-size: 20px;
  text-align: center;
  width: 100%;
  font-weight: bold;
}

/*Disk使用率Echart样式*/
#Disk_Usage_Chart {
  transition: all 0.3s;
  border-radius: 2px;
  border: 1px solid #0000001f;
  box-shadow: var(--el-box-shadow-light);
  margin-bottom: 20px
}

#Disk_Usage_Chart:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light #Disk_Usage_Chart:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark #Disk_Usage_Chart:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*Disk信息字段*/
.Disk_info_name {
  font-size: 14px;
  padding: 6px 7px;
  text-align: left;
  line-height: 35px;
  transition: all 0.3s;
}

.dark .Disk_info_name:hover {
  background-color: #edede9;
  color: #303133;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.light .Disk_info_name:hover {
  background-color: #caf0f8;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.Disk_info_value {
  font-weight: normal;
  text-align: center;
  display: inline-block;
  width: 150px;
  float: right;
}
</style>