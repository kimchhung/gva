import type {
  CreateNode,
  GetManyMeta,
  GetManyNode,
  GetNode,
  Node,
  UpdateNode,
} from './types';

import { req } from '#/utils/axios';

export class ResourceAPI<T extends object = Node<any>> {
  base: string;
  name: string;

  constructor(name: string) {
    this.name = name;
    this.base = `/${this.name.replace('/', '')}`;
  }

  create({ body }: CreateNode<T>) {
    return req.post<Node<T>>({ data: body, url: this.base });
  }

  delete({ id }: GetNode) {
    return req.delete<Node<T>>({ url: `${this.base}/${id}` });
  }

  get({ id }: GetNode) {
    return req.get<Node<T>>({ url: `${this.base}/${id}` });
  }

  getMany({ query }: GetManyNode<T>) {
    // to do use parse usePagi
    return req.get<Node<T>[], GetManyMeta>({ params: query, url: this.base });
  }

  update({ body, id }: UpdateNode<T>) {
    return req.put<Node<T>>({ data: body, url: `${this.base}/${id}` });
  }
}
