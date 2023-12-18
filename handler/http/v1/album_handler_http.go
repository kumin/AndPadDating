package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/handler"
	"github.com/kumin/BityDating/services"
)

type AlbumHandler struct {
	albumService *services.AlbumService
}

func NewAlbumHandler(
	albumService *services.AlbumService,
) *AlbumHandler {
	return &AlbumHandler{
		albumService: albumService,
	}
}

func (a *AlbumHandler) CreateOne(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	image, err := a.albumService.CreateOne(c.Request.Context(), file)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, image)
}

func (a *AlbumHandler) CreateMany(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	files := form.File["files"]
	images, err := a.albumService.CreateMany(c.Request.Context(), files)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, images)
}

func (a *AlbumHandler) GetUserAlbum(c *gin.Context) {
	album, err := a.albumService.GetUserAlbum(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, album)
}
