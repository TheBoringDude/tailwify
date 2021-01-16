package cmd

import (
	"github.com/spf13/cobra"
)

var projectName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate and setup a starter template",
	Long: `It will configure a simple and starter app for you 
with automatically configured TailwindCSS.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("generate called")
	// },
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	generateCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Your projectname or directory (do not add spaces & do not use uppercase) (required)")
	generateCmd.MarkFlagRequired("project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
