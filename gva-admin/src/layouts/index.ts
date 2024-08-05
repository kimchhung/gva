const BasicLayout = () => import('./basic.vue');

const IFrameView = () => import('@gva/layouts').then((m) => m.IFrameView);

const AuthPageLayout = () => import('@gva/layouts').then((m) => m.AuthPageLayout);

export { AuthPageLayout, BasicLayout, IFrameView };
