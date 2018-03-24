// Brian Bright

// See page 112.
// My solution to ex 4.11

package main

import (
	"fmt"
	"log"
	//	"os"
//	"strings"
	"time"
	"github.com/bbright024/gopl_ex/ch4/github"
	"flag"
	"strings"
)

var n = flag.Int("n", 0, "ticket number")
var v = flag.Bool("v", false, "verbose printouts")
var u = flag.String("u", "", "github username")
var g = flag.Bool("g", false, "get a specific issue - requires username, repo, #")
var r = flag.String("r", "", "requested github repo")
//var e = flag.String("e", "/usr/bin/emacs", "specified editor ")
var s = flag.Bool("s", false, "search terms (-s='hello world')")
//var i = flag.Bool("i", false, "name of issue")
//var w = flag.Boool("w", false, "writing an issue")
var o = flag.String("o", "", "owner of repo" )
var c = flag.Bool("c", false, "create a new issue")
var t = flag.String("t", "", "title for new issue")
//func edit_issue(issue *github.Issue) {
	
//}

func create_issue(owner, repo, title, body, user string) {
//	user2 := []string{user}
	result, err := github.CreateIssue(owner, repo, title, body, user)
	if err != nil {
		fmt.Println("create issue error")
//		fmt.Error("Create push failed")
		log.Fatal(err)
	}
	github.PrintIssue(result, false)
}

func get_issue(owner, repo string, number int) *github.Issue {
	result, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	return result
}


func main() {
	flag.Parse()
	if *s {
		query := flag.Args()
		search(query)
	} else if *g {
		issue := get_issue(*o, *r, *n)
		github.PrintIssue(issue, true)
	} else if *c {
		fmt.Println("calling create_issue")
		body := flag.Args()
		create_issue(*o, *r, *t, strings.Join(body, " "), *u)
	}

}

func search(query []string) {
	fmt.Println(query)
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	monthLater := now.AddDate(0, -1, 0)
	yearLater := now.AddDate(-1, 0, 0)
	
	month := []*github.Issue{}
	lessYear := []*github.Issue{}
	olderYear := []*github.Issue{}
	for _, item := range result.Items {
		if monthLater.Before(item.CreatedAt) {
			month = append(month, item)
		} else if yearLater.Before(item.CreatedAt) {
			lessYear = append(lessYear, item)
		} else {
			olderYear = append(olderYear, item)
		}
	}

	fmt.Printf("\n<1 month:")
	github.PrintIssues(month, *v)
	fmt.Printf("\n<1 year:\n")
	github.PrintIssues(lessYear, *v)
	fmt.Printf("\n>1 year:\n")
	github.PrintIssues(olderYear, *v)
}

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
