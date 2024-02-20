package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/module/permission/dto"
	"github.com/kimchhung/gva/app/module/permission/service"
	"github.com/kimchhung/gva/internal/rctrl"
	"github.com/kimchhung/gva/utils/request"
)

var _ interface {
	rctrl.FiberRouter
	IPermissionController
} = (*PermissionController)(nil)

type IPermissionController interface {
	Create(meta *rctrl.RouteMeta) rctrl.MetaHandler
	List(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Get(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Update(meta *rctrl.RouteMeta) rctrl.MetaHandler
	Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler
}

type PermissionController struct {
	service *service.PermissionService
}

func (con *PermissionController) Routes(r fiber.Router) {
	r.Route("permission",
		func(router fiber.Router) {
			rctrl.Register(router, con)
		},
	)
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
// @Success  200 {object} request.Response{data=[]dto.PermissionResponse} "Successfully retrieved Permissions"
// @Router /permission [get]
func (con *PermissionController) List(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("get many Permissions").Do(func(c *fiber.Ctx) error {
		list, err := con.service.GetPermissions(c.UserContext())
		if err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "Permission list retreived successfully!",
			Data:    list,
		})
	})
}

// @Tags Permission
// @Summary Get a Permission
// @Description Get a Permission by ID
// @ID get-Permission-by-id
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Permission ID"
// @Success   200 {object} request.Response{data=dto.PermissionResponse}
// @Router /permission/{id} [get]
func (con *PermissionController) Get(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/:id").Name("get one Permission").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		data, err := con.service.GetPermissionByID(c.UserContext(), id)
		if err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "The permission retrieved successfully!",
			Data:    data,
		})
	})
}

// @Tags Permission
// @Summary Create a Permission
// @Description Create a new Permission with the provided details
// @ID create-Permission
// @Accept  json
// @Produce  json
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} request.Response{data=dto.PermissionResponse} "Successfully created Permission"
// @Router /permission [post]
func (con *PermissionController) Create(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Post("/").Name("create one Permission").DoWithScope(func() []fiber.Handler {
		req := new(dto.PermissionRequest)

		return []fiber.Handler{
			request.Validate(request.BodyParser(&req)),

			func(c *fiber.Ctx) error {
				data, err := con.service.CreatePermission(c.UserContext(), *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The permission was created successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Permission
// @Summary Update a Permission
// @Description Update a Permission by ID
// @ID update-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Param Permission body dto.PermissionRequest true "Permission data"
// @Success  200 {object} request.Response{data=dto.PermissionResponse} "Successfully updated Permission"
// @Router /permission/{id} [patch]
func (con *PermissionController) Update(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Patch("/:id").Name("update one Permission").DoWithScope(func() []fiber.Handler {
		req := new(dto.PermissionRequest)

		return []fiber.Handler{
			request.Validate(request.BodyParser(&req)),
			func(c *fiber.Ctx) error {
				id, err := strconv.Atoi(c.Params("id"))
				if err != nil {
					return err
				}

				data, err := con.service.UpdatePermission(c.UserContext(), id, *req)
				if err != nil {
					return err
				}

				return request.Resp(c, request.Response{
					Message: "The permission was updated successfully!",
					Data:    data,
				})
			},
		}
	})
}

// @Tags Permission
// @Summary Delete a Permission
// @Description Delete a Permission by ID
// @ID delete-Permission-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Permission ID"
// @Success  200 {object} request.Response{} "Successfully deleted Permission"
// @Router /permission/{id} [delete]
func (con *PermissionController) Delete(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Delete("/:id").Name("delete one Permission").Do(func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err = con.service.DeletePermission(c.UserContext(), id); err != nil {
			return err
		}

		return request.Resp(c, request.Response{
			Message: "The Permission was deleted successfully!",
		})
	})
}
