import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import unocss from 'unocss/vite'
import path from 'path'

export default defineConfig(({ mode }) => {
    // 加载环境变量
    const env = loadEnv(mode, process.cwd(), '')

    return {
        plugins: [vue(), unocss()],
        resolve: {
            alias: {
                '@': path.resolve(__dirname, 'src'), // 配置 @ 指向 src 目录
            },
        },
        server: {
            host: '0.0.0.0',
            port: 5173,
            open: false,
            proxy: {
                '/api': {
                    target: env.VITE_BACKEND_URL, // 使用加载的环境变量
                    changeOrigin: true,
                },
            },
        },
    }
})