/* ------------------------------
  ! Generated files, do not edit
  ! Use pnpm run pull:permission instead
  ------------------------------ */

export enum ADMIN_PERMISSION {
  ADD = 'admin:add',
  DELETE = 'admin:delete',
  EDIT = 'admin:edit',
  SUPER = 'admin:super',
  VIEW = 'admin:view',
}

export enum ADMIN_ROLE_PERMISSION {
  ADD = 'admin_role:add',
  DELETE = 'admin_role:delete',
  EDIT = 'admin_role:edit',
  SUPER = 'admin_role:super',
  VIEW = 'admin_role:view',
}

export enum CONFIGURATION_PERMISSION {
  ADD = 'configuration:add',
  DELETE = 'configuration:delete',
  EDIT = 'configuration:edit',
  SUPER = 'configuration:super',
  VIEW = 'configuration:view',
}

export enum DOCUMENT_PERMISSION {
  SUPER = 'document:super',
  VIEW = 'document:view',
}

export enum OPERATION_LOG_PERMISSION {
  SUPER = 'operation_log:super',
  VIEW = 'operation_log:view',
}

export type PERMISSION =
  | ADMIN_PERMISSION
  | ADMIN_ROLE_PERMISSION
  | CONFIGURATION_PERMISSION
  | DOCUMENT_PERMISSION
  | OPERATION_LOG_PERMISSION;
