<script setup lang="ts">
import { MenuRoute } from '@/api/menu/types'
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref, unref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Write from './components/Write.vue'

// const { emit } = useEventBus()

const { push, currentRoute } = useRouter()
const goBack = () => {
  const p = currentRoute.value.path.split('/')
  push(p.toSpliced(p.length - 1, 1).join('/'))
}

const { t } = useI18n()
const { query } = useRoute()
const currentRow = ref<MenuRoute>()

const getTableDetail = async () => {
  if (!query.id) return
  const [res, err] = await api.menu.get({ id: String(query.id) })
  if (err) return
  if (res) {
    currentRow.value = res.data
  }
}

getTableDetail()

const writeRef = ref<ComponentRef<typeof Write>>()
const loading = ref(false)

const save = async () => {
  const write = unref(writeRef)
  const formData = await write?.submit()

  if (!formData) return

  const { id, ...body } = formData as MenuRoute
  const [data] = await api.menu.update({
    id,
    body
  })

  if (data) goBack()
}
</script>

<template>
  <ContentDetailWrap :title="t('button.detail')">
    <Write ref="writeRef" :current-row="currentRow" />

    <template #header>
      <BaseButton @click="goBack">
        <Icon icon="ep:back" />
      </BaseButton>
      <BaseButton type="primary" :loading="loading" @click="save">
        {{ t('button.save') }}
      </BaseButton>
    </template>
  </ContentDetailWrap>
</template>
