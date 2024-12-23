//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
	"github.com/myorg/provider-artfactory/e2e"
)

func SetupE2E() error {
	err := e2e.EnsureKindCluster("kind")

	if err != nil {
		return err
	}

	err = e2e.EnsureArtifactory()

	if err != nil {
		return err
	}

	return e2e.UpdateCredentials()
}

func TestE2E() error {
	// See: https://onsi.github.io/ginkgo/#recommended-continuous-integration-configuration
	return sh.RunV("ginkgo", "-r", "-v", "e2e",
		"--fail-on-pending",
		"--fail-on-empty",
		"--randomize-all",
		"--randomize-suites",
		"--keep-going",
		"--procs=4",
	)
}
