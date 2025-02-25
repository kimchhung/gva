import type {
Blog,
CreateBlog,
UpdateBlog,
UpdatePartialBlog,
} from './types';

import { ResourceAPI } from '../node';

export class BlogAPI extends ResourceAPI<
Blog,
CreateBlog,
UpdateBlog,
UpdatePartialBlog
> {
  constructor() {
    super('blog');
  }

  // # expose from resources
  get = this._get;
  getMany = this._getMany;
  create = this._create;
  update = this._update;
  updatePartial = this._updatePartial;
  delete = this._delete;

  // add custom api here
}

export const module: APIModule = {
  name: 'blog',
  resource: new BlogAPI(),
};
