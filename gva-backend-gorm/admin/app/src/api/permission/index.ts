import type { Permission } from './types';

import { ResourceAPI } from '../node';

export class PermissionAPI extends ResourceAPI<Permission> {
  constructor() {
    super('permission');
  }

  // extends
  getMany = this._getMany;
}

export const module: APIModule = {
  name: 'permission',
  resource: new PermissionAPI(),
};
