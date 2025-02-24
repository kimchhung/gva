import type {
  CreateOperationLog,
  OperationLog,
  UpdateOperationLog,
  UpdatePartialOperationLog,
} from './types';

import { ResourceAPI } from '../node';

export class OperationLogAPI extends ResourceAPI<
  OperationLog,
  CreateOperationLog,
  UpdateOperationLog,
  UpdatePartialOperationLog
> {
  constructor() {
    super('operation-log');
  }

  // extends
  getMany = this._getMany;
  get = this._get;
}

export const module: APIModule = {
  name: 'operationLog',
  resource: new OperationLogAPI(),
};
