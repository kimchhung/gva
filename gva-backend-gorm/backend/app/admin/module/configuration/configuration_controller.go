package configuration

import (
	"backend/app/admin/module/configuration/dto"
	"backend/app/share/permission"
	"backend/app/share/service"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/internal/ctr"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*ConfigurationController)(nil)

type ConfigurationController struct {
	service *ConfigurationService
	jwt_s   *service.JwtService
}

func NewConfigurationController(service *ConfigurationService, jwt_s *service.JwtService) *ConfigurationController {
	return &ConfigurationController{
		service: service,
		jwt_s:   jwt_s,
	}
}

func (con *ConfigurationController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/configuration", con.jwt_s.RequiredAdmin()),
	)
}

// @Tags			Configuration
// @Summary			Create Configuration
// @Description		Create Configuration
// @ID				create-configuration
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreateConfigurationRequest	true	"Configuration data"
// @Success			200		{object}	response.Response{data=dto.ConfigurationResponse}	"Successfully created Configuration"
// @Router			/configuration [post]
func (con *ConfigurationController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.CreateConfigurationRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationAdd),
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.CreateConfiguration(c.Request().Context(), body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Get Configuration
// @Description		Get Configuration
// @ID				get-configuration
// @Accept			json
// @Param			id	path	int	true	"Configuration ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.ConfigurationResponse}	"Successfully get Configuration"
// @Router			/configuration/{id} [get]
func (con *ConfigurationController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetConfigurationRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetConfiguration(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Get Tags
// @Description		Get Tags
// @ID				get-configurations
// @Accept			json
// @Param 			page	query	int	false	"page"
// @Param 			limit	query	int	false	"limit"
// @Param 			search	query	string	false	"search"
// @Param 			filters	query	string	false	"filters"
// @Param 			sorts	query	string	false	"sorts"
// @Param 			orders	query	string	false	"orders"
// @Param 			selects	query	string	false	"selects: list, totalCount"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]dto.ConfigurationResponse}	"Successfully get Tags"
// @Router			/configuration [get]
func (con *ConfigurationController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {
				list, meta, err := con.service.GetConfigurations(c.Request().Context(), query)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(list),
					response.Meta(meta),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Update Configuration
// @Description		Update Configuration
// @ID				update-configuration
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Configuration ID"
// @Param			body	body		dto.UpdateConfigurationRequest	true	"Configuration data"
// @Success			200		{object}	response.Response{data=dto.ConfigurationResponse}	"Successfully updated Configuration"
// @Router			/configuration/{id} [put]
func (con *ConfigurationController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetConfigurationRequest)
		body := new(dto.UpdateConfigurationRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateConfiguration(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Update Configuration partial
// @Description		Update Configuration partial
// @ID				update-configuration-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Configuration ID"
// @Param			body	body	dto.UpdatePatchConfigurationRequest	true	"Configuration data"
// @Success			200		{object}	response.Response{data=dto.ConfigurationResponse}	"Successfully updated Configuration"
// @Router			/configuration/{id} [patch]
func (con *ConfigurationController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetConfigurationRequest)
		body := new(dto.UpdatePatchConfigurationRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatchConfiguration(c.Request().Context(), params.ID, body)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Delete Configuration
// @Description		Delete Configuration
// @ID				delete-configuration
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Configuration ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted Configuration"
// @Router			/configuration/{id} [delete]
func (con *ConfigurationController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetConfigurationRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationDelete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteConfiguration(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted Configuration"),
				)
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Get Configuration by Key
// @Description		Get Configuration by Key
// @ID				get-configuration-by-key
// @Accept			json
// @Param			key	path	string	true	"Configuration Key"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.ConfigurationResponse}	"Successfully get Configuration"
// @Router			/configuration/key/{key} [get]
func (con *ConfigurationController) GetByKey() *ctr.Route {
	return ctr.GET("/key/:key").Do(func() []ctr.H {
		params := new(dto.GetConfigurationByKeyRequest)

		return []ctr.H{
			permission.RequireAny(permission.ConfigurationView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetConfigurationByKey(c.Request().Context(), params.Key)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(data),
				)
			},
		}
	})
}
