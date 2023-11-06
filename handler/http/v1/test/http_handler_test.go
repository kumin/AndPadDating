//go:build integration
// +build integration

package httphandlertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kumin/AndPadDating/entities"
	mocks "github.com/kumin/AndPadDating/mocks/data"
	"github.com/stretchr/testify/suite"
)

type HttpHandlerTestSuite struct {
	suite.Suite
	httpHandler *HttpHandler
	router      *gin.Engine
}

func (h *HttpHandlerTestSuite) SetupTest() {
	h.httpHandler, _ = BuildHttpHandler()
	h.router = gin.Default()
}

func (h *HttpHandlerTestSuite) TestAuthHandler_Register() {
	h.router.POST("/auth/register", h.httpHandler.authHandler.Register)
	body, _ := json.Marshal(mocks.Users[0])
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.router.ServeHTTP(w, req)

	var user entities.RegisteredUser
	_ = json.Unmarshal(w.Body.Bytes(), &user)
	h.Equal(http.StatusOK, w.Code)
	h.Equal(user.User.Phone, mocks.Users[0].Phone)
}

func TestHttpHandler(t *testing.T) {
	suite.Run(t, new(HttpHandlerTestSuite))
}
