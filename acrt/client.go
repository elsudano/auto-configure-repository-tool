package acrt

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func NewClient(input string) (users []*gitlab.User) {
	git, err := gitlab.NewClient(input)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	users, _, err = git.Users.ListUsers(&gitlab.ListUsersOptions{})
	return
}
