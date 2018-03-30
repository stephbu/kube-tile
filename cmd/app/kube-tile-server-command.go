package app

import (
	"fmt"
	"os"

	"github.com/stephbu/kube-tiles/cmd/app/options"

	pkgServer "github.com/stephbu/kube-tiles/pkg/server"
	pkgUtilFlag "github.com/stephbu/kube-tiles/pkg/util/flag"
	pkgVersion "github.com/stephbu/kube-tiles/pkg/version"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func NewKubeTileServer() *cobra.Command {
	settings := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:  "kube-tile",
		Long: `The tile server provides interactive status for multiple kubernetes clusters`,
		Run: func(cmd *cobra.Command, args []string) {
			pkgVersion.PrintAndExitIfRequested()
			pkgUtilFlag.PrintFlags(cmd.Flags())

			stopChannel := pkgServer.SetupSignalHandler()
			if err := Run(settings, stopChannel); err != nil {
				fmt.Fprint(os.Stderr, "%v", err)
				os.Exit(1)
			}
		},
	}
	settings.Add(cmd.Flags)

	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(runOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately logs version
	glog.Infof("Version: %+v", pkgVersion.Get())

	server, err := CreateServerChain(runOptions, stopCh)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run(stopCh)
}
