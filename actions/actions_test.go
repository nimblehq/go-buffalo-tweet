package actions

import (
	"testing"

	"github.com/bufftwitt/models"
	"github.com/gobuffalo/suite"
	"github.com/markbates/pop/nulls"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func (as *ActionSuite) Login() *models.User {
	user := &models.User{
		Name:       "Trung",
		Email:      nulls.NewString("hello@nimbl3.com"),
		Provider:   "something",
		ProviderID: "123",
	}

	as.NoError(as.DB.Create(user))
	as.Session.Set("current_user_id", user.ID)
	return user
}

func (as *ActionSuite) CreateTweet(user *models.User) *models.Tweet {
	tweet := &models.Tweet{
		Message: "Hello world!",
		UserID:  user.ID,
	}

	res, err := as.DB.ValidateAndCreate(tweet)
	as.NoError(err)
	as.False(res.HasAny())

	return tweet
}
