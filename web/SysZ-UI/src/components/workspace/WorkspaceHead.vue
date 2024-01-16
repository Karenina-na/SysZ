<template>
  <div class="workspace-head-background">
    <span class="workspace-head-title" @click="fresh_data()">Themis</span>
    <div class="workspace-namespace-card">
      <div class="workspace-namespace-icon">
        <ProudIcon/>
      </div>
      <NamespaceCard :namespace="namespaces"/>
      <div class="workspace-namespace-icon">
        <TurtleIcon/>
      </div>
    </div>
    <div class="workspace-theme">
      <ThemeButton class="workspace-theme-button"></ThemeButton>
    </div>
    <div class="workspace-icons">
      <User class="workspace-icon" width="30px"/>
      <Setting class="workspace-icon" width="30px"/>
      <InfoFilled class="workspace-icon" width="30px"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {InfoFilled, Setting, User} from '@element-plus/icons-vue'
import ThemeButton from '@/components/ThemeButton.vue'
import NamespaceCard from './NamespacePage/NamespaceCard.vue';
import {SetupServersStore} from '@/stores/SetupServersStore'
import TurtleIcon from '@/components/TurtleIcon.vue'
import ProudIcon from '@/components/ProudIcon.vue'
import {ref, watchEffect} from 'vue'

let store = SetupServersStore()
const namespaces = ref()

//监控命名空间变化
watchEffect(() => {
  namespaces.value = store.GetNamespace()
})

let timer: any = null

//刷新网页
function fresh_data() {
  if (timer) {
    clearTimeout(timer)
  }

  //设置定时器防止刷新过快
  timer = setTimeout(() => {
    window.location.reload()
  }, 1000)
}


</script>

<style scoped>
/** 背景 */
.workspace-head-background {
  height: 100%;
  width: 100%;
  display: flex;
}

/**主题按钮 大小放大*/
.workspace-theme {
  margin-top: 12px;
  height: 100%;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.workspace-theme-button {
  scale: 1.2;
  margin-right: 30px;
}


/**标题 */
.workspace-head-title {
  font-family: Quotes, -apple-system, BlinkMacSystemFont,
  "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
  font-weight: 800;
  font-size: 55px;
  cursor: pointer;
  transition: all 0.3s;
}

.workspace-head-title:hover {
  transform: scale(1.1);
  transition: all 0.5s;
  color: #409eff;
}

.workspace-head-title:active {
  transform: scale(1.0);
  transition: all 0.5s;
  color: #409eff;
}

/**图标 */
.workspace-icons {
  margin-top: 25px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.workspace-icon {
  margin-right: 30px;
  transition: all 0.3s;
}

.workspace-icon:hover {
  cursor: pointer;
  transform: scale(1.3);
  transition: all 0.5s;
}

.workspace-icon:active {
  transform: scale(1.1);
  transition: all 0.1s;
}


/**当前命名空间 */
.workspace-namespace-card {
  margin-left: 22px;
  margin-top: 10px;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

/**命名空间加载球 */
.workspace-namespace-icon {
  scale: 0.7;
}
</style>