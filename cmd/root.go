package cmd

import (

	"github.com/spf13/cobra"

	L "GeeksCli/lib"
)

var rootCmd = &cobra.Command{
	Use:   "GeeksCli",
	Short: "GeeksCli gestión de tus entornos.",
	Long: `Una forma rápida de gestionar tus entornos.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		L.Info()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init(){
	rootCmd.AddCommand(initCmd)
}
