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

func Test_GetAssetsByDepartmentId(t *testing.T) {
	scenarios := []struct {
		name         string
		context      func() (*gin.Context, *httptest.ResponseRecorder)
		mockService  func(ctrl *gomock.Controller) services.IAssetService
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success case",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/assets?department_id=1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IAssetService {
				m := services.NewMockIAssetService(ctrl)
				m.EXPECT().GetAssetsByDepartmentId(uint32(1)).Return([]*models.Asset{}, nil)
				return m
			},
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name: "Error case: Invalid department_id",
			context: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/assets?department_id=1", nil)
				return c, w
			},
			mockService: func(ctrl *gomock.Controller) services.IAssetService {
				m := services.NewMockIAssetService(ctrl)
				m.EXPECT().GetAssetsByDepartmentId(uint32(1)).Return(nil, errors.New("error"))
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

			funcGetAssetHandler = func() services.IAssetService {
				return scenario.mockService(ctrl)
			}

			c, recorder := scenario.context()
			handler := NewAssetHandler()
			handler.GetAssetsByDepartmentId(c)
			assert.Equal(t, scenario.expectedCode, recorder.Code)
			assert.Equal(t, scenario.expectedBody, recorder.Body.String())
		})
	}
}
