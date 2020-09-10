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


// GetSubjects to kreep
func (s *Subject) GetSubjects() string {
	rows := db.Query("select userName from subjects")
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
