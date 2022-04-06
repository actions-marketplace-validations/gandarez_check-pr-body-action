package github

import (
	"context"
	"fmt"
	"net/http"

	gh "github.com/google/go-github/v43/github"
)

//go:generate mockery --name Client --structname Mock --filename client_mock.go --inpackage
// Client is an interface for github.
type Client interface {
	PullRequestBody(number int) (string, error)
}

type github struct {
	client *gh.Client
	owner  string
	repo   string
}

// NewClient creates a new github client.
func NewClient(httpClient *http.Client, owner, repo string) Client {
	return github{
		client: gh.NewClient(httpClient),
		owner:  owner,
		repo:   repo,
	}
}

// PullRequestBody returns the body of pull request.
func (c github) PullRequestBody(number int) (string, error) {
	pr, resp, err := c.client.PullRequests.Get(context.Background(), c.owner, c.repo, number)
	if err != nil {
		return "", fmt.Errorf("failed to get pull request %d: %s", number, err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get pull request %d: with status code: %d", number, resp.StatusCode)
	}

	return pr.GetBody(), nil
}
