package plugin

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Message string `required:"true"`
}

type EnvironmentConfigFetcher struct {
}

// TODO: Replace "EXAMPLE_GO" with the name of your plugin, also in upper case
const pluginEnvironmentPrefix = "BUILDKITE_PLUGIN_EXAMPLE_GO"

func (f EnvironmentConfigFetcher) Fetch(config *Config) error {
	return envconfig.Process(pluginEnvironmentPrefix, config)
}
