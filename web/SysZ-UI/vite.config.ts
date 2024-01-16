import { fileURLToPath, URL } from 'node:url'
import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd(), 'VITE_')
  return {
    //加载环境变量  process.env.VITE_*
    define: {
      'process.env': {
        VITE_SYSZ_UI_TITLE: env.VITE_SYSZ_UI_TITLE,
        VITE_SYSZ_URL: JSON.stringify(env.VITE_SYSZ_URL),
        VITE_SERVER_PORT: JSON.stringify(env.VITE_SERVER_PORT),
      }
    },
    plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server: {
      port: Number(env.VITE_SERVER_PORT),
      proxy: {
        '/api': {
          target: env.VITE_SYSZ_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      }
    },
  }
})
