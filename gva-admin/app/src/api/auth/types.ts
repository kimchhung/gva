export interface BaseAdmin {
  [key: string]: any;
  createdAt: Date;
  id: string;
  username: string;
}

export interface AuthLoginBody {
  password: string;
  username: string;
}
export interface AuthLoginReq {
  body: AuthLoginBody;
}

export interface AuthLoginResp {
  admin: BaseAdmin;
  token: string;
}
