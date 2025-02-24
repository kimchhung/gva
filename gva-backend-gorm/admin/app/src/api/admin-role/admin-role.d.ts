import { AdminRoleAPI } from '.';

declare global {
  interface API {
    adminRole: AdminRoleAPI;
  }
}
