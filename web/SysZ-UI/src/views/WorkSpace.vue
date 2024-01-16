<template>
  <div class="workspace-background">
    <el-scrollbar>
      <el-header class="workspace-head">
        <WorkspaceHeadVue></WorkspaceHeadVue>
      </el-header>
      <el-divider height="0px"/>
      <el-container class="workspace-container">
        <el-aside width="200px">
          <WorkspaceAsideVue></WorkspaceAsideVue>
        </el-aside>
        <router-view></router-view>
      </el-container>
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import WorkspaceAsideVue from '@/components/workspace/WorkspaceAside.vue';
import WorkspaceHeadVue from '@/components/workspace/WorkspaceHead.vue';</script>

<style scoped>
/**背景*/
.workspace-background {
  width: 100%;
  height: 100%;
  user-select: none;
}

/**顶部*/
.workspace-head {
  height: 60px;
  width: 100%;
}

/**下部分容器 */
.workspace-container {
  margin-right: 15px;
  width: 100%;
}
</style>