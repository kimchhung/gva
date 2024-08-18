import { AuthAPI } from '.';

declare global {
  interface API {
    auth: AuthAPI;
  }
}
