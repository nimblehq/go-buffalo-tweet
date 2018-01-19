package actions

import "github.com/bufftwitt/models"

func (as *ActionSuite) Test_TweetsResource_List() {
	user := as.Login()

	tweet := &models.Tweet{
		UserID: user.ID,
		Message: "Hello world!!",
	}
	as.NoError(as.DB.Create(tweet))

	res := as.HTML("/tweets").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), tweet.Message)
}
