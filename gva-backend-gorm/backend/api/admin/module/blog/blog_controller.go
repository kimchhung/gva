package blog

import (
	"backend/api/admin/module/blog/dto"
	"backend/app/common/permission"
	"backend/app/common/service"
	"backend/internal/ctr"
	"backend/internal/request"
	"backend/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*BlogController)(nil)

type BlogController struct {
	service *BlogService
	jwt_s   *service.JwtService
}

func NewBlogController(service *BlogService, jwt_s *service.JwtService) *BlogController {
	return &BlogController{
		service: service,
		jwt_s:   jwt_s,
	}
}

func (con *BlogController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/blog", con.jwt_s.RequiredAdmin()),
	)
}

// @Tags			Blog
// @Summary			Create Blog
// @Description		Create Blog
// @ID				create-blog
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			body	body		dto.CreateBlogRequest	true	"Blog data"
// @Success			200		{object}	response.Response{data=dto.BlogResponse}	"Successfully created Blog"
// @Router			/blog [post]
func (con *BlogController) Create() *ctr.Route {
	return ctr.POST("/").Do(func() []ctr.H {
		body := new(dto.CreateBlogRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogAdd),
			request.Validate(
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.CreateBlog(c.Request().Context(), body)
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

// @Tags			Blog
// @Summary			Get Blog
// @Description		Get Blog
// @ID				get-blog
// @Accept			json
// @Param			id	path	int	true	"Blog ID"
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=dto.BlogResponse}	"Successfully get Blog"
// @Router			/blog/{id} [get]
func (con *BlogController) Get() *ctr.Route {
	return ctr.GET("/:id").Do(func() []ctr.H {
		params := new(dto.GetBlogRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogView),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				data, err := con.service.GetBlog(c.Request().Context(), params.ID)
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

// @Tags			Blog
// @Summary			Get Blogs
// @Description		Get Blogs
// @ID				get-blogs
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
// @Success			200		{object}	response.Response{data=[]dto.BlogResponse}	"Successfully get Blogs"
// @Router			/blog [get]
func (con *BlogController) GetMany() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		query := new(dto.GetManyQuery)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogView),
			request.Validate(
				request.PaginateParser(&query.QueryDto),
			),
			func(c echo.Context) error {

				list, meta, err := con.service.GetBlogs(c.Request().Context(), query)
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

// @Tags			Blog
// @Summary			Update Blog
// @Description		Update Blog
// @ID				update-blog
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Blog ID"
// @Param			body	body		dto.UpdateBlogRequest	true	"Blog data"
// @Success			200		{object}	response.Response{data=dto.BlogResponse}	"Successfully updated Blog"
// @Router			/blog/{id} [put]
func (con *BlogController) Update() *ctr.Route {
	return ctr.PUT("/:id").Do(func() []ctr.H {
		params := new(dto.GetBlogRequest)
		body := new(dto.UpdateBlogRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdateBlog(c.Request().Context(), params.ID, body)
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

// @Tags			Blog
// @Summary			Update Blog partial
// @Description		Update Blog partial
// @ID				update-blog-partial
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path int true "Blog ID"
// @Param			body	body	dto.UpdatePatchBlogRequest	true	"Blog data"
// @Success			200		{object}	response.Response{data=dto.BlogResponse}	"Successfully updated Blog"
// @Router			/blog/{id} [patch]
func (con *BlogController) Patch() *ctr.Route {
	return ctr.PATCH("/:id").Do(func() []ctr.H {
		params := new(dto.GetBlogRequest)
		body := new(dto.UpdatePatchBlogRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogEdit),
			request.Validate(
				request.ParamsParser(params),
				request.BodyParser(body),
			),
			func(c echo.Context) error {
				data, err := con.service.UpdatePatchBlog(c.Request().Context(), params.ID, body)
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

// @Tags			Blog
// @Summary			Delete Blog
// @Description		Delete Blog
// @ID				delete-blog
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Param			id	path	int	true	"Blog ID"
// @Success			200		{object}	response.Response{data=string}	"Successfully deleted Blog"
// @Router			/blog/{id} [delete]
func (con *BlogController) Delete() *ctr.Route {
	return ctr.DELETE("/:id").Do(func() []ctr.H {
		params := new(dto.GetBlogRequest)

		return []ctr.H{
			permission.RequireAnyWithSuper(permission.BlogDelete),
			request.Validate(
				request.ParamsParser(params),
			),
			func(c echo.Context) error {
				err := con.service.DeleteBlog(c.Request().Context(), params.ID)
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data("Successfully deleted Blog"),
				)
			},
		}
	})
}
