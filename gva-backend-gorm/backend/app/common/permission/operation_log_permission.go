package permission

const (
	OperationLogGroup TPermissionGroup = "operation_log"
)

var (
	OperationLogSuper = newKey(OperationLogGroup, ActionSuper)
	OperationLogView  = newKey(OperationLogGroup, ActionView)

	OperationLogSeeder = NewSeeder(OperationLogGroup,
		OperationLogSuper,
		OperationLogView,
	)
)

func init() {
	allSeeders = append(allSeeders, OperationLogSeeder)
}
