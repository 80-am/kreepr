package main

import (
	"log"
	"strings"

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
	Location  string
}

// AddSubject to insert into followers table
func (s *Subject) AddSubject(new string) {
	if !s.isNewSubject(new) {
		stmt := db.Prepare("INSERT INTO subjects(userName) VALUES(?)")
		stmt.Exec(new)
	}
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

// UpdateSubject data
func (s *Subject) UpdateSubject(su Subject) {
	stmt := db.Prepare("UPDATE subjects SET userId = (?), name = (?), followers = (?), friends = (?)," + 
	"tweets = (?), location = (?) WHERE userName = (?);")
	stmt.Exec(su.UserID, su.Name, su.Followers, su.Friends, su.Tweets, su.Location, su.UserName)
}