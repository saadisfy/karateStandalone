package main

import (
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/envs"
	"github.com/kubeshop/testkube/pkg/executor/agent"
	"github.com/kubeshop/testkube/pkg/executor/output"

	"github.com/saadisfy/karateStandalone/pkg/runner"
)

func main() {
	contextBackground := context.Background()
	params, err := envs.LoadTestkubeVariables()
	if err != nil {
		output.PrintError(os.Stderr, errors.Errorf("could not initialize Karate Executor environment variables: %v", err))
		os.Exit(1)
	}
	karateRunner, err := runner.NewKarateRunner(contextBackground, os.Getenv("DEPENDENCY_MANAGER"), params)
	if err != nil {
		output.PrintError(os.Stderr, errors.Wrap(err, "Error instantiating Karate Executor"))
		os.Exit(1)
	}
	agent.Run(contextBackground, karateRunner, os.Args)
}
