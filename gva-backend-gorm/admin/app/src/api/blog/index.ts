import type { Blog, CreateBlog, UpdateBlog, UpdatePartialBlog } from './types';

import { ResourceAPI } from '../node';

export class BlogAPI extends ResourceAPI<
  Blog,
  CreateBlog,
  UpdateBlog,
  UpdatePartialBlog
> {
  create = this._create;

  delete = this._delete;
  // # expose from resources
  get = this._get;
  getMany = this._getMany;
  update = this._update;
  updatePartial = this._updatePartial;
  constructor() {
    super('blog');
  }

  // add custom api here
}

export const module: APIModule = {
  name: 'blog',
  resource: new BlogAPI(),
};
