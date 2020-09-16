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

// UpdateTweets of subject
func (t *Tweet) UpdateTweets(s Subject, tw Tweet) {
	stmt := db.Prepare("INSERT INTO tweets(id, userId, userName, text, created, likes, reTweets, replyTo) VALUES(?, ?, ?, ?, ?, ?, ?, ?);")
	stmt.Exec(tw.ID, s.UserID, s.UserName, tw.Text, tw.Created, tw.Likes, tw.ReTweets, tw.ReplyTo)
}

// GetLastTweetID of subject
func (t *Tweet) GetLastTweetID(s Subject) string {
	r := db.QueryRow("SELECT id FROM tweets WHERE userId =(SELECT MAX(?) FROM tweets);", s.UserID)
	var id string
	r.Scan(&id)
	return id
}
