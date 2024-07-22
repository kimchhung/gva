<script setup lang="ts">
import { MenuRoute } from '@/api/menu/types'
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Detail from './components/Detail.vue'

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

  currentRow.value = res.data
}

getTableDetail()
</script>

<template>
  <ContentDetailWrap :title="t('button.detail')">
    <template #header>
      <BaseButton @click="goBack">
        <Icon icon="ep:back" />
      </BaseButton>
    </template>
    <Detail :current-row="currentRow" />
  </ContentDetailWrap>
</template>
