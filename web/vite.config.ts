import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  base: './', // 设置打包目录
  server: {
    open: true,
    host: true,
    port: 3333,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8888', // 代理到 目标路径
        ws: true,
        changeOrigin: true,
        rewrite: path => path.replace('^/api', ''),
      }
    },
  }
})
