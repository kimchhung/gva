package permission

const (
	DocumentGroup TPermissionGroup = "document"
)

var (
	DocumentSuper = newKey(DocumentGroup, ActionSuper)
	DocumentView  = newKey(DocumentGroup, ActionView)

	DocumentSeeder = NewSeeder(DocumentGroup,
		DocumentView,
	)
)

func init() {
	allSeeders = append(allSeeders, DocumentSeeder)
}
