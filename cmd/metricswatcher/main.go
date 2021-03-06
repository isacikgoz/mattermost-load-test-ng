package main

import (
	"os"

	"github.com/mattermost/mattermost-load-test-ng/cmd/metricswatcher/config"
	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:    "metricswatcher",
		RunE:   runMetricsWatcher,
		PreRun: config.SetupMetricsCheck,
	}

	persistentFlags := rootCmd.PersistentFlags()
	persistentFlags.StringP("config", "c", "", "path to the configuration file to use")

	if err := rootCmd.Execute(); err != nil {
		mlog.Error(err.Error())
		os.Exit(1)
	}
}

func runMetricsWatcher(cmd *cobra.Command, args []string) error {
	configuration, err := config.GetMetricsCheckConfig()

	if err != nil {
		return err
	}

	errChan := make(chan error, 1)
	defer close(errChan)

	go healthcheck(errChan, configuration)
	go checkMetrics(errChan, configuration)

	err = <-errChan

	return err
}
