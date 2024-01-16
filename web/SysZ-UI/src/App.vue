<script lang="ts" setup>
import {RouterView} from 'vue-router'</script>

<template>
  <el-scrollbar>
    <RouterView/>
  </el-scrollbar>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

html,
body,
#app {
  height: 100%;
  width: 100%;
}
</style>
