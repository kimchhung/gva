package ctr_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gva/internal/ctr"
)

type H func(context.Context) error
type CTR = ctr.CTR

var _ interface {
	ctr.CTR
	ctr.Children
} = (*UserController)(nil)

type UserController struct {
}

func (c *UserController) Init(ctm *ctr.Ctr) *ctr.Ctr {
	ctm.SetGroup("user")
	return ctm
}

func (c *UserController) Children() []ctr.CTR {
	return []ctr.CTR{
		&UserRoleController{},
		&AdminControler{},
	}
}

func (c *UserController) GetUser(r *ctr.Route) *ctr.Route {
	r.Get("/getUser").Do(func(ctx context.Context) error {

		return nil
	})
	return r
}

func (c *UserController) CreateUser(r *ctr.Route) *ctr.Route {
	r.Get("/createtUser").Do(func(ctx context.Context) error {

		return nil
	})
	return r
}

type AdminControler struct {
}

func (c *AdminControler) Init(ctm *ctr.Ctr) *ctr.Ctr {
	ctm.SetGroup("admin")
	return ctm
}

func (c *AdminControler) GetUser(r *ctr.Route) *ctr.Route {
	r.Get("/create").Do(func(ctx context.Context) error {

		return nil
	})
	return r
}

func (c *AdminControler) CreateUser(r *ctr.Route) *ctr.Route {
	return r.Get("/createtUser").Do(func(ctx context.Context) error {

		return nil
	})
}

var _ interface {
	ctr.CTR
} = (*UserRoleController)(nil)

type UserRoleController struct {
}

func (c *UserRoleController) Children() []ctr.CTR {
	return []ctr.CTR{
		&AdminControler{},
	}
}

func (c *UserRoleController) Init(ctm *ctr.Ctr) *ctr.Ctr {
	ctm.SetGroup("role")
	return ctm
}

func (c *UserRoleController) GetRole(r *ctr.Route) *ctr.Route {
	r.Get("/GetRole").Do(func(ctx context.Context) error {

		return nil
	})
	return r
}

func (c *UserRoleController) CreateRole(r *ctr.Route) *ctr.Route {
	r.Get("/CreateRole").Do(func(ctx context.Context) error {

		return nil
	})
	return r
}

func TestPrintTree(t *testing.T) {

	con := ctr.New()
	con.Add(&UserController{})
	con.Add(&AdminControler{})
	con.PrintTree("")

	i := 0
	con.ForEach(func(c *ctr.Ctr) {
		i++
		fmt.Println(i)
	})
}
