<script setup lang="ts">
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Detail from './components/Detail.vue'

const { push, currentRoute } = useRouter()
const goBack = () => {
  const p = currentRoute.value.path.split('/')
  push(p.toSpliced(p.length - 1, 1).join('/'))
}

const { t } = useI18n()

const currentRow = ref<Nullable<Recordable>>(null)

const getTableDet = async () => {
  // const res = await getTableDetApi(query.id as string)
  // if (res) {
  //   currentRow.value = res.data
  // }
}

getTableDet()
</script>

<template>
  <ContentDetailWrap :title="t('button.detail')">
    <template #header>
      <BaseButton @click="goBack">
        {{ t('common.back') }}
      </BaseButton>
    </template>
    <Detail :current-row="currentRow" />
  </ContentDetailWrap>
</template>
