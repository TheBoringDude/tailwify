package cmd

import (
	"strings"

	"github.com/TheBoringDude/tailwify/v1/worker"
	"github.com/spf13/cobra"
)

// viteVue3Cmd represents the viteVue3 command
var viteVue3Cmd = &cobra.Command{
	Use:   "vite-vue3",
	Short: "Vue3 APP w/ Vite",
	Long:  `Configure and setup a Vue3 APP with Vite`,
	Run: func(cmd *cobra.Command, args []string) {
		generate := &worker.Worker{
			AppType:     "vite-vue3",
			ProjectName: strings.ToLower(projectName), // npm & yarn doesn't allow having caps in project names
			JsApp:       true,
		}

		// generate
		generate.Start(UseNPM)
	},
}

func init() {
	generateCmd.AddCommand(viteVue3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viteVue3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viteVue3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
