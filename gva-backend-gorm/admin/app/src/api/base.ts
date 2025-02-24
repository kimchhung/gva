import type { Configuration } from './configuration/types';
import type { AdminIndexConfig, UploadedImage } from './types';

import { req, type RequestOption } from '#/utils/axios';

export class BaseAPI {
  getConfig() {
    return req.get<AdminIndexConfig>({ url: `/config` });
  }

  getDocs() {
    return req.get<Configuration[]>({ url: `/config/docs` });
  }

  now() {
    return req.get<string>({ url: `/now` });
  }

  uploadImage(data: FormData, opts?: RequestOption<UploadedImage>) {
    return req.post<{ filename: string; url: string }>(
      {
        url: `/upload/image`,
        data,
        headers: { 'Content-Type': 'multipart/form-data' },
      },
      opts,
    );
  }
}
