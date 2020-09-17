package cmd

import (
	"github.com/80-am/kreepr/db"
)

// Tweet to kreep
type Tweet struct {
	ID       int64
	Text     string
	Created  string
	Likes    int
	ReTweets int
	ReplyTo  string
}

// GetLastTweetID of subject
func (t *Tweet) GetLastTweetID(s Subject) string {
	r := db.QueryRow("SELECT id FROM tweets WHERE userId =(SELECT MAX(?) FROM tweets);", s.UserID)
	var id string
	r.Scan(&id)
	return id
}

// GetNumberOfTimesTweetedAbout a keyword defined
func (t *Tweet) GetNumberOfTimesTweetedAbout(about string, s Subject) int64 {
	q := "SELECT id FROM tweets WHERE userName = '" + s.UserName + "' AND text LIKE '%" + about + "%';"
	r := db.Query(q)
	defer r.Close()
	var times int64
	for r.Next() {
		var id int64
		r.Scan(&id)
		if id != 0 {
			times++
		}
	}
	return times
}

// SubjectTweetedAbout a keyword defined
func (t *Tweet) SubjectTweetedAbout(about string, s Subject) {
	q := "SELECT id, text, created, likes, reTweets, replyTo FROM tweets WHERE userName = '" + s.UserName + "' AND text LIKE '%" + about + "%';"
	r := db.Query(q)
	defer r.Close()
	tweets := []Tweet{}
	for r.Next() {
		var id int64
		var text string
		var created string
		var likes int
		var reTweets int
		var replyTo string
		var tweet Tweet
		r.Scan(&id, &text, &created, &likes, &reTweets, &replyTo)
		tweet.ID = id
		tweet.Text = text
		tweet.Created = created
		tweet.Likes = likes
		tweet.ReTweets = reTweets
		tweet.ReplyTo = replyTo
		tweets = append(tweets, tweet)
	}
}

// UpdateTweets of subject
func (t *Tweet) UpdateTweets(s Subject, tw Tweet) {
	stmt := db.Prepare("INSERT INTO tweets(id, userId, userName, text, created, likes, reTweets, replyTo) VALUES(?, ?, ?, ?, ?, ?, ?, ?);")
	stmt.Exec(tw.ID, s.UserID, s.UserName, tw.Text, tw.Created, tw.Likes, tw.ReTweets, tw.ReplyTo)
}