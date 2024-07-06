package pm

// permission manager

import (
	"fmt"
)

type Event int

const (
	EventRole Event = iota
	EventResource
	EventRule
)

// Store interface defines methods for data persistence
type Store interface {
	// all roles
	LoadRoles() ([]*Role, error)

	// all resources
	LoadResources() ([]*Resource, error)

	// rules
	LoadRules() ([]*Rule, error)
}

type Resource interface {
	Attributes() map[string]any
}

type User interface {
	Roles() []string
	Attributes() map[string]any
}

type Role map[string][]string

type Rule struct {
	UserAttrKey        string
	ResourceAttrKey    string
	Action             string
	ComparisonOperator string
	Value              interface{}
}

func (r *Rule) Evaluate(user User, resource Resource, action string) (bool, error) {
	userAttrs := user.Attributes()
	resourceAttrs := resource.Attributes()

	switch r.ComparisonOperator {
	case "==":
		return r.Value == userAttrs[r.UserAttrKey] && r.Value == resourceAttrs[r.ResourceAttrKey], nil
	case "<>":
		return r.Value != userAttrs[r.UserAttrKey] && r.Value != resourceAttrs[r.ResourceAttrKey], nil
	default:
		return false, fmt.Errorf("unsupported comparison operator: %s", r.ComparisonOperator)
	}
}

// AccessControlManager struct now includes a Store
type AccessControlManager struct {
	Name      string
	store     *Store
	roles     map[string][]string
	resources map[string]Resource
	rules     []*Rule
}

func (acm *AccessControlManager) option(opt ...Option) {
	for _, opt := range opt {
		opt(acm)
	}
}

type Option func(*AccessControlManager)

func NewAccessControlManager(name string, opt ...Option) *AccessControlManager {
	acm := &AccessControlManager{
		Name:      name,
		roles:     map[string][]string{},
		resources: map[string]Resource{},
		rules:     []*Rule{},
	}

	acm.option(opt...)
	return acm
}
