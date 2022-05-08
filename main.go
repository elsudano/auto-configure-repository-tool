package main

import (
	"auto-configure-repository-tool/acrt"
	"context"
	"flag"
	"fmt"
	"log"
)

func main() {
	token := flag.String("token", "", "you need put your access token in order to make actions in your repository")
	repo := flag.String("repo", "", "You need put the URL to the repo that you want to read")
	flag.Parse()

	if *token != "" {
		fmt.Printf(*token + "\n")
		fmt.Printf(*repo + "\n")
		context := context.Background()
		client := acrt.NewClient(*token, context)
		// list all repositories for the authenticated user
		repos, resp, err := client.Repositories.List(context, "", nil)
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
		fmt.Printf("Repos: %#v", repos)
		// Rate.Limit should most likely be 5000 when authorized.
		log.Printf("Rate: %#v\n", resp.Rate)
	} else {
		flag.Usage()
	}
}
