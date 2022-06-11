package acrt

import (
	"fmt"
	"log"

	"github.com/shurcooL/githubv4"
)

func (c *Client) ShowLogin() string {
	var query struct {
		Viewer struct {
			Login     string
			CreatedAt githubv4.DateTime
		}
		RateLimit struct {
			Limit     githubv4.Int
			Cost      githubv4.Int
			Remaining githubv4.Int
			ResetAt   githubv4.DateTime
		}
	}
	err := c.Client.Query(c.Context, &query, nil)
	if err != nil {
		log.Printf("[ACRT] Fi: acrt.go Fu: ShowLogin Obj:Error %s\n", err)
	} else {
		fmt.Println("      Login:", query.Viewer.Login)
		fmt.Println("  CreatedAt:", query.Viewer.CreatedAt)
		if c.debugIsEnabled() {
			fmt.Println(" Rate Limit:", query.RateLimit.Limit)
			fmt.Println("       Cost:", query.RateLimit.Cost)
			fmt.Println("  Remaining:", query.RateLimit.Remaining)
			fmt.Println("   Reset At:", query.RateLimit.ResetAt)
		}
	}
	return query.Viewer.Login
}

func (c *Client) Last10Repositories(owner string) {
	// type myRepository struct {
	// 	Name   githubv4.String
	// 	sshURL githubv4.String
	// }

	var query struct {
		User struct {
			Id           githubv4.ID
			Repositories struct {
				// totalCount githubv4.Int
				// Nodes    []myRepository
				// PageInfo struct {
				// 	EndCursor   githubv4.String
				// 	HasNextPage bool
				// }
			} `graphql:"repositories(first: 10, orderBy: {field: CREATED_AT, direction: DESC}, after: $RepositoriesCursor)"`
		} `graphql:"user(login: $owner)"`
		RateLimit struct {
			Limit     githubv4.Int
			Cost      githubv4.Int
			Remaining githubv4.Int
			ResetAt   githubv4.DateTime
		}
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		// "RepositoriesCursor": (*githubv4.String)(nil),
	}
	// var allRepositories []myRepository
	// for {
	err := c.Client.Query(c.Context, &query, variables)
	if err != nil {
		log.Printf("[ACRT] Fi: acrt.go Fu: Last10Repositories Obj:Error %s\n", err)
	} else {
		fmt.Println("    Id:", query.User.Id)
		// fmt.Println(" Count:", query.User.Repositories.totalCount)
		// fmt.Println(" Repo name:", query.User.Repositories.nodes.name)
		// fmt.Println("    sshURL:", query.User.Repositories.nodes.sshUrl)
		if c.debugIsEnabled() {
			fmt.Println(" Rate Limit:", query.RateLimit.Limit)
			fmt.Println("       Cost:", query.RateLimit.Cost)
			fmt.Println("  Remaining:", query.RateLimit.Remaining)
			fmt.Println("   Reset At:", query.RateLimit.ResetAt)
		}
	}
	// 	allRepositories = append(allRepositories, query.User.Repositories.Nodes...)
	// 	if !query.User.Repositories.PageInfo.HasNextPage {
	// 		break
	// 	}
	//   variables["RepositoriesCursor"] = githubv4.NewString(query.User.Repositories.PageInfo.EndCursor)
	// }
}

func (c *Client) ShowRepo(owner string, repo string) {
	var query struct {
		Repository struct {
			Id          githubv4.ID
			Description githubv4.String
			DiskUsage   githubv4.Int
		} `graphql:"repository(owner: $owner, name: $repo)"`
		RateLimit struct {
			Limit     githubv4.Int
			Cost      githubv4.Int
			Remaining githubv4.Int
			ResetAt   githubv4.DateTime
		}
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
	}
	err := c.Client.Query(c.Context, &query, variables)
	if err != nil {
		log.Printf("[ACRT] Fi: acrt.go Fu: ShowRepo Obj:Error %s\n", err)
	} else {
		fmt.Println("    Repo ID:", query.Repository.Id)
		fmt.Println("Description:", query.Repository.Description)
		fmt.Println(" Disk Usage:", query.Repository.DiskUsage)
		if c.debugIsEnabled() {
			fmt.Println(" Rate Limit:", query.RateLimit.Limit)
			fmt.Println("       Cost:", query.RateLimit.Cost)
			fmt.Println("  Remaining:", query.RateLimit.Remaining)
			fmt.Println("   Reset At:", query.RateLimit.ResetAt)
		}
	}
}
