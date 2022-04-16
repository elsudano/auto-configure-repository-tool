package main

import (
	"auto-configure-repository-tool/acrt"
	"flag"
	"fmt"
)

func main() {
	repo := flag.String("repo", "", "You need put the URL to the repo that you want to read")
	flag.Parse()

	if *repo != "" {
		fmt.Printf(*repo + "\n")
		acrt.ReadRepo("scm.capside.com/terraform/google-cloud/ntt-gcp-resource-policy.git")
	} else {
		flag.Usage()
	}
}
