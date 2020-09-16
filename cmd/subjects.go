package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/80-am/kreepr/db"
)

// Subject to kreep
type Subject struct {
	UserID    int64
	Name      string
	UserName  string
	Followers int
	Friends   int
	Tweets    int64
	JoinDate  string
	Location  string
}

// AddSubject to insert into followers table
func (s *Subject) AddSubject(new string) {
	if !s.isNewSubject(new) {
		stmt := db.Prepare("INSERT INTO subjects(userName) VALUES(?)")
		stmt.Exec(new)
	}
}

// DropSubject drops all history and stops kreeping subject
func (s *Subject) DropSubject(user string) {
	stmt := db.Prepare("DELETE s, h, t from subjects s INNER JOIN history h ON s.userId = h.userId INNER JOIN tweets t ON s.userId = t.userId WHERE s.userName = (?)")
	stmt.Exec(user)
}

// GetSubjects to kreep
func (s *Subject) GetSubjects() string {
	rows := db.Query("SELECT userName FROM subjects;")
	var subjects strings.Builder

	for rows.Next() {
		var userName string
        if err := rows.Scan(&userName); err != nil {
                log.Fatal(err)
		}
		subjects.WriteString(userName + ", ")
	}
	return subjects.String()
}

// UpdateHistory with cron once per day to keep the history
func updateHistory(su Subject) {
	today := time.Now().UTC().Format("2006-01-02")
	r := db.QueryRow("SELECT date FROM history WHERE date = '" + today + "' AND userName = (?);", su.UserName)
	var date string
	r.Scan(&date)
	if date == "" {
		stmt := db.Prepare("INSERT INTO history(date, userId, userName, followers, friends, tweets) VALUES(?, ?, ?, ?, ?, ?);")
		stmt.Exec(today, su.UserID, su.UserName, su.Followers, su.Friends, su.Tweets)
	}
}

// UpdateSubject data
func (s *Subject) UpdateSubject(su Subject, dailyJob bool) {
	stmt := db.Prepare("UPDATE subjects SET userId = (?), name = (?), followers = (?), friends = (?)," +
	"tweets = (?), joinDate = (?), location = (?) WHERE userName = (?);")
	stmt.Exec(su.UserID, su.Name, su.Followers, su.Friends, su.Tweets, su.JoinDate, su.Location, su.UserName)
	if (dailyJob) {
		updateHistory(su)
	}
}

func (s *Subject)isNewSubject(userName string) bool {
	stmt := db.Prepare("SELECT userName FROM subjects where userName = (?);")
	rows, _ := stmt.Query(userName)
	cols, _ := rows.Columns()
	rawResult := make([][]byte, len(cols))
    dest := make([]interface{}, len(cols))
    for i := range rawResult {
        dest[i] = &rawResult[i]
    }
    for rows.Next() {
		if rawResult != nil {
			return true
		}
		return false
	}
	return false
}
