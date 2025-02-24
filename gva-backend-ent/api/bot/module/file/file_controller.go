package file

import (
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*FileController)(nil)

type FileController struct {
	service *FileService
}

func (con *FileController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("/file"),
	)
}

func NewFileController(service *FileService) *FileController {
	return &FileController{
		service: service,
	}
}

// @Tags        File
// @Summary     Serve static files
// @Description Serves files from `storage/static` directory
// @ID          serve-static-files
// @Success     200
// @Param       name path string true "filename"
// @Router      /file/static/img/{name} [get]
func (con *FileController) Static() *ctr.Route {
	return ctr.Add(func(g *echo.Group) {
		g.Static("/static/img", "storage/static")
	})
}

// @Tags        File
// @Summary     Upload a file
// @Description Upload a file
// @ID          upload-file
// @Accept      multipart/form-data
// @Produce     json
// @Param       file formData file true "File"
// @Success     200 {object} response.Response{}
// @Router      /file/upload-img [post]
func (con *FileController) UploadImg() *ctr.Route {
	return ctr.POST("/upload-img").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				fileHeader, err := c.FormFile("file")
				if err != nil {
					return err
				}

				data, err := con.service.UploadFile(c.Request().Context(), fileHeader)
				if err != nil {
					return err
				}

				return request.Response(c, response.Data(data))
			},
		}
	})
}
