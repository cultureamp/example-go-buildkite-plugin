package main

import (
	"context"
	"os"

	"github.com/cultureamp/example-go-buildkite-plugin/buildkite"
	"github.com/cultureamp/example-go-buildkite-plugin/plugin"
)

func main() {
	ctx := context.Background()
	agent := &buildkite.Agent{}
	fetcher := plugin.EnvironmentConfigFetcher{}
	examplePlugin := plugin.ExamplePlugin{}

	err := examplePlugin.Run(ctx, fetcher, agent)

	if err != nil {
		buildkite.LogFailuref("plugin execution failed: %s\n", err.Error())
		os.Exit(1)
	}
}
