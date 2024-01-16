<template>
  <div class="ThemeButton-box-themeButton">
    <el-switch v-model="theme" :active-icon="MoonIcon" :inactive-icon="SunnyIcon" active-color="#635c5c"
               active-value="dark"
               inactive-color="#dcdfe6" inactive-value="light" inline-prompt @click="changeTheme"/>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue'
import SunnyIcon from '@/components/icon/SunnyIcon.vue'
import MoonIcon from '@/components/icon/MoonIcon.vue'
import {useDark, useToggle} from '@vueuse/core'
import {useGlobalStore} from '@/stores/GlobalStore'

const store = useGlobalStore()

const isDark = useDark({
  selector: 'html',
  attribute: 'class',
  storageKey: 'Theme',
  valueDark: 'dark',
  valueLight: 'light',
})
const toggleDark = useToggle(isDark)
let theme = ref("light")

onMounted(() => {
  theme.value = isDark.value ? "dark" : "light"
  store.setTheme(theme.value)
})

let changeTheme = () => {
  if (theme.value == "light") {
    toggleDark(false);
    store.setTheme("light");
  } else {
    toggleDark(true);
    store.setTheme("dark");
  }
}
</script>

<style scoped>
.ThemeButton-box-themeButton {
  align-items: center;
  justify-content: center;
}
</style>