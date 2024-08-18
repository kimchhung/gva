import type { AuthLoginReq, AuthLoginResp, BaseAdmin } from './types';

import { req } from '#/utils/axios';

export class AuthAPI {
  base: string;
  name: string;

  constructor(name: string) {
    this.name = name;
    this.base = `/${this.name.replace('/', '')}`;
  }

  login({ body }: AuthLoginReq) {
    return req.post<AuthLoginResp>({ data: body, url: `${this.base}/login` });
  }

  me() {
    return req.get<BaseAdmin>({ url: `${this.base}/me` });
  }
}

export const module: APIModule = {
  name: 'auth',
  resource: new AuthAPI('auth'),
};
