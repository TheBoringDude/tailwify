package cmd

import (
	"strings"

	"github.com/ootiq/tailwify/worker"
	"github.com/spf13/cobra"
)

// reactCmd represents the react command
var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "React APP",
	Long:  `Configure and setup a React APP`,
	Run: func(cmd *cobra.Command, args []string) {
		generate := &worker.Worker{
			AppType:     "react",
			ProjectName: strings.ToLower(projectName), // npm & yarn doesn't allow having caps in project names
			JsApp:       true,
		}

		// generate
		generate.Start(UseNPM)
	},
}

func init() {
	generateCmd.AddCommand(reactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
