package acrt

import (
	"context"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
	"net/http"
)

func NewClient(token string, ctx context.Context) (ghc *github.Client) {
	cwt := Authentication(token, ctx)
	ghc = github.NewClient(cwt)
	return
}

func Authentication(token string, ctx context.Context) (cwt *http.Client) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	cwt = oauth2.NewClient(ctx, ts)
	return
}
