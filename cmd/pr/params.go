package pr

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gandarez/check-pr-body-action/internal/actions"
	"github.com/gandarez/check-pr-body-action/internal/github"

	"golang.org/x/oauth2"
)

var tokenRegex = regexp.MustCompile(`[0-9a-fA-F]{40}`) // nolint

// Params contains pr command parameters.
type Params struct {
	Client      github.Client
	Contains    string
	NotContains string
	PrNumber    int
}

// LoadParams loads pr config params.
func LoadParams() (Params, error) {
	var ghToken string

	if ghTokenStr := os.Getenv("GITHUB_TOKEN"); ghTokenStr != "" {
		if !tokenRegex.MatchString(ghTokenStr) {
			return Params{}, fmt.Errorf("invalid github_token format: %s", ghTokenStr)
		}

		ghToken = ghTokenStr
	}

	var (
		owner      string
		repository string
	)

	if githubRepositoryStr := os.Getenv("GITHUB_REPOSITORY"); githubRepositoryStr != "" {
		splitted := strings.SplitN(githubRepositoryStr, "/", 2)

		owner = splitted[0]
		repository = splitted[1]
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	ghClient := github.NewClient(tc, owner, repository)

	var prNumber int

	if prNumberStr := actions.GetInput("pr_number"); prNumberStr != "" {
		parsed, err := strconv.Atoi(prNumberStr)
		if err != nil {
			return Params{}, fmt.Errorf("failed to convert pull request number to int: %s", prNumberStr)
		}

		prNumber = parsed
	}

	return Params{
		Client:      ghClient,
		Contains:    actions.GetInput("contains"),
		NotContains: actions.GetInput("not_contains"),
		PrNumber:    prNumber,
	}, nil
}
