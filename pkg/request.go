package pkg

import (
	"time"
	"net/http"
	"net/http/cookiejar"
	"fmt"
	"io"
	"errors"
)

var Token string
var URL string = "https://adventofcode.com"
var Year int
// var Month int
var Day int

func init() {
    t := time.Now()
	Year = t.Year()
	// Month = t.Month()
	Day = t.Day()
}

func GetTodaysInput() string {
	resp, err := GetInput(Year, Day)
	if err != nil {
		// probably a day early
		resp, err = GetInput(Year, Day - 1)
	}
	return resp
}

func GetInput(year int, day int) (string, error) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
        Jar: jar,
    }
	url := fmt.Sprintf("%s/%d/day/%d/input", URL, year, day)
	req, _ := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
        Name:   "session",
        Value:  Token,
        MaxAge: 300,
    }
	req.AddCookie(cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Status code = %d", resp.StatusCode))
	}
	return string(bodyBytes), nil
}