package main

import (
	"os"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
