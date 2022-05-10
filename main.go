package main

import (
	"auto-configure-repository-tool/acrt"
	"flag"
	"fmt"
	"log"
)

func main() {
	token := flag.String("token", "", "you need put your access token in order to make actions in your repository")
	repo := flag.String("repo", "", "You need put the URL to the repo that you want to read")
	flag.Parse()

	if *repo != "" {
		fmt.Printf(*token + "\n")
		fmt.Printf(*repo + "\n")
		client, err := acrt.NewClient()
		if err != nil {
			log.Printf("[ACRT] Fi: main.go Fu: main Obj:GitHub Client %#v\n", client.Client)
		}
		client.Debug = true
		client.ShowRepo(client.ShowLogin(), *repo)
	} else {
		flag.Usage()
	}
}
