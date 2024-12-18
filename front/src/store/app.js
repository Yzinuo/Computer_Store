import { defineStore } from 'pinia'
import api from '@/api'

export const useAppStore = defineStore('app', {
    state: () => ({
        searchFlag: false,
        loginFlag: false,
        registerFlag: false,
        collapsed: false, // 侧边栏折叠（移动端）
    }),

    actions: {
        setCollapsed(flag) { this.collapsed = flag },
        setLoginFlag(flag) { this.loginFlag = flag },
        setRegisterFlag(flag) { this.registerFlag = flag },
        setSearchFlag(flag) { this.searchFlag = flag },


        async getPageList() {
            const resp = await api.getPageList()
            if (resp.data.code === 0) {
                this.page_list = resp.data.data
                this.page_list?.forEach(e => (e.cover = convertImgUrl(e.cover)))
            }
        },
    },
})



