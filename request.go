package main

import (
	"time"
	"net/http"
	"net/http/cookiejar"
	"fmt"
	"io"
	"errors"
	"strings"
)

type Data struct {
	token string	
}

func (d Data) GetInput(year int, day int) (string, error) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
        Jar: jar,
    }
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, _ := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
        Name:   "session",
        Value:  d.token,
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

func (d Data) GetTodaysInput() string {
	t := time.Now()
	resp, err := d.GetInput(t.Year(), t.Day())
	if err != nil {
		// probably a day early
		resp, err = d.GetInput(t.Year(), t.Day() - 1)
	}
	return resp
}

func (d Data) GetTodaysInputLines() []string {
	return strings.Split(d.GetTodaysInput(), "\n")
}

