package main

import (
	"fmt"
	"../db"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var c conf
	c.getConf()
	api := anaconda.NewTwitterApiWithCredentials(c.AccessToken, c.AccessSecret, c.Key, c.Secret)

	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		panic(err.Error())
	}

	defer database.Close()

	searchResult, err := api.GetSearch("TODO", nil)

	if err != nil {
		panic(err.Error())
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
