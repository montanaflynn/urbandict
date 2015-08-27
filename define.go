package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	List       []struct {
		Defid       int    `json:"defid"`
		Word        string `json:"word"`
		Author      string `json:"author"`
		Permalink   string `json:"permalink"`
		Definition  string `json:"definition"`
		Example     string `json:"example"`
		ThumbsUp    int    `json:"thumbs_up"`
		ThumbsDown  int    `json:"thumbs_down"`
		CurrentVote string `json:"current_vote"`
	} `json:"list"`
}

func getDefinitions(w string) (data, error) {
	var d = data{}

	url := "http://api.urbandictionary.com/v0/define?term=" + w

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return d, err
	}

	req.Header.Add("user-agent", "urbandic-cli")

	printDebug("Sending request to API")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return d, err
	}

	defer res.Body.Close()
	printDebug("Getting response from API")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return d, err
	}

	printDebug("Converting JSON response to data")
	err = json.Unmarshal(body, &d)
	if err != nil {
		return d, err
	}

	if *tags == true && len(d.Tags) < 1 {
		return d, errors.New("No tags found")
	}

	printDebug("Checking if there are any definitions")
	if d.ResultType == "no_results" || len(d.List) < 1 {
		return d, errors.New("No definitions found")
	}

	return d, nil
}

func printDebug(msg string) {
	if *debug {
		fmt.Println("DEBUG: " + msg)
	}
}
