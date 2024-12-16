import { defineStore } from 'pinia'
import { convertImgUrl } from '@/utils'
import api from '@/api'

export const useUserStore = defineStore ('user', {
    persist: {
        key: 'computer_mall_user',
        paths: ['token'],
    },
    state: () => ({
        userInfo: {
            id: '',
            nickname: '',
            avatar: 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
            email: '',
            intro: '',
        },
        token: null,
    }),
    getters: {
        userId: state => state.userInfo.id ?? '',
        nickname: state => state.userInfo.nickname ?? '',
        avatar: state => state.userInfo.avatar ?? 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
        intro: state => state.userInfo.intro ?? '',
        email: state => state.userInfo.email ?? '',
    },
    actions: {
        setToken(token) {
            this.token = token
        },
        resetLoginState() {
            this.$reset()
        },
        async logout() {
            await api.logout()
            this.$reset()
        },
        async getUserInfo() {
            if (!this.token) {
                return
            }
            try {
                const resp = await api.getUser()
                if (resp.code === 0) {
                    const data = resp.data
                    this.userInfo = {
                        id: data.id,
                        nickname: data.nickname,
                        avatar: data.avatar ? convertImgUrl(data.avatar) : 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
                        intro: data.intro,
                        email: data.email,
                    }
                    return Promise.resolve(resp.data)
                }
                else {
                    return Promise.reject(resp)
                }
            }
            catch (error) {
                return Promise.reject(error)
            }
        },

    },
})