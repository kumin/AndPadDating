package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/erroz"
	"github.com/kumin/BityDating/handler"
	"github.com/kumin/BityDating/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(
	authService *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	if err := handler.ValidateCreateUser(&user); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	registeredUser, err := a.authService.Register(c.Request.Context(), &user)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}

	c.IndentedJSON(http.StatusOK, registeredUser)
}

func (a *AuthHandler) Login(c *gin.Context) {
	phone, ok := c.GetQuery("phone")
	if !ok {
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(erroz.ErrPhoneIsMissing))
	}
	loginUser, err := a.authService.Login(c.Request.Context(), phone)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, loginUser)
}
