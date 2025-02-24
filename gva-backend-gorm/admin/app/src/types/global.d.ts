// eslint-disable-next-line @typescript-eslint/no-empty-object-type
declare interface API {}

declare interface APIModule {
  name: string;
  resource: any;
}

declare global {
  declare interface Window {
    api: API;
  }

  declare interface APIModule {
    name: string;
    resource: any;
  }
}

declare module '@wangeditor/editor-for-vue';
declare module 'chinese-lunar-calendar';
