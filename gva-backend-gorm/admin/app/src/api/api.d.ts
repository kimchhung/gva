import { BaseAPI } from './base';

declare global {
  // eslint-disable-next-line @typescript-eslint/no-empty-object-type
  interface API extends BaseAPI {}
}
