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
  updatePartial = this._updatePartial;
  delete = this._delete;
}

export const module: APIModule = {
  name: 'adminRole',
  resource: new AdminRoleAPI(),
};
