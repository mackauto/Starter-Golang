package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/data_structure/github"
	"log"
)

func main() {
	//result, err := github.SearchIssues(os.Args[1:])
	args := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
