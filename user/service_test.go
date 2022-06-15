package user_test

import (
	"bankapp/app"
	"bankapp/db"
	storemocks "bankapp/db/mocks"
	"bankapp/user"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceTestSuite struct {
	suite.Suite
	store   *storemocks.Storer
	logger  *zap.SugaredLogger
	service user.Service
}

func init() {
	app.InitLogger()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.logger = app.GetLogger()
	suite.store = &storemocks.Storer{}
	suite.service = user.NewService(suite.store, suite.logger)
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.store.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestUpdate() {
	t := suite.T()
	t.Run("when user created successfully", func(t *testing.T) {
		suite.SetupTest()
		updateRequest := user.UpdateRequest{
			Name:     "john",
			Password: "example",
		}
		userID := "example-id"
		dbUser := db.User{
			Password: updateRequest.Password,
			Name:     updateRequest.Name,
			ID:       userID,
		}
		suite.store.On("UpdateUser", context.Background(), &dbUser).Return(nil).Once()

		gotErr := suite.service.Update(context.Background(), updateRequest, userID)

		assert.NoError(t, gotErr)
		suite.TearDownTest()
	})
}
