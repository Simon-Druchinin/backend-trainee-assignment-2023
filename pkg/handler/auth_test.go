package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"
	"user_segmentation"
	"user_segmentation/pkg/service"
	service_mocks "user_segmentation/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
)

func TestHabdler_register(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockAuthorization, user user_segmentation.User)

	tests := []struct {
		name                 string
		inputUser            user_segmentation.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputUser: user_segmentation.User{},
			mockBehavior: func(r *service_mocks.MockAuthorization, user user_segmentation.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Service Error",
			inputUser: user_segmentation.User{},
			mockBehavior: func(r *service_mocks.MockAuthorization, user user_segmentation.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := &Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/auth/register", handler.register)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/register",
				bytes.NewBufferString(""))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
