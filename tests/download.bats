#!/usr/bin/env bats

load '/usr/local/lib/bats/load.bash'

load '../lib/download.bash'

#
# Tests for top-level docker bootstrap command. The rest of the plugin runs in Go.
#

# Uncomment the following line to debug stub failures
# export [stub_command]_STUB_DEBUG=/dev/tty
#export DOCKER_STUB_DEBUG=/dev/tty

setup() {
  export BUILDKITE_PLUGIN_EXAMPLE_GO_BUILDKITE_PLUGIN_TEST_MODE=true
}

teardown() {
    unset BUILDKITE_PLUGIN_EXAMPLE_GO_BUILDKITE_PLUGIN_TEST_MODE
    rm ./example-go-buildkite-plugin || true
}

create_script() {
cat > "$1" << EOM
set -euo pipefail

echo "executing $1:\$@"

EOM
}

@test "Downloads and runs the command for the current architecture" {

  function downloader() {
    echo "$@";
    create_script $2
  }
  export -f downloader

  run download_binary_and_run

  unset downloader

  assert_success
  assert_line --regexp "https://github.com/cultureamp/example-go-buildkite-plugin/releases/latest/download/example-go-buildkite-plugin_linux_amd64 example-go-buildkite-plugin"
  assert_line --regexp "executing example-go-buildkite-plugin"
}
