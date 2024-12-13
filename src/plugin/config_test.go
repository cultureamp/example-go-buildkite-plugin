package plugin_test

import (
	"os"
	"testing"

	"github.com/cultureamp/examplego/plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailOnMissingEnvironment(t *testing.T) {
	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	unsetEnv(t, "BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE")

	err := fetcher.Fetch(&config)

	expectedErr := "required key BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE missing value"
	assert.EqualError(t, err, expectedErr, "fetch should error on missing environment variable")
}

func TestFetchConfigFromEnvironment(t *testing.T) {
	var config plugin.Config
	fetcher := plugin.EnvironmentConfigFetcher{}

	t.Setenv("BUILDKITE_PLUGIN_EXAMPLE_GO_MESSAGE", "test-message")

	err := fetcher.Fetch(&config)

	require.NoError(t, err, "fetch should not error")
	assert.Equal(t, "test-message", config.Message, "fetched message should match environment")
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()

	// ensure state is restored correctly
	currValue, exists := os.LookupEnv(key)
	t.Cleanup(func() {
		if exists {
			os.Setenv(key, currValue)
		} else {
			os.Unsetenv(key)
		}
	})

	// clear the value
	os.Unsetenv(key)
}
