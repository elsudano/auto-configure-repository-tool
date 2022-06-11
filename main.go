package main

import (
	"auto-configure-repository-tool/acrt"
	"flag"
	"fmt"
	"log"
)

func main() {
	repo := flag.String("repo", "", "You need put the URL to the repo that you want to read")
	flag.Parse()

	if *repo != "" {
		fmt.Printf(*repo + "\n")
		client, err := acrt.NewClient()
		if err != nil {
			log.Printf("[ACRT] Fi: main.go Fu: main Obj:GitHub Client %#v\n", client.Client)
		}
		client.Debug = true
		login := client.ShowLogin()
		fmt.Printf("Login: %v\n", login)
		client.Last10Repositories(login)
		client.ShowRepo(login, *repo)
	} else {
		flag.Usage()
	}
}
