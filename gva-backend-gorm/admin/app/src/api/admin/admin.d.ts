import { AdminAPI } from '.';

declare global {
  interface API {
    admin: AdminAPI;
  }
}
