<script lang="ts" setup>
import type { AdminIndexConfig } from '#/api/types';

import { onMounted, ref } from 'vue';

import { AuthenticationLogin } from '@vben/common-ui';
import { Copy } from '@vben/icons';
import { $t } from '@vben/locales';

import { Button } from 'ant-design-vue';

import { useAuthStore } from '#/store';
import { copyToClipboard } from '#/utils/helper/copy';

defineOptions({ name: 'Login' });

const authStore = useAuthStore();

const adminConfig = ref<AdminIndexConfig | null>(null);

onMounted(async () => {
  const [res] = await api.getConfig();

  if (res?.code === 0) {
    adminConfig.value = res.data;
  }
});

const handleCopy = () => {
  if (adminConfig?.value?.publicIp) {
    copyToClipboard(adminConfig?.value?.publicIp);
  }
};
</script>

<template>
  <AuthenticationLogin
    :loading="authStore.loginLoading"
    @submit="authStore.authLogin"
  >
    <template v-if="adminConfig" #bottom>
      <div class="text-muted-foreground flex items-center justify-center">
        {{ $t('auth.currentIP') }}: {{ adminConfig?.publicIp }}
        <Button class="ml-2" @click="handleCopy">
          <Copy class="size-4" />
        </Button>
      </div>
    </template>
  </AuthenticationLogin>
</template>
