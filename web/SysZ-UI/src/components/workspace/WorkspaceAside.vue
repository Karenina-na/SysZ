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

        <!--namespace-->
        <el-menu-item index="2" @click="getNamespaces()">
          <el-icon :size="20">
            <LocationFilled/>
          </el-icon>
          <template #title><span :class="{ 'workspace-first-level-title-close': isCollapse }"
                                 class="workspace-first-level-title">namespace</span></template>
        </el-menu-item>

        <!--servers status-->
        <el-sub-menu index="3">
          <template #title>
            <el-icon :size="20">
              <Monitor/>
            </el-icon>
            <span :class="{ 'workspace-first-level-title-close': isCollapse }"
                  class="workspace-first-level-title">servers status</span>
          </template>

          <!--colony status-->
          <el-sub-menu v-for="(List, colonyName) of Colonies"
                       :key="colonyName" :index="String(colonyName)">
            <template #title>
              <el-icon class="workspace-second-level-icon">
                <ArrowDown :size="16"/>
              </el-icon>
              <span class="workspace-second-level-title">
                                <span class="workspace-second-colonyName">{{ colonyName }}</span>
                            </span>
            </template>

            <!--server status-->
            <el-menu-item v-for="(serverName, s_index) of List"
                          :key="s_index" :index="String(colonyName) + '-' + String(serverName)"
                          @click="choice_server(colonyName, serverName)">
              <el-icon class="workspace-third-icon">
                <ArrowRight :size="14"/>
              </el-icon>
              <span class="workspace-third-servername">{{ serverName }}</span>
            </el-menu-item>
          </el-sub-menu>
        </el-sub-menu>

        <!--black list-->
        <el-menu-item index="4" @click="getBlackList()">
          <el-icon :size="20">
            <DeleteFilled/>
          </el-icon>
          <template #title><span :class="{ 'workspace-first-level-title-close': isCollapse }"
                                 class="workspace-first-level-title">blacklist</span></template>
        </el-menu-item>

        <!--operator status-->
        <el-menu-item index="5" @click="getOperator()">
          <el-icon :size="20">
            <icon-menu/>
          </el-icon>
          <template #title><span :class="{ 'workspace-first-level-title-close': isCollapse }"
                                 class="workspace-first-level-title">operator</span></template>
        </el-menu-item>

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

//监控集群数据变化
watchEffect(() => {
  Colonies.value = store.GetColoniesAndInstancesNameList()
})

const router = useRouter()
let length = document.documentElement.clientHeight - 110
const isCollapse = ref(true)

//获取命名空间
function getNamespaces() {
  router.push({
    path: '/workspace/namespace',
  })
}

//获取黑名单
function getBlackList() {
  router.push({
    path: '/workspace/blacklist',
  })
}

//操作页面
function getOperator() {
  router.push({
    path: '/workspace/operator',
  })
}

//文档页面
function getDocument() {
  router.push({
    path: '/workspace/document',
  })
}

//选择当前命名空间下的服务器
function choice_server(colony: number, name: number) {
  let namespace = store.GetNamespace();
  router.push({
    path: '/workspace/server',
    query: {
      namespace: namespace,
      colony: colony,
      name: name
    }
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