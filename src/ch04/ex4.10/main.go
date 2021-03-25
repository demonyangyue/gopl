package main

import (
	"fmt"
	"github.com/demonyangyue/gopl/src/ch04/github"
	"log"
	"time"
)

func main() {
	ageStat := make(map[string]([]string))

	result, err := github.SearchIssues([]string{"repo:golang/go"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		ageCate := calculateAgeCategory(item)
		ageStat[ageCate] =  append(ageStat[ageCate], fmt.Sprintf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title))
	}

	for k, v := range ageStat {
		fmt.Printf("%s:\n%v\n", k ,v)
	}
	
}

func calculateAgeCategory(item *github.Issue) string {

	now := time.Now()

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	switch  {
	case item.CreatedAt.Before(year) :
		return "More then one year "
	case (!item.CreatedAt.Before(year)) && item.CreatedAt.Before(month):
		return "More than one month"
	case (!item.CreatedAt.Before(month)) && item.CreatedAt.Before(day):
		return "More than one day"
	default:
		return "Less then one day"

	}
	
}
