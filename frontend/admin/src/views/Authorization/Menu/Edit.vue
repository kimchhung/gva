<script setup lang="ts">
import { updateRouter } from '@/api/route'
import { MenuRoute } from '@/api/route/types'
import { useApi } from '@/axios'
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref, unref } from 'vue'
import { useRouter } from 'vue-router'
import Write from './components/Write.vue'

// const { emit } = useEventBus()

const { push, currentRoute } = useRouter()
const goBack = () => {
  const p = currentRoute.value.path.split('/')
  push(p.toSpliced(p.length - 1, 1).join('/'))
}

const { t } = useI18n()

const currentRow = ref<MenuRoute>()

const getTableDetail = async () => {
  // const res = await getTableDetApi(query.id as string)
  // if (res) {
  //   currentRow.value = res.data
  // }
}

getTableDetail()

const writeRef = ref<ComponentRef<typeof Write>>()
const loading = ref(false)

const save = async () => {
  const write = unref(writeRef)
  const formData = await write?.submit()

  if (formData) {
    const [res] = await useApi(() => updateRouter(formData as MenuRoute), {
      loading
    })

    if (res) goBack()
  }
}
</script>

<template>
  <ContentDetailWrap :title="t('button.detail')">
    <Write ref="writeRef" :current-row="currentRow" />

    <template #header>
      <BaseButton @click="goBack">
        {{ t('common.back') }}
      </BaseButton>
      <BaseButton type="primary" :loading="loading" @click="save">
        {{ t('button.save') }}
      </BaseButton>
    </template>
  </ContentDetailWrap>
</template>
