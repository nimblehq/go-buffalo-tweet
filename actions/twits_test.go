package actions

import (
    "github.com/bufftwitt/models"
    "github.com/markbates/pop/nulls"
)

func (as *ActionSuite) Test_TwitsResource_List() {
    user := &models.User{
        Name: "Trung",
        Email: nulls.NewString("test@test.com"),
        Provider: "test",
        ProviderID: "123",

    }
    as.NoError(as.DB.Create(user))

    as.Session.Set("current_user_id", user.ID)

    res := as.HTML("/twits").Get()
    as.Equal(200, res.Code)
}