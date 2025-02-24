schema "gva_backend" {
  charset = "utf8mb4"
  collate = "utf8mb4_general_ci"
}

table "operation_logs" {
  schema = schema.gva_backend
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = true
    type = datetime(3)
  }
  column "admin_id" {
    null     = true
    type     = bigint
    unsigned = true
  }
  column "role_ids" {
    null = true
    type = json
  }
  column "method" {
    null = true
    type = varchar(7)
  }
  column "path" {
    null = true
    type = varchar(255)
  }
  column "scope" {
    null = true
    type = varchar(255)
  }
  column "ip" {
    null = true
    type = varchar(39)
  }
  column "data" {
    null = true
    type = json
  }
  column "code" {
    null = true
    type = int
  }
  column "error" {
    null = true
    type = longtext
  }
  column "msg" {
    null = true
    type = longtext
  }
  column "latency" {
    null     = true
    type     = int
    unsigned = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_operation_logs_admin" {
    columns     = [column.admin_id]
    ref_columns = [table.admins.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "fk_operation_logs_admin" {
    columns = [column.admin_id]
  }
}

table "permissions" {
  schema = schema.gva_backend
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = false
    type = datetime(3)
  }
  column "group" {
    null = false
    type = varchar(255)
  }
  column "scope" {
    null = false
    type = varchar(255)
  }
  column "order" {
    null    = false
    type    = int
    default = 0
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_uniq_permissions_scope" {
    unique  = true
    columns = [column.scope]
  }
}

table "admin_admin_roles" {
  schema = schema.gva_backend
  column "admin_id" {
    null     = false
    type     = bigint
    unsigned = true
  }
  column "admin_role_id" {
    null     = false
    type     = bigint
    unsigned = true
  }
  primary_key {
    columns = [column.admin_id, column.admin_role_id]
  }
  foreign_key "admin_admin_roles_admin" {
    columns     = [column.admin_id]
    ref_columns = [table.admins.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "admin_admin_roles_admin_role" {
    columns     = [column.admin_role_id]
    ref_columns = [table.admin_roles.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "admin_admin_roles_role" {
    columns = [column.admin_role_id]
  }
}

table "admin_role_permissions" {
  schema = schema.gva_backend
  column "admin_role_id" {
    null     = false
    type     = bigint
    unsigned = true
  }
  column "permission_id" {
    null     = false
    type     = bigint
    unsigned = true
  }
  primary_key {
    columns = [column.admin_role_id, column.permission_id]
  }
  foreign_key "fk_admin_role_permissions_admin_role" {
    columns     = [column.admin_role_id]
    ref_columns = [table.admin_roles.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fk_admin_role_permissions_permission" {
    columns     = [column.permission_id]
    ref_columns = [table.permissions.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "fk_admin_role_permissions_permission" {
    columns = [column.permission_id]
  }
}

table "admin_roles" {
  schema = schema.gva_backend
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = true
    type = datetime(3)
  }
  column "updated_at" {
    null = true
    type = datetime(3)
  }
  column "deleted_at" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
  }
  column "name" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null    = true
    type    = bigint
    default = 1
  }
  column "order" {
    null    = false
    type    = int
    default = 0
  }
  column "type" {
    null    = false
    type    = int
    default = 1
  }
  column "name_id" {
    null = false
    type = varchar(30)
  }
  column "description" {
    null = true
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_admin_roles_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_admin_roles_name_id_deleted_at" {
    unique  = true
    columns = [column.name_id, column.deleted_at]
  }
}

table "admins" {
  schema = schema.gva_backend
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "created_at" {
    null = true
    type = datetime(3)
  }
  column "updated_at" {
    null = true
    type = datetime(3)
  }
  column "deleted_at" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
  }
  column "password_hash" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null    = true
    type    = bigint
    default = 1
  }
  column "name" {
    null = true
    type = varchar(255)
  }
  column "username" {
    null = true
    type = varchar(255)
  }
  column "ip_white_list" {
    null = true
    type = json
  }
  column "current_login_ip" {
    null = true
    type = varchar(255)
  }
  column "last_login_ip" {
    null = true
    type = varchar(255)
  }
  column "current_region" {
    null = true
    type = varchar(255)
  }
  column "last_region" {
    null = true
    type = varchar(255)
  }
  column "current_login_at" {
    null = true
    type = datetime(3)
  }
  column "last_login_at" {
    null = true
    type = datetime(3)
  }
  column "google_secret_key" {
    null = true
    type = varchar(255)
  }
  column "google_otp" {
    null = true
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_admins_deleted_at" {
    columns = [column.deleted_at]
  }
  index "idx_admins_username_deleted_at" {
    unique  = true
    columns = [column.username, column.deleted_at]
  }
}

table "configurations" {
  schema = schema.gva_backend
  column "id" {
    null           = false
    type           = bigint
    unsigned       = true
    auto_increment = true
  }
  column "key" {
    null = false
    type = varchar(255)
  }
  column "type" {
    null = false
    type = varchar(255)
  }
  column "value" {
    null = true
    type = json
  }
  column "parent_id" {
    null     = true
    type     = bigint
    unsigned = true
  }
  column "status" {
    null    = false
    type    = tinyint
    default = 1
  }
  column "updated_at" {
    null = true
    type = datetime(3)
  }
  column "created_at" {
    null = false
    type = datetime(3)
  }
  column "deleted_at" {
    null     = false
    type     = bigint
    default  = 0
    unsigned = true
  }
  column "description" {
    null = true
    type = varchar(255)
  }
  column "metadata" {
    null = true
    type = json
  }
  column "root_id" {
    null     = true
    type     = bigint
    unsigned = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_configuration_parent_id" {
    columns     = [column.parent_id]
    ref_columns = [table.configurations.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fk_configuration_root_id" {
    columns     = [column.root_id]
    ref_columns = [table.configurations.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "fk_configuration_parent_id" {
    columns = [column.parent_id]
  }
  index "fk_configuration_root_id" {
    columns = [column.root_id]
  }
  index "key" {
    unique  = true
    columns = [column.key, column.deleted_at]
  }
}


