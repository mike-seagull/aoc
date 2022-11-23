package main

import (
	"time"
	"net/http"
	"net/http/cookiejar"
	"fmt"
	"io"
	"errors"
	"strings"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

type Input struct {
	token string	
}

func (i Input) GetInput(year int, day int) (string, error) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
        Jar: jar,
    }
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, _ := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
        Name:   "session",
        Value:  i.token,
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

func (i Input) GetInputLines(year int, day int) []string {
	resp, err := i.GetInput(year, day)
	if err != nil {
		return []string{}
	} else {
		return strings.Split(resp, "\n")
	}
}

func (i Input) GetTodaysInput() string {
	t := time.Now()
	resp, err := i.GetInput(t.Year(), t.Day())
	if err != nil {
		// probably a day early
		resp, err = i.GetInput(t.Year(), t.Day() - 1)
	}
	return resp
}

func (i Input) GetTodaysInputLines() []string {
	return strings.Split(i.GetTodaysInput(), "\n")
}

func (i Input) GetScenario(year int, day int) (string, error) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
        Jar: jar,
    }
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, _ := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
        Name:   "session",
        Value:  i.token,
        MaxAge: 300,
    }
	req.AddCookie(cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Status code = %d", resp.StatusCode))
	}
	html := string(bodyBytes)
	converter := md.NewConverter("", true, nil)	
	converter.Use(plugin.GitHubFlavored())
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", err
	}
	return markdown, nil
}

func (i Input) GetTodaysScenario() string {
	t := time.Now()
	resp, err := i.GetScenario(t.Year(), t.Day())
	if err != nil {
		// probably a day early
		resp, err = i.GetScenario(t.Year(), t.Day() - 1)
		if err != nil {
			return ""
		}
	}
	return resp
}
