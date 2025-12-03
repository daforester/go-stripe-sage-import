import {defineConfig} from 'vite'
import path from 'path'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@font": path.resolve(__dirname, "./src/assets/fonts"),
      "@image": path.resolve(__dirname, "./src/assets/images"),
      "@view": path.resolve(__dirname, "./src/views"),
    },
  },
  plugins: [
    vue(),
  ],
})
