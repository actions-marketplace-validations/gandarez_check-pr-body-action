package pr

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gandarez/check-pr-body-action/internal/github"
)

type Pr struct {
	client github.Client
}

func New(c github.Client) Pr {
	return Pr{
		client: c,
	}
}

func (pr Pr) CheckBody(number int, contains, notContains string) error {
	body, err := pr.client.PullRequestBody(number)
	if err != nil {
		return err
	}

	if body == "" {
		return errors.New("pull request body is empty")
	}

	if contains != "" && !strings.Contains(body, contains) {
		return fmt.Errorf("pull request body does not contain: %s", contains)
	}

	if notContains != "" && strings.Contains(body, notContains) {
		return fmt.Errorf("pull request body contains: %s", notContains)
	}

	return nil
}
