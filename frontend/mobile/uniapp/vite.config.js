import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    uni()
  ],
  // 解决路径别名问题
  resolve: {
    alias: {
      '@': './'
    }
  },
  // 开发服务器配置
  server: {
    port: 10086,
    open: true,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false
      }
    }
  }
})