import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server:{
    host: '0.0.0.0',  // 允许所有IP访问
    port: 5173,        // 可以修改为你需要的端口
    strictPort: true,  // 如果端口已经被占用，Vite会退出而不是自动更改端口
    proxy:{
      '/api':{
        target:'http://localhost:3456',
        changeOrigin:true,
        secure:false,
        rewrite:(path)=>{
          return path.replace(/^\/api/,'')
        }
      }
    }
  }
})
