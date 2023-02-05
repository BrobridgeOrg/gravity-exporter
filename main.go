package main

import (
	"os"

	"github.com/BrobridgeOrg/gravity-exporter/pkg/configs"
	"github.com/BrobridgeOrg/gravity-exporter/pkg/connector"
	"github.com/BrobridgeOrg/gravity-exporter/pkg/exporter"
	"github.com/BrobridgeOrg/gravity-exporter/pkg/logger"
	"github.com/spf13/cobra"

	"go.uber.org/fx"
)

var config *configs.Config

var rootCmd = &cobra.Command{
	Use:   "gravity-exporter",
	Short: "Prometheus exporter for Gravity",
	Long:  "The gravity exporter is a component that exports Gravity server metries to Prometheus for monitoring.",
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := run(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	config = configs.GetConfig()
}

func main() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {

	config.SetConfigs(map[string]interface{}{})

	fx.New(
		fx.Supply(config),
		fx.Provide(
			logger.GetLogger,
			connector.New,
		),
		fx.Invoke(exporter.New),
		fx.NopLogger,
	).Run()

	return nil
}
