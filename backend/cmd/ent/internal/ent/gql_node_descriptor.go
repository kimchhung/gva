// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"

	"github.com/gva/app/database/schema/pulid"
	"github.com/gva/internal/ent/admin"
	"github.com/gva/internal/ent/permission"
	"github.com/gva/internal/ent/role"
	"github.com/gva/internal/ent/route"
)

// Node in the graph.
type Node struct {
	ID     pulid.ID `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string     `json:"type,omitempty"` // edge type.
	Name string     `json:"name,omitempty"` // edge name.
	IDs  []pulid.ID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

// Node implements Noder interface
func (a *Admin) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Admin",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.IsEnable); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "bool",
		Name:  "is_enable",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Username); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "username",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Password); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "password",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.WhitelistIps); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "[]string",
		Name:  "whitelist_ips",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DisplayName); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "display_name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Role",
		Name: "roles",
	}
	err = a.QueryRoles().
		Select(role.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (pe *Permission) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pe.ID,
		Type:   "Permission",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(pe.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Group); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "group",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Key); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "key",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Order); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "order",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Role",
		Name: "roles",
	}
	err = pe.QueryRoles().
		Select(role.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (r *Role) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Role",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(r.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.IsEnable); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "bool",
		Name:  "is_enable",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Description); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Order); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "order",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.IsChangeable); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "bool",
		Name:  "is_changeable",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Admin",
		Name: "admins",
	}
	err = r.QueryAdmins().
		Select(admin.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Permission",
		Name: "permissions",
	}
	err = r.QueryPermissions().
		Select(permission.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Route",
		Name: "routes",
	}
	err = r.QueryRoutes().
		Select(route.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (r *Route) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Route",
		Fields: make([]*Field, 12),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(r.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.IsEnable); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "bool",
		Name:  "is_enable",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.ParentID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "pulid.ID",
		Name:  "parent_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Path); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "path",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Component); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "component",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Redirect); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "redirect",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Order); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "int",
		Name:  "order",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Type); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "route.Type",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Meta); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "types.RouteMeta",
		Name:  "meta",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Route",
		Name: "parent",
	}
	err = r.QueryParent().
		Select(route.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Route",
		Name: "children",
	}
	err = r.QueryChildren().
		Select(route.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Role",
		Name: "roles",
	}
	err = r.QueryRoles().
		Select(role.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node returns the node with given global ID.
//
// This API helpful in case you want to build
// an administrator tool to browser all types in system.
func (c *Client) Node(ctx context.Context, id pulid.ID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}
