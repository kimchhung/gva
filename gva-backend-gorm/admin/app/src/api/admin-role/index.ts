import type { AdminRole, CreateAdminRole, UpdateAdminRole } from './types';

import { ResourceAPI } from '../node';

export class AdminRoleAPI extends ResourceAPI<
  AdminRole,
  CreateAdminRole,
  UpdateAdminRole
> {
  create = this._create;

  delete = this._delete;
  // #expose from resources
  get = this._get;
  getMany = this._getMany;
  update = this._update;
  updatePartial = this._updatePartial;
  constructor() {
    super('adminrole');
  }
}

export const module: APIModule = {
  name: 'adminRole',
  resource: new AdminRoleAPI(),
};
