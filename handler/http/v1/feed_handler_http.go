package http_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kumin/AndPadDating/configs"
	"github.com/kumin/AndPadDating/handler"
	"github.com/kumin/AndPadDating/services"
)

type FeedHandler struct {
	feedService *services.FeedService
}

func NewFeedHandler(
	feedService *services.FeedService,
) *FeedHandler {
	return &FeedHandler{
		feedService: feedService,
	}
}

func (m *FeedHandler) GetFeed(c *gin.Context) {
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
	users, err := m.feedService.GetFeed(c.Request.Context(), userId, page, limit)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, users)
}
