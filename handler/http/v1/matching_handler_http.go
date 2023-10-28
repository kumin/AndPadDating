package http_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kumin/AndPadDating/configs"
	"github.com/kumin/AndPadDating/entities"
	"github.com/kumin/AndPadDating/handler"
	"github.com/kumin/AndPadDating/services"
)

type MatchingHandler struct {
	matchingService *services.MatchingService
}

func NewMatchingHandler(
	matchingService *services.MatchingService,
) *MatchingHandler {
	return &MatchingHandler{
		matchingService: matchingService,
	}
}

func (m *MatchingHandler) CreateMatching(c *gin.Context) {
	var matching entities.UserMatching
	if err := c.BindJSON(&matching); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	if err := handler.ValidateCreateMatching(&matching); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	_, err := m.matchingService.CreateMatching(c.Request.Context(), &matching)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, matching)
}

func (m *MatchingHandler) WhoILike(c *gin.Context) {
	id, err := handler.GetParam("userid", &c.Params)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	users, err := m.matchingService.WhoILike(c.Request.Context(), userId)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, users)
}

func (m *MatchingHandler) WhoLikeMe(c *gin.Context) {
	id, err := handler.GetParam("partnerid", &c.Params)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}

	users, err := m.matchingService.WhoLikeMe(c.Request.Context(), userId)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, users)
}

func (m *MatchingHandler) ListMatching(c *gin.Context) {
	id, err := handler.GetParam("userid", &c.Params)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	page := 0
	if pageStr, ok := c.GetQuery("page"); ok {
		val, err := strconv.ParseInt(pageStr, 10, 64)
		if err == nil {
			page = int(val)
		}
	}
	limit := 10
	if limitStr, ok := c.GetQuery("limit"); ok {
		val, err := strconv.ParseInt(limitStr, 10, 32)
		if err == nil {
			limit = int(val)
		}
	}
	if limit > configs.MaxPageLimit {
		limit = configs.MaxPageLimit
	}
	users, err := m.matchingService.ListMatching(c.Request.Context(), userId, page, limit)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, users)
}
