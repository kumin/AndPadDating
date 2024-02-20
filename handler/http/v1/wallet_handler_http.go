package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/configs"
	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/handler"
	"github.com/kumin/BityDating/pkg/numberx"
	"github.com/kumin/BityDating/services"
)

type WalletHandler struct {
	walletService *services.WalletService
}

func NewWalletHandler(
	walletService *services.WalletService,
) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (w *WalletHandler) CreateTransaction(c *gin.Context) {
	var transaction entities.WalletTransaction
	if err := c.BindJSON(&transaction); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage)
		return
	}
	_, err := w.walletService.CreateTransaction(c.Request.Context(), &transaction)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, handler.ErrorMessage)
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (w *WalletHandler) ListTransactions(c *gin.Context) {
	id, err := handler.GetParam("userid", &c.Params)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	userId, err := numberx.ParseInt(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	page := 0
	if pageStr, ok := c.GetQuery("page"); ok {
		val, err := numberx.ParseInt(pageStr)
		if err == nil {
			page = int(val)
		}
	}
	limit := 10
	if limitStr, ok := c.GetQuery("limit"); ok {
		val, err := numberx.ParseInt(limitStr)
		if err == nil {
			limit = int(val)
		}
	}
	if limit > configs.MaxPageLimit {
		limit = configs.MaxPageLimit
	}
	transaction, err := w.walletService.ListTransactions(c.Request.Context(), userId, page, limit)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, handler.ErrorMessage(err))
		return
	}
	c.JSON(http.StatusOK, transaction)
}
