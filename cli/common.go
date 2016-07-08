package main

import (
	"flag"
	"fmt"
	"os"
)

func defaultUsage(name string, arguments string, flagSet *flag.FlagSet) func() {
	return func() {
		flagSet.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s %s %s [OPTIONS]\n", os.Args[0], name, arguments)
	}
}

func initCommonProjectOptions(flagSet *flag.FlagSet) {
	defaultProject := os.Getenv("OS_TENANT_NAME")
	if len(defaultProject) == 0 {
		defaultProject = "admin"
	}

	flagSet.StringVar(&policyCommonOpts.project, "project", defaultProject,
		"Project name (Env: OS_TENANT_NAME)")
	flagSet.StringVar(&policyCommonOpts.projectId, "project-id",
		os.Getenv("OS_TENANT_ID"), "Project id (Env: OS_TENANT_ID)")
}
