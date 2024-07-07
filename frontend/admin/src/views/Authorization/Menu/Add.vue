<script setup lang="ts">
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref, unref } from 'vue'
import { useRouter } from 'vue-router'
import Write from './components/Write.vue'

import { createRouter } from '@/api/route'
import { MenuRoute } from '@/api/route/types'
import { useApi } from '@/axios'
import { BaseButton } from '@/components/Button'
// import { saveTableApi } from '@/api/table'

// const { emit } = useEventBus()

const { push, currentRoute } = useRouter()
const goBack = () => {
  const p = currentRoute.value.path.split('/')
  push(p.toSpliced(p.length - 1, 1).join('/'))
}

const { t } = useI18n()

const writeRef = ref<ComponentRef<typeof Write>>()
const loading = ref(false)

const save = async () => {
  const write = unref(writeRef)
  const formData = await write?.submit()

  if (formData) {
    const [res] = await useApi(() => createRouter(formData as MenuRoute), {
      loading
    })

    if (res) goBack()
  }
}

// const save = async () => {
//   const write = unref(writeRef)
//   const formData = (await write?.submit()) as MenuRoute

//   if (!formData) {
//     return
//   }

//   const createOrUpdate = () => {
//     const isCreate = !formData?.id
//     return (
//       useApi(() => (isCreate ? createRouter(formData) : updateRouter(formData))),
//       { loading: saveLoading, onFinally: () => (dialogVisible.value = false) }
//     )
//   }

//   return await createOrUpdate()
// }
</script>

<template>
  <ContentDetailWrap :title="t('button.add')">
    <Write ref="writeRef" />

    <template #header>
      <BaseButton @click="goBack">
        <Icon icon="ep:back" />
        <!-- {{ t('common.back') }} -->
      </BaseButton>
      <BaseButton type="primary" :loading="loading" @click="save"
        >{{ t('button.save') }}
      </BaseButton>
    </template>
  </ContentDetailWrap>
</template>
