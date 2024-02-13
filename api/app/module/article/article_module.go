package article

import (
	"fmt"
	"gva/app/module/article/controller"
	"gva/app/module/article/repository"
	"gva/app/module/article/service"
	"gva/internal/control_route"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var _ interface {
	control_route.Router
} = &ArticleRouter{}

type ArticleRouter struct {
	app        fiber.Router
	controller *controller.ArticleController
}

// Router methods
func NewArticleRouter(fiber *fiber.App, controller *controller.ArticleController) *ArticleRouter {
	return &ArticleRouter{
		app:        fiber,
		controller: controller,
	}
}

func (r *ArticleRouter) Register() {
	r.controller.Routes(r.app)

	fmt.Print("register router")
}

// Register bulkly
var NewArticleModule = fx.Module("ArticleModule",
	// Register Repository & Service
	fx.Provide(repository.NewArticleRepository),
	fx.Provide(service.NewArticleService),

	// Regiser Controller
	fx.Provide(controller.NewArticleController),

	// Register Router
	fx.Provide(fx.Annotate(
		NewArticleRouter,
		fx.As(new(control_route.Router)),
		fx.ResultTags(`group:"routers"`),
	)),
)
