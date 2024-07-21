<script setup lang="ts">
import { MenuRoute } from '@/api/menu/types'
import { BaseButton } from '@/components/Button'
import { ContentDetailWrap } from '@/components/ContentDetailWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { ref, unref } from 'vue'
import { useRouter } from 'vue-router'
import Write from './components/Write.vue'

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
    const [res] = await api.department.create({
      body: formData as MenuRoute,
      opt: { loading }
    })

    if (res) goBack()
  }
}
</script>

<template>
  <ContentDetailWrap :title="t('button.add')">
    <Write ref="writeRef" />

    <template #header>
      <BaseButton @click="goBack">
        <Icon icon="ep:back" />
      </BaseButton>
      <BaseButton type="primary" :loading="loading" @click="save"
        >{{ t('button.save') }}
      </BaseButton>
    </template>
  </ContentDetailWrap>
</template>
