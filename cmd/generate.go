package cmd

import (
	"github.com/ootiq/tailwify/v1/cmd/internal/hooks"
	"github.com/spf13/cobra"
)

var projectName string

// UseNPM gets if the user wants to use npm not yarn.
var UseNPM bool

// NodeApps is the main required apps for installing.
var NodeApps = map[string]nodeAppCommand{
	"node": {
		Command: "node",
	},
	"npm": {
		Command: "npm",
	},
	"yarn": {
		Command: "yarn",
	},
}

type nodeAppCommand struct {
	Command string
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate and setup a starter template",
	Long: `It will configure a simple and starter app for you 
with automatically configured TailwindCSS.`,
	PreRun: hooks.PreCheckNode,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Your projectname or directory (do not add spaces & do not use uppercase) (required)")
	generateCmd.PersistentFlags().BoolVar(&UseNPM, "use-npm", false, "Use npm for installing (not applicable to others) [defaults: false]")
	generateCmd.MarkFlagRequired("project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
