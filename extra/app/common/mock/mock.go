package mock

import (
	"encoding/json"
	"fmt"

	"github.com/kimchhung/gva/extra/internal/ent"
)

var ROUTE_LIST = `[
	{
	   "id":1,
	   "path": "/dashboard",
	   "component": "#",
	   "redirect": "/dashboard/",
	   "name": "Dashboard",
	   "meta": {
		 "title": "router.dashboard",
		 "icon": "ant-design:dashboard-filled",
		 "alwaysShow": true
	   },
	   "edges": {
		 "children": [
		   {
			 "id":2,
			 "parentId":1,
			 "path": "/",
			 "component": "views/Dashboard/Welcome",
			 "name": "Welcome",
			 "meta": {
			   "title": "router.welcome",
			   "noCache": true,
			   "affix": true
			 }
		   }
		 ]
	   }
	},
	{
	   "id":3,
	   "path": "/authorization",
	   "component": "#",
	   "redirect": "/authorization/admin",
	   "name": "Authorization",
	   "meta": {
		 "title": "router.authorization",
		 "icon": "eos-icons:role-binding",
		 "alwaysShow": true
	   },
	   "edges": {
		 "children": [
		   {
			"id":4,
			"parentId":3,
			 "path": "admin",
			 "component": "views/Authorization/Admin/Admin",
			 "name": "User",
			 "meta": {
			   "title": "router.admin"
			 }
		   },
		   {
			"id":5,
			"parentId":3,
			 "path": "menu",
			 "component": "views/Authorization/Menu/Menu",
			 "name": "Menu",
			 "meta": {
			   "title": "router.menuManagement"
			 }
		   },
		   {
			"id":6,
			"parentId":3,
			 "path": "role",
			 "component": "views/Authorization/Role/Role",
			 "name": "Role",
			 "meta": {
			   "title": "router.role"
			 }
		   }
		 ]
	   }
	}
]
`

func GetRoutes() []*ent.Route {
	var routes []*ent.Route
	err := json.Unmarshal([]byte(ROUTE_LIST), &routes)
	if err != nil {
		fmt.Println(err)
	}

	return routes
}
