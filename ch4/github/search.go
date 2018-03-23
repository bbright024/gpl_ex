// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"strconv"
)

func GetIssue(owner, repo string, number int) (*Issue, error){
	owner = url.QueryEscape(owner)
	repo = url.QueryEscape(repo)
	q := owner + "/" + repo + "/issues/"
	req, err := http.NewRequest("GET", IssueURL + q + strconv.Itoa(number), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println("full query:")
	fmt.Printf("%q \t\n", IssuesURL + "?q=" + q)

	req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
	   return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}


/*
func send_req(req *http.Request) (*http.Response, error) {

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	fmt.Println(req)

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	// this makes the code non-modular: having to have a specific struct to insert results
	// before closing the body.  maybe that's what "Defer" helps with? otherwise would
	// need to pass in a generic and assign the json decoder to that parameter.
	// basically i need to learn more go to be able to do this, so commenting out for now.
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil	
}
*/
