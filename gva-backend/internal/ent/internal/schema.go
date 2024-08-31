// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/gva/app/database/schema\",\"Package\":\"github.com/gva/internal/ent\",\"Schemas\":[{\"name\":\"Admin\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\"},{\"name\":\"department\",\"type\":\"Department\",\"field\":\"department_id\",\"ref_name\":\"members\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"is_enable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isEnable\\\"  rql:\\\"filter,sort\\\"\",\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"deleted_at\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"-\\\"\",\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3},\"annotations\":{\"EntSQL\":{\"default\":\"0\"}}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"username\\\" rql:\\\"column=username,filter,sort\\\"\",\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"sensitive\":true},{\"name\":\"whitelist_ips\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"tag\":\"json:\\\"whitelistIps\\\"\",\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"display_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"displayName,omitempty\\\" rql:\\\"filter,sort\\\"\",\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"department_id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"tag\":\"json:\\\"departmentId,omitempty\\\" rql:\\\"filter,sort\\\"\",\"nillable\":true,\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]},{\"unique\":true,\"fields\":[\"username\",\"deleted_at\"]}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"admin\"}}},{\"name\":\"Department\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"parent\",\"type\":\"Department\",\"field\":\"pid\",\"ref\":{\"name\":\"children\",\"type\":\"Department\"},\"unique\":true,\"inverse\":true},{\"name\":\"members\",\"type\":\"Admin\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"-\\\"\",\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"annotations\":{\"EntSQL\":{\"default\":\"0\"}}},{\"name\":\"is_enable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isEnable\\\"  rql:\\\"filter,sort\\\"\",\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3}},{\"name\":\"name_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"nameId\\\" rql:\\\"column=name_id,filter,sort\\\"\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name\\\" rql:\\\"column=name,filter,sort\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"pid\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"tag\":\"json:\\\"pid,omitempty\\\" rql:\\\"filter,sort\\\"\",\"nillable\":true,\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]},{\"unique\":true,\"fields\":[\"name_id\",\"deleted_at\"]}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"EntGQL\":{\"QueryField\":{},\"RelayConnection\":true},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"dpm\"}}},{\"name\":\"Genre\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"mangas\",\"type\":\"Manga\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"is_enable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isEnable\\\"  rql:\\\"filter,sort\\\"\",\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"deleted_at\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"-\\\"\",\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3},\"annotations\":{\"EntSQL\":{\"default\":\"0\"}}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name\\\" rql:\\\"column=name,filter,sort\\\"\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name\\\" rql:\\\"column=name,filter,sort\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]},{\"unique\":true,\"fields\":[\"name_id\",\"deleted_at\"]}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"magr\"}}},{\"name\":\"Manga\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"chapters\",\"type\":\"MangaChapter\"},{\"name\":\"genres\",\"type\":\"Genre\",\"ref_name\":\"mangas\",\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"is_enable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isEnable\\\"  rql:\\\"filter,sort\\\"\",\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"deleted_at\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"-\\\"\",\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3},\"annotations\":{\"EntSQL\":{\"default\":\"0\"}}},{\"name\":\"name_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name_id\\\" rql:\\\"column=name_id,filter,sort\\\"\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name\\\" rql:\\\"column=name,filter,sort\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"desc\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"desc\\\"\",\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"prodiver\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"provider\\\"\",\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"thumbnail_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"thumbnailUrl\\\"\",\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"authors\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"tag\":\"json:\\\"authors\\\"\",\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]},{\"unique\":true,\"fields\":[\"name\",\"name_id\",\"deleted_at\"]}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"mga\"}}},{\"name\":\"MangaChapter\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"manga\",\"type\":\"Manga\",\"field\":\"manga_id\",\"ref_name\":\"chapters\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"manga_id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"tag\":\"json:\\\"mangaId,omitempty\\\" rql:\\\"filter,sort\\\"\",\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"title\\\" rql:\\\"filter,sort\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"img_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"imgUrl\\\"\",\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"number\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"number\\\" rql:\\\"filter,sort\\\"\",\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"provider_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"providerName\\\"\",\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"chapter_updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"chapterUpdatedAt\\\" rql:\\\"filter,sort\\\"\",\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"unique\":true,\"fields\":[\"provider_name\",\"manga_id\",\"number\"]}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"mcr\"}}},{\"name\":\"Permission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\",\"ref_name\":\"permissions\",\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"group\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"group,omitempty\\\"\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name,omitempty\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"scope\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"scope,omitempty\\\"\",\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"type\",\"type\":{\"Type\":6,\"Ident\":\"permission.Type\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"key,omitempty\\\"\",\"enums\":[{\"N\":\"dynamic\",\"V\":\"dynamic\"},{\"N\":\"static\",\"V\":\"static\"}],\"optional\":true,\"default\":true,\"default_value\":\"dynamic\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"order\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"order,omitempty\\\"\",\"optional\":true,\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"perm\"}}},{\"name\":\"Role\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"admins\",\"type\":\"Admin\",\"ref_name\":\"roles\",\"inverse\":true},{\"name\":\"permissions\",\"type\":\"Permission\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"pxid.ID\",\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"PkgName\":\"pxid\",\"Nillable\":false,\"RType\":{\"Name\":\"ID\",\"Ident\":\"pxid.ID\",\"Kind\":24,\"PkgPath\":\"github.com/gva/app/database/schema/pxid\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"Prefix\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"PrefixIndex\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"XID\":{\"In\":[],\"Out\":[{\"Name\":\"ID\",\"Ident\":\"xid.ID\",\"Kind\":17,\"PkgPath\":\"github.com/rs/xid\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"id\"}}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"createdAt,omitempty\\\" rql:\\\"filter,sort\\\"\",\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"updatedAt,omitempty\\\"\",\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"is_enable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isEnable\\\"  rql:\\\"filter,sort\\\"\",\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"deleted_at\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"-\\\"\",\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3},\"annotations\":{\"EntSQL\":{\"default\":\"0\"}}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"name,omitempty\\\"\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"description,omitempty\\\"\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"order\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"order,omitempty\\\"\",\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"is_changeable\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"tag\":\"json:\\\"isChangeable,omitempty\\\"\",\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3}],\"annotations\":{\"Edges\":{\"StructTag\":\"json:\\\"edges\\\" rql:\\\"-\\\"\"},\"Fields\":{\"ID\":null,\"StructTag\":{\"id\":\"json:\\\"id\\\" rql:\\\"filter,sort\\\"\"}},\"PXID\":{\"Prefix\":\"role\"}}}],\"Features\":[\"privacy\",\"intercept\",\"entql\",\"namedges\",\"bidiedges\",\"schema/snapshot\",\"sql/schemaconfig\",\"sql/lock\",\"sql/modifier\",\"sql/execquery\",\"sql/upsert\",\"sql/versioned-migration\",\"namedges\"]}"
