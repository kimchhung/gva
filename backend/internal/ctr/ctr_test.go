package ctr_test

import (
	"testing"

	"github.com/gva/internal/ctr"
)

var _ interface {
	ctr.CTR
} = (*UserController)(nil)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("user"),
	)
}

func TestPrintTree(t *testing.T) {
	con := ctr.New()
	con.Add(&UserController{})
}

func TestGroupController(t *testing.T) {
	// Setup mock controllers
	ctrls := []*ctr.Ctr{
		{ID: "1", PID: ""},
		{ID: "2", PID: "1"},
		{ID: "3", PID: "1"},
		{ID: "4", PID: "2"},
		{ID: "5", PID: "4"},
		{ID: "6", PID: "5"},
	}

	// Test Case 1: Valid Inputs
	t.Run("Valid Inputs", func(t *testing.T) {
		nesteds := ctr.GroupController(ctrls...)
		if len(nesteds) != 1 {
			t.Errorf("wrong nesteds, expect %v, got %v", 1, len(nesteds))
		}

		flats, _ := ctr.FlatController(nesteds...)
		if len(flats) != len(ctrls) {
			t.Errorf("wrong len, expect %v, got %v", len(ctrls), len(flats))
		}
	})
}
