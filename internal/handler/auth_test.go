package handler_test

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/quanergyO/avito_assingment/internal/service"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user types.SignInInput)
	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            types.SignInInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"Test","password":"test"}`,
			inputUser: types.UserType{
				Username: "Test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user types.UserType) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty Fields",
			inputBody:            `{"username":"Test"}`,
			mockBehavior:         func(s *mock_service.MockAuthorization, user types.UserType) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"username":"Test","password":"test"}`,
			inputUser: types.UserType{
				Username: "Test",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user types.UserType) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init deps
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
