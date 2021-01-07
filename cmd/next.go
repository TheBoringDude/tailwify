package cmd

import (
	"github.com/TheBoringDude/tailwify/worker"
	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Next.js APP",
	Long:  `Configure and setup a Next.js APP`,
	Run: func(cmd *cobra.Command, args []string) {
		generate := &worker.Worker{
			AppType:     "next",
			ProjectName: projectName,
			JsApp:       true,
		}

		// generate
		generate.Start()
	},
}

func init() {
	generateCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
