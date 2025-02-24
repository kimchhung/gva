export interface Configuration {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  key: string;
  value?: any;
  description: string;
  metadata?: { [key: string]: any };
  type: string;
  children: Configuration[];
  parentId?: number;
  parent?: Configuration;
  rootId?: number;
  root?: Configuration;
  parents?: Configuration[];
  allChildren?: Configuration[];
}

export interface CreateConfiguration {
  [key: string]: any;
}

export interface UpdateConfiguration {
  [key: string]: any;
}

export interface UpdatePartialConfiguration {
  [key: string]: any;
}
