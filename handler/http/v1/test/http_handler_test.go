//go:build integration
// +build integration

package httphandlertest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/services"
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
	h.router.POST("/auth/register", h.httpHandler.authHandler.Register)
	h.router.POST("/auth/login", h.httpHandler.authHandler.Login)
	h.router.POST("/matching/like", h.httpHandler.matchingHandler.CreateMatching)
	h.router.GET("/matching/whoilike/:userid", h.httpHandler.matchingHandler.WhoILike)
	h.router.GET("/matching/wholikeme/:userid", h.httpHandler.matchingHandler.WhoLikeMe)
	h.router.GET("/matching/list/:userid", h.httpHandler.matchingHandler.ListMatching)
}

func (h *HttpHandlerTestSuite) _TestAuthHandler() {
	user := h.getTestUser()

	//test register
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.router.ServeHTTP(w, req)
	var registerdUser entities.RegisteredUser
	_ = json.Unmarshal(w.Body.Bytes(), &registerdUser)
	h.Equal(http.StatusOK, w.Code)
	h.Equal(registerdUser.User.Phone, user.Phone)

	//test login
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/auth/login?phone=%s", user.Phone),
		http.NoBody)
	w = httptest.NewRecorder()
	h.router.ServeHTTP(w, req)
	var loginUser entities.RegisteredUser
	_ = json.Unmarshal(w.Body.Bytes(), &loginUser)
	h.Equal(http.StatusOK, w.Code)
	ok := services.ValidateToken(loginUser.Token)
	h.True(ok)
}

func (h *HttpHandlerTestSuite) TestMatchingHandler() {
	user1 := h.getTestUser()
	user1, err := h.httpHandler.userService.CreateUser(context.Background(), user1)
	h.Nil(err)
	user2 := h.getTestUser()
	user2, err = h.httpHandler.userService.CreateUser(context.Background(), user2)
	h.Nil(err)
	user3 := h.getTestUser()
	user3, err = h.httpHandler.userService.CreateUser(context.Background(), user3)
	h.Nil(err)

	//test like
	body, _ := json.Marshal(&entities.UserMatching{
		UserId:    user1.Id,
		PartnerId: user2.Id,
		IsLike:    1,
	})
	req, _ := http.NewRequest(http.MethodPost, "/matching/like", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.router.ServeHTTP(w, req)
	var matchingUser entities.UserMatching
	_ = json.Unmarshal(w.Body.Bytes(), &matchingUser)
	h.Equal(http.StatusOK, w.Code)
	h.Equal(user1.Id, matchingUser.UserId)

	body, _ = json.Marshal(&entities.UserMatching{
		UserId:    user2.Id,
		PartnerId: user1.Id,
		IsLike:    1,
	})
	req, _ = http.NewRequest(http.MethodPost, "/matching/like", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.router.ServeHTTP(w, req)
	var matchingUser2 entities.UserMatching
	_ = json.Unmarshal(w.Body.Bytes(), &matchingUser2)
	h.Equal(http.StatusOK, w.Code)
	h.Equal(user2.Id, matchingUser2.UserId)

	body, _ = json.Marshal(&entities.UserMatching{
		UserId:    user3.Id,
		PartnerId: user1.Id,
		IsLike:    0,
	})
	req, _ = http.NewRequest(http.MethodPost, "/matching/like", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.router.ServeHTTP(w, req)
	var matchingUser3 entities.UserMatching
	_ = json.Unmarshal(w.Body.Bytes(), &matchingUser3)
	h.Equal(http.StatusOK, w.Code)
	h.Equal(user3.Id, matchingUser3.UserId)

	//test whoilike
	req4, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/matching/whoilike/%d", user1.Id), http.NoBody)
	w4 := httptest.NewRecorder()
	h.router.ServeHTTP(w4, req4)
	var users4 []*entities.User
	_ = json.Unmarshal(w4.Body.Bytes(), &users4)
	h.Equal(http.StatusOK, w4.Code)
	h.Len(users4, 1)
	h.Equal(user2.Id, users4[0].Id)

	//test wholikeme
	req5, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/matching/wholikeme/%d", user2.Id), http.NoBody)
	w5 := httptest.NewRecorder()
	h.router.ServeHTTP(w5, req5)
	var users5 []*entities.User
	_ = json.Unmarshal(w5.Body.Bytes(), &users5)
	h.Equal(http.StatusOK, w5.Code)
	h.Len(users5, 1)
	h.Equal(user1.Id, users5[0].Id)

	//test list matching
	req6, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/matching/list/%d?page=0&limit=10", user1.Id), http.NoBody)
	w6 := httptest.NewRecorder()
	h.router.ServeHTTP(w6, req6)
	var users6 []*entities.User
	_ = json.Unmarshal(w6.Body.Bytes(), &users6)
	h.Equal(http.StatusOK, w6.Code)
	h.Len(users6, 1)
	h.Equal(user2.Id, users6[0].Id)
}

func (h *HttpHandlerTestSuite) getTestUser() *entities.User {
	user := &entities.User{
		Username: fmt.Sprintf("tester%d", time.Now().Unix()),
		Phone:    fmt.Sprintf("%d", 84389728300+time.Now().Unix()+rand.Int63n(100)),
		Email:    "test9@.com",
		BirthDay: time.Now(),
		Gender:   "female",
		IsActive: 1,
	}

	return user
}

func TestHttpHandler(t *testing.T) {
	suite.Run(t, new(HttpHandlerTestSuite))
}
