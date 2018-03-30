package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/stephbu/kube-tiles/cmd/app"
	pkgUtilFlag "github.com/stephbu/kube-tiles/pkg/util/flag"
)

func main() {

	// Seed RNG
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewKubeTileServer()

	pflag.CommandLine.SetNormalizeFunc(pkgUtilFlag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	// start server
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
