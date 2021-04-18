package hooks

import (
	"os"

	"github.com/TheBoringDude/tailwify/utils"
	"github.com/leaanthony/spinner"
	"github.com/spf13/cobra"
)

// PreCheckNode checks if the defined package manager is installed
func PreCheckNode(cmd *cobra.Command, args []string) {
	preCheckSpinner := spinner.New("Checking for Node installation...")
	preCheckSpinner.Start()

	// check for node install
	if check := utils.CheckRequiredInstalled("node", "-v"); check {
		preCheckSpinner.Success("Node is installed")
	} else {
		preCheckSpinner.Error("Please install NODE in order to use this CLI!")

		// exit if Node is not installed
		os.Exit(0)
	}
}
