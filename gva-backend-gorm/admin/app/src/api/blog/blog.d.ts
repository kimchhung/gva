import { BlogAPI } from '.';

declare global {
  interface API {
    blog: BlogAPI;
  }
}
