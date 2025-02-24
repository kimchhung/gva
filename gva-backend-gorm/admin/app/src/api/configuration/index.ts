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
  constructor() {
    super('configuration');
  }

  create = this._create;
  update = this._update;
  getMany = this._getMany;
  delete = this._delete;
}

export const module: APIModule = {
  name: 'configuration',
  resource: new ConfigurationAPI(),
};
