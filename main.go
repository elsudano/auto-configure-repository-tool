package main

import (
	"auto-configure-repository-tool/acrt"
	"flag"
	"fmt"
)

func main() {
	token := flag.String("token", "", "you need put your access token in order to make actions in your repository")
	repo := flag.String("repo", "", "You need put the URL to the repo that you want to read")
	flag.Parse()

	if *repo != "" {
		fmt.Printf(*repo + "\n")
		acrt.NewClient(*token)
	} else {
		flag.Usage()
	}
}
