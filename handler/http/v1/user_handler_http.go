package http_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/handler"
	"github.com/kumin/BityDating/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(
	userService *services.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)
		return
	}
	if err := handler.ValidateCreateUser(&user); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	_, err := u.userService.CreateUser(c.Request.Context(), &user)

	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	user, err := u.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)
		return
	}
	user.Id = id
	_, err = u.userService.UpdateUser(c.Request.Context(), &user)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	if err := u.userService.DeleteUser(c.Request.Context(), id); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete successfully"})
}

func (u *UserHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
	}
	fileUrl, err := u.userService.UploadFile(c.Request.Context(), file)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
	}

	c.IndentedJSON(http.StatusOK, gin.H{"file_url": fileUrl})
}
