package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Juniper/contrail-go-api"
)

type createOptions struct {
	parent   string
	parentId string
}

var (
	createOpts createOptions
)

func create(client *contrail.Client, flagSet *flag.FlagSet) {
	fmt.Printf("%d: %s\n", flagSet.NArg(), flagSet.Args())
	if flagSet.NArg() < 2 {
		flagSet.Usage()
		os.Exit(2)
	}

	resource := flagSet.Args()[0]
	//	name := flagSet.Args()[1]

	var parentId string
	if len(createOpts.parent) > 0 {
		obj := contrail.TypeToObject(resource)
		var err error
		parentId, err = client.UuidByName(obj.GetDefaultParentType(), createOpts.parent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			flagSet.Usage()
			os.Exit(2)
		}
	} else if len(createOpts.parentId) > 0 {
		parentId = createOpts.parentId
	} else {
		fmt.Fprintf(os.Stderr, "Error: parent name or uuid must be specified\n")
		flagSet.Usage()
		os.Exit(2)
	}
	fmt.Printf("Uuid: %s\n", parentId)
}

func init() {
	createFlags := flag.NewFlagSet("create", flag.ExitOnError)
	initCommonProjectOptions(createFlags)
	createFlags.StringVar(&createOpts.parent,
		"parent", "", "Parent resource")
	createFlags.StringVar(&createOpts.parentId,
		"parent-id", "", "Uuid of parent resource")
	createFlags.Usage = defaultUsage("create", "resource-type name", createFlags)
	RegisterCliCommand("create", createFlags, create)
}
