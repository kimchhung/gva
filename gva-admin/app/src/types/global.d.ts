declare let api: API;

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
