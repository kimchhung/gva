import { api } from '@/api'
import { useApi } from '@/axios'
import dayjs from 'dayjs'
import { defineStore } from 'pinia'
import { onBeforeMount } from 'vue'
import { store } from '..'

export const useTimeStore = defineStore('time', {
  state: () => ({
    timeDiff: 0,
    loading: false,
    currentTime: null as dayjs.Dayjs | null,
    timer: null as NodeJS.Timeout | null
  }),
  actions: {
    async sync() {
      const [data] = await useApi(() => api().getServerTime(), {
        loading: this.$state,
        onError: (error) => console.error('Failed to sync time:', error)
      })

      if (data) {
        const now = dayjs()
        this.timeDiff = dayjs(data).diff(now)
      }
    },
    now() {
      return dayjs().add(this.timeDiff)
    },

    async startTimer() {
      if (this.timer) {
        return
      }

      await this.sync()

      this.timer = setInterval(() => {
        this.currentTime = this.now()
      }, 1000)

      this.timer = setInterval(() => {
        this.sync()
      }, 15000)
    }
  }
})

export const useServerTime = () => {
  const timeStore = useTimeStore(store)

  onBeforeMount(() => {
    timeStore.startTimer()
  })

  return timeStore
}
