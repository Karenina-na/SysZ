<template>
  <el-scrollbar :height="length">
    <div class="workspace-aside-background">
      <el-radio-group v-model="isCollapse" class="workspace-expand-and-collapse" style="margin-bottom: 20px">
        <el-radio-button :label="false" class="workspace-expand-and-collapse-button" size="large">
          <span>expand</span>
        </el-radio-button>
        <el-radio-button :label="true" class="workspace-expand-and-collapse-button" size="large">
          <span>collapse</span>
        </el-radio-button>
      </el-radio-group>
      <el-menu :collapse="isCollapse" class="workspace-el-menu-vertical" default-active="-1">

        <!--system status-->
        <el-sub-menu index="1">
          <template #title>
            <el-icon :size="20">
              <House/>
            </el-icon>
            <span :class="{ 'workspace-first-level-title-close': isCollapse }"
                  class="workspace-first-level-title">system status</span>
          </template>
          <el-menu-item v-for="(system, index) of systemStatus" :key="index" :index="'1-' + String(system)"
                        @click="choice_system(system)">
            <el-icon class="workspace-second-icon">
              <DArrowRight :size="16"/>
            </el-icon>
            <span class="workspace-third-system">{{ system }}</span>
          </el-menu-item>
        </el-sub-menu>

        <!--document status-->
        <el-menu-item index="6" @click="getDocument()">
          <el-icon :size="20">
            <Document/>
          </el-icon>
          <template #title><span :class="{ 'workspace-first-level-title-close': isCollapse }"
                                 class="workspace-first-level-title">document</span></template>
        </el-menu-item>
      </el-menu>
    </div>
  </el-scrollbar>
</template>

<script lang="ts" setup>
import {onMounted, ref, watchEffect} from 'vue'
import {
  ArrowDown,
  ArrowRight,
  DArrowRight,
  DeleteFilled,
  Document,
  House,
  LocationFilled,
  Menu as IconMenu,
  Monitor
} from '@element-plus/icons-vue'
import {useRouter} from 'vue-router'
import {SetupServersStore} from '@/stores/SetupServersStore'

let store = SetupServersStore()
const Colonies = ref()

const router = useRouter()
let length = document.documentElement.clientHeight - 110
const isCollapse = ref(true)

//文档页面
function getDocument() {
  router.push({
    path: '/workspace/document',
  })
}

//选择查看的系统环境
function choice_system(status: string) {
  if (status === "SystemStatus") {
    router.push({
      path: '/workspace/system',
    })
    return
  }
  router.push({
    path: '/workspace/system',
    query: {
      status: status
    }
  })
}

const systemStatus = ['SystemStatus', 'HostStatus', 'CpuStatus', 'MemStatus', 'DiskStatus', 'NetworkStatus', 'PoolStatus']

//加载初始数据
onMounted(() => {
})

</script>
<style scoped>
/**背景 */
.workspace-aside-background {
  width: 100%;
  height: 100%;
}

/** 下拉栏顶部按钮*/
.workspace-expand-and-collapse {
  width: 100%;
  display: flex;
  justify-content: center;
}

.workspace-expand-and-collapse-button span {
  font-size: 14px;
  font-weight: 800;
}

/**垂直下拉栏 */
.workspace-el-menu-vertical:not(.el-menu--collapse) {
  width: 100%;
}

/**下拉栏字体样式 */
.workspace-first-level-title {
  font-size: 15px;
  margin-left: 5px;
  font-weight: 500;
}

/**下拉栏收缩字体样式 */
.workspace-first-level-title-close {
  font-size: 15px;
  font-weight: 400;
}

.workspace-second-level-icon {
  margin-left: 10px;
}

.workspace-second-colonyName {
  font-size: 13px;
  font-weight: 500;
  margin-left: 5px;
}

.workspace-third-icon {
  margin-left: 10px;
}

.workspace-third-servername,
.workspace-third-system {
  font-size: 13px;
  margin-left: 10px;
}
</style>