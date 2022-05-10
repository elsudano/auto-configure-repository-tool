package acrt

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Client object, this object contain:
type Client struct {
	// Github client used to communicate with the GitHub API.
	Client *githubv4.Client
	// User to access
	User string
	// Password of User
	Password string
	// Token of access
	Token string
	// Context for the client
	Context context.Context
	// Debug Mode
	Debug bool
}

func NewClient() (*Client, error) {
	ghc := new(Client)
	var ok bool
	ghc.User = ""
	ghc.Password = ""
	ghc.Token, ok = os.LookupEnv("GITHUB_TOKEN")
	ghc.Context = context.Background()
	ghc.Debug = false
	if ok {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghc.Token},
		)
		ghc.Client = githubv4.NewClient(
			oauth2.NewClient(ghc.Context, ts),
		)
	} else {
		log.Printf("[ACRT] Fi: client.go Fu: NewClient Obj:GitHub Client %#v\n", ghc.Client)
	}
	if ghc.Debug {
		log.SetOutput(os.Stderr)
	}
	if !ghc.Debug {
		log.SetOutput(ioutil.Discard)
	}
	return ghc, nil
}

// func NewClient(user string, pass string) (*Client, error) {
// 	ghc := new(Client)
// 	ghc.User = user
// 	ghc.Password = pass
// 	ghc.Token = ""
// 	ghc.Context = context.Background()
// 	ghc.Debug = false

// 	log.Printf("[ACRT] Fi: client.go Fu: NewClient Obj:GitHub Client %#v\n", ghc.Client)
// 	if ghc.Debug {
// 		log.SetOutput(os.Stderr)
// 	}
// 	if !ghc.Debug {
// 		log.SetOutput(ioutil.Discard)
// 	}
// 	return ghc, nil
// }

func (c *Client) debugIsEnabled() bool {
	return c.Debug
}
