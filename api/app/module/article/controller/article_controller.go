package controller

import (
	"strconv"

	"gva/app/module/article/dto"
	"gva/app/module/article/service"

	"gva/internal/control_route"
	"gva/utils/response"

	"github.com/gofiber/fiber/v2"
)

var _ interface {
	control_route.FiberRouter
} = (*ArticleController)(nil)

type ArticleController struct {
	service *service.ArticleService
}

func (con *ArticleController) Routes(r fiber.Router) {

	r.Route(
		"/articles", func(router fiber.Router) {
			router.Get("/", con.List).Name("get many articles")
			router.Get("/:id", con.Get).Name("get one article")
			router.Post("/", con.Create).Name("create one article")
			router.Patch("/:id", con.Update).Name("update one article")
			router.Delete("/:id", con.Delete).Name("delete one article")
		},
	)
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{
		service: service,
	}
}

func (con *ArticleController) List(c *fiber.Ctx) error {
	list, err := con.service.GetArticles(c.UserContext())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "Article list retreived successfully!",
		Data:    list,
	})
}

func (con *ArticleController) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := con.service.GetArticleByID(c.UserContext(), id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The article retrieved successfully!",
		Data:    data,
	})
}

func (con *ArticleController) Create(c *fiber.Ctx) error {
	req := new(dto.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.CreateArticle(c.UserContext(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The article was created successfully!",
		Data:    data,
	})
}

func (con *ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(dto.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	data, err := con.service.UpdateArticle(c.UserContext(), id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The article was updated successfully!",
		Data:    data,
	})
}

func (con *ArticleController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.service.DeleteArticle(c.UserContext(), id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Message: "The article was deleted successfully!",
	})
}
