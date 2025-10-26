package index

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	admincontext "backend/app/admin/context"
	middleware "backend/app/admin/middleware"
	"backend/app/admin/module/configuration"
	"backend/app/admin/module/index/dto"
	apperror "backend/app/share/error"
	"backend/app/share/permission"
	"backend/app/share/service"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"backend/env"
	"backend/internal/ctr"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*IndexController)(nil)

type IndexController struct {
	cfg        *env.Config
	index_s    *IndexService
	s3_s       *service.S3Service
	ip_s       *service.IPService
	config_s   *configuration.ConfigurationService
	middleware *middleware.Middleware
}

func (con *IndexController) Init() *ctr.Ctr {
	return ctr.New()
}

func NewIndexController(
	cfg *env.Config,
	index_s *IndexService,
	s3_s *service.S3Service,
	ip_s *service.IPService,
	jwt_s *service.JwtService,
	config_s *configuration.ConfigurationService,
	middleware *middleware.Middleware,
) *IndexController {
	return &IndexController{
		index_s:    index_s,
		ip_s:       ip_s,
		s3_s:       s3_s,
		cfg:        cfg,
		config_s:   config_s,
		middleware: middleware,
	}
}

// @Tags        Index
// @Summary     Current Server Time
// @ID          now
// @Produce     json
// @Success     200 {object} response.Response{data=string} "format time.RFC3339"
// @Router      /now [get]
// @Security    Bearer
func (con *IndexController) Now() *ctr.Route {
	return ctr.GET("/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				now, err := con.index_s.Now(c.Request().Context())
				if err != nil {
					return err
				}

				return request.Response(c, response.Data(now.UTC().Format(time.RFC3339)))
			},
		}
	})
}

// @Tags        Index
// @Summary     Config
// @ID          config
// @Produce     json
// @Success     200 {object} response.Response{data=dto.ConfigResponse}
// @Router      /config [get]
func (con *IndexController) Config() *ctr.Route {
	return ctr.GET("/config").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				now, err := con.index_s.Now(c.Request().Context())
				if err != nil {
					return err
				}

				configResponse := dto.ConfigResponse{
					Now:      now.UTC().Format(time.RFC3339),
					PublicIp: con.ip_s.GetCurrentIP(c),
				}

				return request.Response(c, response.Data(configResponse))
			},
		}
	})
}

// @Tags			Index
// @Summary			Get Permission Scopes
// @Description		Get Permission Scopes
// @ID				get-permission-scopes
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]string}
// @Router			/permission-scope [get]
func (con *IndexController) PermissionScope() *ctr.Route {
	return ctr.GET("/permission-scope").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				if con.cfg.IsProd() {
					return apperror.ErrNotFound
				}

				return request.Response(c,
					response.Data(permission.Scopes()),
				)
			},
		}
	})
}

// @Tags        Index
// @Summary     Upload Image
// @ID          upload
// @Accept      multipart/form-data
// @Produce     json
// @Param       file formData file true "Image File"
// @Success     200 {object} response.Response{data=string} "Image URL"
// @Router      /upload/image [post]
// @Security    Bearer
func (con *IndexController) Upload() *ctr.Route {
	return ctr.POST("/upload/image").
		Use(
			con.middleware.JwtGuard(),
			con.middleware.IpGuard(),
		).
		Do(func() []ctr.H {
			return []ctr.H{
				func(c echo.Context) error {
					file, err := c.FormFile("file")
					if err != nil {
						return err
					}

					uploadObject, err := con.s3_s.UploadFile(file)
					if err != nil {
						return err
					}

					// appctx.SetRequestParams(c.Request().Context(), map[string]interface{}{
					// 	"filename":    uploadObject.Filename,
					// 	"size":        file.Size,
					// 	"mime":        file.Header.Get("Content-Type"),
					// 	"uploadedURL": uploadObject.URL,
					// })

					return request.Response(c, response.Data(uploadObject))
				},
			}
		})
}

// @Tags        Index
// @Summary     Serve static files
// @Description Serves files from `storage/docs` directory
// @ID          serve-static-files
// @Success     200
// @Param       name path string true "filename"
// @Router      /file/docs/{name} [get]
func (con *IndexController) Static() *ctr.Route {
	mdToHtml := func(md []byte) []byte {
		// create markdown parser with extensions
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
		p := parser.NewWithExtensions(extensions)
		doc := p.Parse(md)

		// create HTML renderer with extensions
		htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.LazyLoadImages
		opts := html.RendererOptions{Flags: htmlFlags}
		renderer := html.NewRenderer(opts)

		content := markdown.Render(doc, renderer)

		// wrap rendered content with GitHub markdown style
		return []byte(fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Document</title>
				<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/4.0.0/github-markdown.min.css">
				<style>
					body {
						box-sizing: border-box;
						min-width: 200px;
						max-width: 980px;
						margin: 0 auto;
						padding: 45px;
					}
				</style>
			</head>
			<body class="markdown-body">
				%s
			</body>
			</html>
		`, content))
	}

	return ctr.GET("/file/docs/:name").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				name := c.Param("name")
				filePath := fmt.Sprintf("storage/docs/%s", name)

				if !strings.HasSuffix(filePath, ".md") {
					return apperror.NewError(apperror.ErrUnsupportedFileFormat, nil)
				}

				data, err := os.ReadFile(filePath)
				if err != nil {
					return apperror.NewError(apperror.ErrUnknownError, apperror.Join(err))
				}

				output := mdToHtml(data)
				return c.HTML(http.StatusOK, string(output))
			},
		}
	})
}

// @Tags			Configuration
// @Summary			Get Docs Configuration
// @Description		Get Docs Configuration
// @ID				get-docs-configuration
// @Accept			json
// @Produce			json
// @Security		Bearer
// @Success			200		{object}	response.Response{data=[]dto.ConfigurationResponse}	"Successfully get Docs Configuration"
// @Router			/config/docs [get]
func (con *IndexController) GetDocs() *ctr.Route {
	return ctr.GET("/config/docs").Use(con.middleware.JwtGuard()).Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				adminCtx := admincontext.MustAdminContext(c.Request().Context())
				if !adminCtx.IsSuperAdmin() {
					if !slices.Contains(adminCtx.PermissionScopes(), string(permission.DocumentView)) {
						return request.Response(c, response.Data([]any{}))
					}
				}

				data, err := con.config_s.GetDocsConfiguration(c.Request().Context())
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
