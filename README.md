# Example Go Buildkite Plugin

An example Buildkite plugin written in Go, demonstrating on-demand download of the
public release files from Github and execution.

To demonstrate interaction with the Buildkite agent, the example includes code
to annotate the build with a message.

This example is reverse engineered from [@cultureamp/ecr-scan-results-buildkite-plugin](https://github.com/cultureamp/ecr-scan-results-buildkite-plugin).

## Example

Add the following lines to your `pipeline.yml`:

```yml
steps:
  - plugins:
      - cultureamp/example-go#v0.1.0:
          message: "This is the message that will be annotated!"
```

## Configuration

### `message` (Required, string)

The message to annotate onto the build.

## Local development

Ensure devbox is setup as per [Local Dev Environments (LDEs) - Getting Started](https://cultureamp.atlassian.net/wiki/spaces/DE/pages/3342434338/Devbox+setup).

Install dependencies:

    devbox run setup

Fast feedback (lint + tests):

    devbox run check

Full sweep (module tidiness check, lint, and race/coverage tests):

    devbox run verify

## Releasing

Push a version tag to trigger new release via [Github Actions workflow](./.github/workflows/release.yaml).

```
git tag v0.1.0
git push --tags
```