import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import unocss from 'unocss/vite'
import path from 'path'

export default defineConfig({
    plugins: [vue(),
              unocss()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'), // 配置 @ 指向 src 目录
        },
    },
})