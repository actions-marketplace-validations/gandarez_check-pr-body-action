package pr_test

import (
	"testing"

	"github.com/gandarez/check-pr-body-action/internal/github"
	"github.com/gandarez/check-pr-body-action/internal/pr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckBody_Empty(t *testing.T) {
	gh := new(github.Mock)
	gh.On("PullRequestBody", mock.AnythingOfType("int")).Return("", nil)

	pr := pr.New(gh)

	err := pr.CheckBody(123, "something", "")

	assert.Error(t, err, "pull request body is empty")
}

func TestCheckBody_Contains(t *testing.T) {
	gh := new(github.Mock)
	gh.On("PullRequestBody", mock.AnythingOfType("int")).
		Return("Changelog:\n- some text 1\n-some text 2", nil)

	pr := pr.New(gh)

	err := pr.CheckBody(123, "Changelog:", "")
	require.NoError(t, err)
}

func TestCheckBody_Contains_Err(t *testing.T) {
	gh := new(github.Mock)
	gh.On("PullRequestBody", mock.AnythingOfType("int")).
		Return("Changelog:\n- some text 1\n-some text 2", nil)

	pr := pr.New(gh)

	err := pr.CheckBody(123, "something", "")

	assert.Error(t, err, "pull request body does not contain: something")
}

func TestCheckBody_NotContains(t *testing.T) {
	gh := new(github.Mock)
	gh.On("PullRequestBody", mock.AnythingOfType("int")).
		Return("Changelog:\n- some text 1\n-some text 2", nil)

	pr := pr.New(gh)

	err := pr.CheckBody(123, "", "some text 3")
	require.NoError(t, err)
}

func TestCheckBody_NotContains_Err(t *testing.T) {
	gh := new(github.Mock)
	gh.On("PullRequestBody", mock.AnythingOfType("int")).
		Return("Changelog:\n- some text 1\n-some text 2", nil)

	pr := pr.New(gh)

	err := pr.CheckBody(123, "", "some text 1")

	assert.Error(t, err, "pull request body contains: something")
}
