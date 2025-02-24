import type { AdminRole } from '../admin-role/types';

export interface Admin {
  createdAt: Date;
  id: number;
  username: string;
  name: string;
  nameId: string;
  roles: AdminRole[];
  ipWhiteList: null | string[];
  status: number;
}

export interface AdminCreate {
  username: string;
  password: string;
  name: string;
  roles: { id: number }[];
}

export interface AdminUpdate {
  username: string;
  name: string;
  roles: { id: number }[];
}

export interface AdminUpdatePatch {
  password?: string;
  roles?: { id: number }[];
  ipWhiteList?: string[];
  status?: number;
}

export interface AdminSetTOTP {
  totp: string;
}

export interface AdminSetTOTPResponse {
  totpKey: string;
  totpURL: string;
}
