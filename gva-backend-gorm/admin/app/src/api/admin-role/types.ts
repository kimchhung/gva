import type { Permission } from '../permission/types';

export interface AdminRole {
  id: number;
  name: string;
  description: string;
  status: number;
  permissions: Permission[];
  createdAt: Date;
  [key: string]: any;
}

export interface CreateAdminRole {
  name: string;
  description: string;
  permissions: string[];
}

export interface UpdateAdminRole {
  name: string;
  description: string;
  permissions: string[];
}
