package cmd

import (
	"time"

	"github.com/80-am/kreepr/db"
)

// UNIXDATE standard
const UNIXDATE = "Mon Jan _2 15:04:05 MST 2006"

// GetFollowersPerDay since joined
func GetFollowersPerDay(s Subject) int {
	r := db.QueryRow("SELECT joinDate, followers FROM subjects WHERE userId = (?);", s.UserID)
	var date string
	var followers int
	r.Scan(&date, &followers)
	joined, _ := time.Parse(UNIXDATE, date)
	accAge := daysAgo(joined)
	fpd := followers / accAge
	return fpd
}

// GetTodaysNew followers, friends or tweets
func GetTodaysNew(data string, s Subject) int {
	q := "SELECT " + data + " FROM history WHERE userName = '" + s.UserName + "' ORDER BY date DESC LIMIT 2;"
	r := db.Query(q)
	defer r.Close()
	days := []int{}
	for r.Next() {
		var data int
		r.Scan(&data)
		days = append(days, data)
	}
	return days[0] - days[1]
}

func daysAgo(start time.Time) int {
    days := -start.YearDay()
	now := time.Now().UTC()
    for year := start.Year(); year < now.Year(); year++ {
        days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
    }
    days += now.YearDay()
    return days
}