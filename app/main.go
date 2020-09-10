package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/80-am/kreepr/db"
	"github.com/ChimeraCoder/anaconda"
)

var api *anaconda.TwitterApi
var c Config
var subject Subject
var addSub string

func init() {
	flag.StringVar(&addSub, "addSubject", "", "Add new subject")
	flag.Parse()
}

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

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
	})
    return found
}

func isEmptySubjectDb() bool {
	s := subject.GetSubjects()
	if s != "" {
		return false
	}
	return true
}

func main() {
	c.GetConfig()
	api := anaconda.NewTwitterApiWithCredentials(c.AccessToken, c.AccessSecret, c.Key, c.Secret)
	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	if isEmptySubjectDb() && !isFlagPassed("addSubject") {
		fmt.Print("Add a subject to kreep: ")
		fmt.Scan(&addSub) 
		subject.AddSubject(addSub)
	} else if isFlagPassed("addSubject") {
		subject.AddSubject(addSub)
	}
	getSubjectData(api)
}
