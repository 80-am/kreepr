package main

import (
	"fmt"
	"net/url"

	"github.com/80-am/kreepr/db"
	"github.com/ChimeraCoder/anaconda"
)

var api *anaconda.TwitterApi
var c Config
var subject Subject

func getSubjectData(api *anaconda.TwitterApi) {
	s := subject.GetSubjects()
	users, err := api.GetUsersLookup(s, url.Values{})
	if err != nil {
		fmt.Println(err)
	}

	t := make([]Subject, len(users))
	for i, u := range users {
		t[i].UserID = u.Id
		t[i].Name = u.Name
		t[i].UserName = u.ScreenName
		t[i].Followers = u.FollowersCount
		t[i].Friends = u.FriendsCount
		t[i].Tweets = u.StatusesCount
		t[i].Location = u.Location
		subject.UpdateSubject(t[i])
	}
}

func main() {
	c.GetConfig()
	api := anaconda.NewTwitterApiWithCredentials(c.AccessToken, c.AccessSecret, c.Key, c.Secret)
	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	getSubjectData(api)
}
