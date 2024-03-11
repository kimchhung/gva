package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/api/admin/module/permission/dto"
	"github.com/kimchhung/gva/extra/api/admin/module/permission/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*PermissionController)(nil)

type PermissionController struct {
	service *service.PermissionService
}

func (con *PermissionController) Init(r fiber.Router) {
	permission := r.Group("permission")
	rctrl.Register(permission, con)
}

func NewPermissionController(service *service.PermissionService) *PermissionController {
	return &PermissionController{
		service: service,
	}
}

// @Tags Permission
// @Summary List all Permissions
// @Description Get a list of all Permissions
// @ID list-all-Permissions
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=map[string]dto.PermissionResponse{list=[]dto.PermissionResponse}} "Successfully retrieved Permissions"
// @Router /permission [get]
// @Security Bearer
func (con *PermissionController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetPermissions(c.UserContext())
		if err != nil {
			return err
		}

		return request.Response(c,
			response.Data(list),
			response.Message("Permission list retreived successfully!"),
		)
	})
}

// @Tags Permission
// @Security Bearer
// @Summary Get a Permission
// @Description Get a Permission by ID
// @ID get-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Success   200 {object} response.Response{data=dto.PermissionResponse}
// @Router /permission/{id} [get]
func (con *PermissionController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Permission").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.GetPermissionByID(c.UserContext(), param.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The permission retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Permission
// @Security Bearer
// @Summary Create a Permission
// @Description Create a new Permission with the provided details
// @ID create-Permission
// @Accept  json
// @Produce  json
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} response.Response{data=dto.PermissionResponse} "Successfully created Permission"
// @Router /permission [post]
func (con *PermissionController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Permission").DoWithScope(func() []fiber.Handler {
		body := new(dto.PermissionRequest)

		return []fiber.Handler{
			request.Validate(
				request.BodyParser(body),
			),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreatePermission(c.UserContext(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The permission retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Permission
// @Security Bearer
// @Summary Update a Permission
// @Description Update a Permission by ID
// @ID update-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} response.Response{data=dto.PermissionResponse} "Successfully updated Permission"
// @Router /permission/{id} [patch]
func (con *PermissionController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Permission").DoWithScope(func() []fiber.Handler {
		body := new(dto.PermissionRequest)
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
				request.BodyParser(body),
			),
			func(c *fiber.Ctx) error {
				data, err := con.service.UpdatePermission(c.UserContext(), param.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
					response.Message("The permission retrieved successfully!"),
				)
			},
		}
	})
}

// @Tags Permission
// @Security Bearer
// @Summary Delete a Permission
// @Description Delete a Permission by ID
// @ID delete-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Success  200 {object} response.Response{} "Successfully deleted Permission"
// @Router /permission/{id} [delete]
func (con *PermissionController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Permission").DoWithScope(func() []fiber.Handler {
		param := &struct {
			ID int `params:"id" validate:"gt=0"`
		}{}

		return []fiber.Handler{
			request.Validate(
				request.ParamsParser(param),
			),
			func(c *fiber.Ctx) error {
				if err := con.service.DeletePermission(c.UserContext(), param.ID); err != nil {
					return err
				}

				return request.Response(c,
					response.Message("The permission retrieved successfully!"),
				)
			},
		}
	})
}
