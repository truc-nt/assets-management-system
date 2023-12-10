package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/models"
	"server/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetUsers(t *testing.T) {
	scenarios := []struct {
		name         string
		context      func() (*gin.Context, *httptest.ResponseRecorder)
		mockService  func(ctrl *gomock.Controller) services.IUserService
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success case",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/users?user_id=1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IUserService {
				m := services.NewMockIUserService(ctrl)
				param := &models.GetUsersParam{
					Role: 1,
				}
				m.EXPECT().GetUsers(param).Return([]*models.User{}, nil)
				return m
			},
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name: "Error case: Invalid user_id",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/users?user_id=1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IUserService {
				m := services.NewMockIUserService(ctrl)
				param := &models.GetUsersParam{
					Role: 1,
				}
				m.EXPECT().GetUsers(param).Return(nil, errors.New("error"))
				return m
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"Server Error"}`,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			funcGetUserHandler = func() services.IUserService {
				return scenario.mockService(ctrl)
			}

			c, recorder := scenario.context()
			handler := NewUserHandler()
			handler.GetUsers(c)
			assert.Equal(t, scenario.expectedCode, recorder.Code)
			assert.Equal(t, scenario.expectedBody, recorder.Body.String())
		})
	}
}

func Test_GetUserById(t *testing.T) {
	scenarios := []struct {
		name         string
		context      func() (*gin.Context, *httptest.ResponseRecorder)
		mockService  func(ctrl *gomock.Controller) services.IUserService
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success case",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Params = []gin.Param{{Key: "user_id", Value: "1"}}
				c.Request = httptest.NewRequest("GET", "/users/1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IUserService {
				m := services.NewMockIUserService(ctrl)
				m.EXPECT().GetUserById(uint32(1)).Return(&models.User{
					ID:        1,
					Username:  "test",
					Password:  "test",
					Login:     false,
					Role:      1,
					Telephone: "0123456789",
				}, nil)
				return m
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"username":"test","password":"test","login":0,"role":1,"telephone":"0123456789"}`,
		},
		{
			name: "Error case: Invalid user_id",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Params = []gin.Param{{Key: "user_id", Value: "1"}}
				c.Request = httptest.NewRequest("GET", "/users/1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IUserService {
				m := services.NewMockIUserService(ctrl)
				m.EXPECT().GetUserById(uint32(1)).Return(nil, errors.New("error"))
				return m
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"Server Error"}`,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			funcGetUserHandler = func() services.IUserService {
				return scenario.mockService(ctrl)
			}

			c, recorder := scenario.context()
			handler := NewUserHandler()
			handler.GetUserById(c)
			assert.Equal(t, scenario.expectedCode, recorder.Code)
			assert.Equal(t, scenario.expectedBody, recorder.Body.String())
		})
	}
}
