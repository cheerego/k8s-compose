package cmd

import (
	"github.com/cheerego/k8s-compose/cmd/compose"
	"github.com/spf13/cobra"
)

func NewRootCommand(compose *compose.Compose) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "k8s-compose",
		Short: "k8s-compose",
		Long: `k8s-compose is a tool for defining and running multi-container Docker applications. 
With KCompose, you use a YAML file to configure your application’s services. 
Then, with a single command, you create and start all the services from your configuration. 
To learn more about all the features of KCompose`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.InitDefaultHelpCmd()
	rootCmd.InitDefaultVersionFlag()

	rootCmd.AddCommand(NewVersionCommand(compose))
	return rootCmd
}
