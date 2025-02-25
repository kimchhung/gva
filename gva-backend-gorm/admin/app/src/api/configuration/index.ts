import type {
  Configuration,
  CreateConfiguration,
  UpdateConfiguration,
  UpdatePartialConfiguration,
} from './types';

import { ResourceAPI } from '../node';

export class ConfigurationAPI extends ResourceAPI<
  Configuration,
  CreateConfiguration,
  UpdateConfiguration,
  UpdatePartialConfiguration
> {
  create = this._create;

  delete = this._delete;
  getMany = this._getMany;
  update = this._update;
  constructor() {
    super('configuration');
  }
}

export const module: APIModule = {
  name: 'configuration',
  resource: new ConfigurationAPI(),
};
