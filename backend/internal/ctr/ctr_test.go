package ctr_test

import (
	"fmt"
	"testing"

	"github.com/gva/internal/ctr"
	"github.com/gva/internal/treeprint"
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
		ctr.ID(c),
		ctr.Group("user"),
	)
}

func TestAddChild(t *testing.T) {
	// Create a root controller
	root := ctr.New(
		ctr.ID("root"),
	)

	// Attempt to add the child to the root
	childs, err := root.AddChildren(
		ctr.New(
			ctr.ParentID("root"),
			ctr.ID("child"),
		),
		ctr.New(
			ctr.ParentID("root"),
			ctr.ID("child2"),
		),
	)

	child := childs[0]
	childSon, err := child.AddChild(ctr.New(
		ctr.ID("childSon"),
		ctr.ParentID("child"),
	))

	if err != nil {
		t.Fatalf("Failed to add childSon: %v", err)
	}

	childUselessSon, err := child.AddChild(ctr.New(
		ctr.ID("childUselessSon"),
		ctr.ParentID("child"),
	))

	if err != nil {
		t.Fatalf("Failed to add childUselessSon: %v", err)
	}

	// Verify the child's properties
	if child.PID != root.ID {
		t.Errorf("Expected child PID to be '%s', got '%s'", root.ID, child.PID)
	}
	if child.Parent() != root {
		t.Errorf("Expected child parent to be '%p', got '%p'", root, child.Parent())
	}
	if child.Root() != root {
		t.Errorf("Expected child root to be '%p', got '%p'", root, child.Root())
	}
	if len(root.Children) != 1 || root.Children[0] != child {
		t.Errorf("Expected root to have one child, got %d", len(root.Children))
	}

	if childUselessSon.Root() != root || childSon.Root() != root {
		t.Errorf("Expected same root , got %v | %v, expected %v", childUselessSon.Root(), childSon.Root(), root)
	}

	if err != nil {
		t.Errorf("failed to add child %v", err)
	}

	fmt.Println("")
	treeprint.Print(ctr.CreateIDTreePrint(nil, root))
	fmt.Println("")
	fmt.Println("")
	treeprint.Print(ctr.CreateIDTreePrint(nil, child))
	fmt.Println("")
}

func TestGroupController(t *testing.T) {
	// Setup mock controllers
	ctrls := []*ctr.Ctr{
		ctr.New(
			ctr.ParentID(""),
			ctr.ID("/system"),
		),
		ctr.New(
			ctr.ParentID("/system"),
			ctr.ID("/menu"),
			ctr.AddRoute(ctr.NewRoute().Get("/get")),
			ctr.AddRoute(ctr.NewRoute().Post("/create")),
			ctr.AddRoute(ctr.NewRoute().Put("/update")),
			ctr.AddRoute(ctr.NewRoute().Delete("/delete")),
			ctr.AddRoute(ctr.NewRoute().Patch("/update")),
		),
		ctr.New(
			ctr.ParentID("/system"),
			ctr.ID("/admin"),
			ctr.AddRoute(ctr.NewRoute().Get("/get")),
			ctr.AddRoute(ctr.NewRoute().Post("/create")),
			ctr.AddRoute(ctr.NewRoute().Put("/update")),
			ctr.AddRoute(ctr.NewRoute().Delete("/delete")),
			ctr.AddRoute(ctr.NewRoute().Patch("/update")),
		),
		ctr.New(
			ctr.ParentID("/system"),
			ctr.ID("/user"),
			ctr.AddRoute(ctr.NewRoute().Get("/get")),
			ctr.AddRoute(ctr.NewRoute().Post("/create")),
			ctr.AddRoute(ctr.NewRoute().Put("/update")),
			ctr.AddRoute(ctr.NewRoute().Delete("/delete")),
			ctr.AddRoute(ctr.NewRoute().Patch("/update")),
		),
		ctr.New(
			ctr.ParentID("/user"),
			ctr.ID("/{id}"),
			ctr.AddRoute(ctr.NewRoute().Get("/get")),
			ctr.AddRoute(ctr.NewRoute().Post("/create")),
			ctr.AddRoute(ctr.NewRoute().Put("/update")),
			ctr.AddRoute(ctr.NewRoute().Delete("/delete")),
			ctr.AddRoute(ctr.NewRoute().Patch("/update")),
		),
		ctr.New(
			ctr.ID("/book"),
			ctr.Group("/book"),
			ctr.AddRoute(ctr.NewRoute().Get("/get")),
			ctr.AddRoute(ctr.NewRoute().Post("/create")),
			ctr.AddRoute(ctr.NewRoute().Put("/update")),
			ctr.AddRoute(ctr.NewRoute().Delete("/delete")),
			ctr.AddRoute(ctr.NewRoute().Patch("/update")),
		),
	}

	nesteds := ctr.GroupController(ctrls...)
	root := ctr.New(ctr.ID("root"))
	root.AddChildren(nesteds...)
	treeprint.Print(ctr.CreateIDTreePrint(nil, root))
	// Test Case 1: Valid Inputs
	t.Run("Valid Inputs", func(t *testing.T) {
		nesteds := ctr.GroupController(ctrls...)

		if len(nesteds) != 2 {
			t.Errorf("wrong nesteds, expect %v, got %v", 2, len(nesteds))
		}

		flats, _ := ctr.FlatController(root)
		if len(flats) != len(ctrls) {
			t.Errorf("wrong len, expect %v, got %v", len(ctrls), len(flats))
		}
	})
}

func AddToParentNested(parent *ctr.Ctr, child *ctr.Ctr) (*ctr.Ctr, error) {
	childNode, err := parent.AddChild(child)
	if err != nil {
		return nil, err // Return the error immediately if adding a child fails
	}

	for _, node := range childNode.Children {
		childNode.AddChild(node)
	}

	return childNode, nil
}
