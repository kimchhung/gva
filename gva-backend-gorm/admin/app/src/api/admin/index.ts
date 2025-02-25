import type { Node } from '../node/types';
import type {
  Admin,
  AdminCreate,
  AdminSetTOTP,
  AdminSetTOTPResponse,
  AdminUpdate,
} from './types';

import { req, type RequestOption } from '#/utils/axios';

import { ResourceAPI } from '../node';

export class AdminAPI extends ResourceAPI<Admin, AdminCreate, AdminUpdate> {
  // # expose from resources
  create = this._create;

  delete = this._delete;
  get = this._get;
  getMany = this._getMany;
  update = this._update;
  updatePartial = this._updatePartial;
  constructor() {
    super('admin');
  }

  setTOTP({
    id,
    body,
    opt,
  }: {
    body: AdminSetTOTP;
    id: number | string;
    opt?: RequestOption<AdminSetTOTPResponse>;
  }) {
    return req.patch<Node<AdminSetTOTPResponse>>(
      { url: `${this.baseUrl}/${id}/totp`, data: body },
      opt,
    );
  }
}

export const module: APIModule = {
  name: 'admin',
  resource: new AdminAPI(),
};
