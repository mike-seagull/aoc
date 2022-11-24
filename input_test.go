package main

import (
	"testing"
	"time"
	"os"
)

func setup() (Input){
	token := os.Getenv("AOC_TOKEN")
	return Input{token: token}
}

func cleanup() {
	time.Sleep(10 * time.Second)
}

func TestGetInput(t *testing.T) {
	session := setup()
	input, err := session.GetInput(2020, 1)
	if input == "" || err != nil {
		t.Error(err)
	}
	t.Cleanup(cleanup)
}
func TestGetScenario(t *testing.T) {
	session := setup()
	scenario, err := session.GetScenario(2020, 1)
	if scenario == "" || err != nil {
		t.Error(err)
	}
	t.Cleanup(cleanup)
}