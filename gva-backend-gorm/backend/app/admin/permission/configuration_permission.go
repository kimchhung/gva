package permission

const (
	ConfigurationGroup TPermissionGroup = "configuration"
)

var (
	ConfigurationSuper  = newKey(ConfigurationGroup, ActionSuper)
	ConfigurationView   = newKey(ConfigurationGroup, ActionView)
	ConfigurationAdd    = newKey(ConfigurationGroup, ActionAdd)
	ConfigurationEdit   = newKey(ConfigurationGroup, ActionEdit)
	ConfigurationDelete = newKey(ConfigurationGroup, ActionDelete)

	ConfigurationSeeder = NewSeeder(ConfigurationGroup,
		ConfigurationSuper,
		ConfigurationView,
		ConfigurationAdd,
		ConfigurationEdit,
		ConfigurationDelete,
	)
)

func init() {
	allSeeders = append(allSeeders, ConfigurationSeeder)
}
