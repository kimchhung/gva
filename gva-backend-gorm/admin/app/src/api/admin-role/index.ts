import type { AdminRole, CreateAdminRole, UpdateAdminRole } from './types';

import { ResourceAPI } from '../node';

export class AdminRoleAPI extends ResourceAPI<
  AdminRole,
  CreateAdminRole,
  UpdateAdminRole
> {
  constructor() {
    super('adminrole');
  }

  // #expose from resources
  get = this._get;
  getMany = this._getMany;
  create = this._create;
  update = this._update;
}

export const module: APIModule = {
  name: 'adminRole',
  resource: new AdminRoleAPI(),
};
