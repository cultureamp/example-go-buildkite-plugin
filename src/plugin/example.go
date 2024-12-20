package plugin

import (
	"context"
	"fmt"

	"github.com/cultureamp/example-go-buildkite-plugin/buildkite"
)

type ExamplePlugin struct {
}

type ConfigFetcher interface {
	Fetch(config *Config) error
}

type Agent interface {
	Annotate(ctx context.Context, message string, style string, annotationContext string) error
}

func (ep ExamplePlugin) Run(ctx context.Context, fetcher ConfigFetcher, agent Agent) error {
	var config Config
	err := fetcher.Fetch(&config)
	if err != nil {
		return fmt.Errorf("plugin configuration error: %w", err)
	}
	annotation := config.Message

	buildkite.Logf("Annotating with message: %s\n", annotation)

	err = agent.Annotate(ctx, annotation, "info", "message")
	if err != nil {
		return fmt.Errorf("buildkite annotation error: %w", err)
	}

	buildkite.Log("done.")
	return nil
}
