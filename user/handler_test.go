package user_test

import (
	"bankapp/api"
	"bankapp/user"
	"bankapp/user/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	service *mocks.Service
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.service = &mocks.Service{}
}

func (suite *HandlerTestSuite) TearDownTest() {
	t := suite.T()
	suite.service.AssertExpectations(t)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestUpdateHandler() {
	t := suite.T()
	userID := "example-user-id"
	reqURL := "/users/example-user-id"
	t.Run("when user updated successfully", func(t *testing.T) {
		suite.SetupTest()
		updateRequest := user.UpdateRequest{
			Name:     "john",
			Password: "example-password",
		}
		requestBody, err := json.Marshal(updateRequest)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, reqURL, bytes.NewBuffer(requestBody))
		vars := map[string]string{
			"userId": userID,
		}
		req = mux.SetURLVars(req, vars)
		rw := httptest.NewRecorder()
		response := api.Response{Message: "Updated Successfully"}
		suite.service.On("Update", req.Context(), updateRequest, userID).Return(nil).Once()

		user.Update(suite.service)(rw, req)

		var gotResponse api.Response
		err = json.Unmarshal(rw.Body.Bytes(), &gotResponse)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Equal(t, response, gotResponse)
		suite.TearDownTest()
	})
}
