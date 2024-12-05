package plugin_test

import (
	"os"
	"testing"

	"github.com/cultureamp/examplego/plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailOnMissingEnvironment(t *testing.T) {
	unsetEnvironmentVariables()

	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	err := fetcher.Fetch(&config)

	expectedErr := "required key BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE missing value"
	assert.EqualError(t, err, expectedErr, "fetch should error on missing environment variable")
}

func TestFetchConfigFromEnvironment(t *testing.T) {
	unsetEnvironmentVariables()
	defer unsetEnvironmentVariables()

	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	t.Setenv("BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE", "test-message")

	err := fetcher.Fetch(&config)

	require.NoError(t, err, "fetch should not error")
	assert.Equal(t, "test-message", config.Message, "fetched message should match environment")
}

// Unsets environment variables through an all-in-one function. Extend this with additional environment variables as
// needed.
func unsetEnvironmentVariables() {
	os.Unsetenv("BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE")
}
