package actions

import (
	"github.com/bufftwitt/models"
	"github.com/markbates/pop/nulls"
)

func (as *ActionSuite) Test_TweetsResource_List() {
	user := as.Login()

	tweet := &models.Tweet{
		UserID:  user.ID,
		Message: "Hello world!!",
	}
	as.NoError(as.DB.Create(tweet))

	res := as.HTML("/tweets").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), tweet.Message)
}

func (as *ActionSuite) Test_TweetsResource_Show() {
	user := as.Login()
	tweet := as.CreateTweet(user)

	result := as.HTML("/tweets/%s", tweet.ID).Get()
	as.Equal(200, result.Code)
	as.Contains(result.Body.String(), tweet.Message)
}

func (as *ActionSuite) Test_TweetsResource_NotShow_FromOtherUser() {
	anotherUser := &models.User{
		Name:       "New user",
		Provider:   "some",
		ProviderID: "129",
		Email:      nulls.NewString("new_user@test.com"),
	}
	tweet := as.CreateTweet(anotherUser)

	as.Login()

	result := as.HTML("/tweets").Get()
	as.Equal(200, result.Code)
	as.NotContains(result.Body.String(), tweet.Message)
}

func (as *ActionSuite) Test_TweetsResource_New() {
	as.Login()

	res := as.HTML("/tweets/new").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "New Tweet")
}

func (as *ActionSuite) Test_TweetsResource_Create() {
	user := as.Login()

	tweet := &models.Tweet{
		Message: "Create tweet",
	}

	res := as.HTML("/tweets").Post(tweet)
	as.Equal(302, res.Code)

	tweetFromDb := &models.Tweet{}
	as.NoError(as.DB.First(tweetFromDb))
	as.Equal(tweet.Message, tweetFromDb.Message)
	as.Equal(tweetFromDb.UserID, user.ID)
}

func (as *ActionSuite) Test_TweetsResource_Edit() {
	user := as.Login()
	tweet := as.CreateTweet(user)

	res := as.HTML("/tweets/%s/edit", tweet.ID).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Edit Tweet")
}

func (as *ActionSuite) Test_TweetsResource_Update() {
	user := as.Login()
	tweet := as.CreateTweet(user)

	oldMessage := tweet.Message
	tweet.Message = "New message"

	res := as.HTML("/tweets/%s", tweet.ID).Put(tweet)
	as.Equal(302, res.Code)

	tweetFromDb := &models.Tweet{}
	as.NoError(as.DB.First(tweetFromDb))
	as.NotEqual(oldMessage, tweetFromDb.Message)
}

func (as *ActionSuite) Test_TweetsResource_Destroy() {
	user := as.Login()
	tweet := as.CreateTweet(user)

	res := as.HTML("/tweets/%s", tweet.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/tweets", res.Location())
}
