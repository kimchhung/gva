import { getServerTime } from '@/api'
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
    isInit: false
  }),
  actions: {
    now() {
      return dayjs().add(this.timeDiff, 'ms')
    },

    async syncTimeDiff() {
      const [data] = await useApi(() => getServerTime(), {
        loading: this.$state,
        onError: (error) => console.error('Failed to sync time:', error)
      })

      if (data) {
        const now = dayjs()
        this.timeDiff = dayjs(data).diff(now, 'ms')
      }
    },

    async startSyncTimeDiff() {
      await this.syncTimeDiff()
      return setInterval(this.syncTimeDiff, 50000)
    },

    startTimer() {
      setInterval(() => {
        this.currentTime = this.now()
      }, 1000)
    }
  }
})

export const useServerTime = () => {
  const timeStore = useTimeStore(store)

  onBeforeMount(async () => {
    if (timeStore.isInit) return
    await timeStore.startSyncTimeDiff()
    timeStore.startTimer()
    timeStore.isInit = true
  })

  return timeStore
}
