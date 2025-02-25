import type { Permission } from './types';

import { ResourceAPI } from '../node';

export class PermissionAPI extends ResourceAPI<Permission> {
  // extends
  getMany = this._getMany;

  constructor() {
    super('permission');
  }
}

export const module: APIModule = {
  name: 'permission',
  resource: new PermissionAPI(),
};
