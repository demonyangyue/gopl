package main

import (
	"fmt"
	"github.com/demonyangyue/gopl/src/ch04/github"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.CacheURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type IssueCache struct {
	Issues []github.Issue
	IdOfIssue map[int]github.Issue
}


func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr,
`usage: ex4.14 <owner> <repo>
  e.g. ex4.14 torbiak gopl`)
		os.Exit(1)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	issueCache, err := NewIssueCache(owner, repo)
	if err != nil {
		log.Fatal("failed to create issues cache: " + err.Error())
	}

	http.Handle("/", issueCache)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewIssueCache(owner string, repo string) (ic IssueCache, err error) {

	issues, err := github.GetIssues(owner, repo)

	if err != nil {
		return
	}

	ic.Issues = issues
	ic.IdOfIssue = make(map[int]github.Issue)

	for _, issue := range issues {
		ic.IdOfIssue[issue.Number] = issue
	}
	return

}

func (ic IssueCache) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := strings.Split(request.URL.Path, "/")
	if len(path) < 1 {
		issueListTemplate.Execute(writer, ic)
		return
	}

	num, err := strconv.Atoi(path[len(path) - 1])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid issue id"))
		return
	}

	issue, ok := ic.IdOfIssue[num]

	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("issue not found"))
		return
	}

	issueTemplate.Execute(writer, issue)
}

