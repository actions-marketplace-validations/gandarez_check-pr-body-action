package pr_test

import (
	"os"
	"testing"

	"github.com/gandarez/check-pr-body-action/cmd/pr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadParams_PrNumber(t *testing.T) {
	os.Setenv("INPUT_PR_NUMBER", "123")

	defer os.Unsetenv("INPUT_PR_NUMBER")

	params, err := pr.LoadParams()
	require.NoError(t, err)

	assert.Equal(t, 123, params.PrNumber)
}

func TestLoadParams_PrNumber_Err(t *testing.T) {
	os.Setenv("INPUT_PR_NUMBER", "abc")

	defer os.Unsetenv("INPUT_PR_NUMBER")

	_, err := pr.LoadParams()

	assert.Error(t, err, "failed to convert pull request number to int: abc")
}

func TestLoadParams_Contains(t *testing.T) {
	os.Setenv("INPUT_CONTAINS", "something")

	defer os.Unsetenv("INPUT_CONTAINS")

	params, err := pr.LoadParams()
	require.NoError(t, err)

	assert.Equal(t, "something", params.Contains)
}

func TestLoadParams_NotContains(t *testing.T) {
	os.Setenv("INPUT_NOT_CONTAINS", "not")

	defer os.Unsetenv("INPUT_NOT_CONTAINS")

	params, err := pr.LoadParams()
	require.NoError(t, err)

	assert.Equal(t, "not", params.NotContains)
}
