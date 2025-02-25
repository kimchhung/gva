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
  get = this._get;

  // extends
  getMany = this._getMany;
  constructor() {
    super('operation-log');
  }
}

export const module: APIModule = {
  name: 'operationLog',
  resource: new OperationLogAPI(),
};
