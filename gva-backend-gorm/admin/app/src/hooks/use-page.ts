import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { useTabs } from '@vben/hooks';

export const usePageRoute = () => {
  const route = useRoute();
  const router = useRouter();
  const tabs = useTabs();

  const goToListPage = () => {
    const currentPath = router.currentRoute.value.path;
    const parentPath = currentPath.slice(0, currentPath.lastIndexOf('/'));

    tabs.closeCurrentTab();
    router.push({
      path: parentPath,
    });
  };

  const goToCreatePage = () => {
    const currentPath = router.currentRoute.value.path;
    router.push({
      path: `${currentPath}/create`,
    });
  };

  const goToEditPage = (id: number | string) => {
    const currentPath = router.currentRoute.value.path;
    router.push({
      path: `${currentPath}/${id}`,
    });
  };

  const isCreatedPage = computed(() => {
    return !route.params.id;
  });

  const isEditPage = computed(() => {
    return !!route.params.id;
  });

  return {
    goToListPage,
    goToCreatePage,
    goToEditPage,
    route,
    router,
    isCreatedPage,
    isEditPage,
    resourceId: route.params.id as string | undefined,
  };
};
