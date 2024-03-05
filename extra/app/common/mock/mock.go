package mock

import (
	"encoding/json"
	"fmt"

	"github.com/kimchhung/gva/extra/internal/ent"
)

var ROUTE_LIST = `[
	{
	   "path": "/dashboard",
	   "component": "#",
	   "redirect": "/dashboard/analysis",
	   "name": "Dashboard",
	   "meta": {
		 "title": "router.dashboard",
		 "icon": "ant-design:dashboard-filled",
		 "alwaysShow": true
	   },
	   "edges": {
		 "children": [
		   {
			 "path": "analysis",
			 "component": "views/Dashboard/Analysis",
			 "name": "Analysis",
			 "meta": {
			   "title": "router.analysis",
			   "noCache": true,
			   "affix": true
			 }
		   }
		 ]
	   }
	},
	{
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
			 "path": "admin",
			 "component": "views/Authorization/Admin/Admin",
			 "name": "User",
			 "meta": {
			   "title": "router.admin"
			 }
		   },
		   {
			 "path": "menu",
			 "component": "views/Authorization/Menu/Menu",
			 "name": "Menu",
			 "meta": {
			   "title": "router.menuManagement"
			 }
		   },
		   {
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

func GetRoutes() []ent.Route {
	var routes []ent.Route
	err := json.Unmarshal([]byte(ROUTE_LIST), &routes)
	if err != nil {
		fmt.Println(err)
	}

	return routes
}
