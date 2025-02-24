export interface Role {
  id: number;
  name: string;
}

export interface BaseAdmin {
  [key: string]: any;
  createdAt: Date;
  id: string;
  username: string;
  roles: Role[];
  permissionScope: string[];
  roleNameId: string[];
  isSuperAdmin: boolean;
}

export interface AuthLoginBody {
  password: string;
  username: string;
  totp: string;
}
export interface AuthLoginReq {
  body: AuthLoginBody;
}

export interface AuthLoginResp {
  admin: BaseAdmin;
  token: string;
}
