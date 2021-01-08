package worker

import (
	"os"
	"strings"

	"github.com/TheBoringDude/tailwify/utils"
	"github.com/leaanthony/spinner"
)

var nodeApps = []map[string]string{
	{
		"name":         "Node",
		"command":      "node",
		"command-args": "-v",
	},
	{
		"name":         "NPM",
		"command":      "npm",
		"command-args": "-v",
	},
	{
		"name":         "Yarn",
		"command":      "yarn",
		"command-args": "-v",
	},
}

// checkApps will check the required apps
// for generating specific templates
func (w *Worker) checkApps() {
	// check for node and npm if it is a JS app
	if w.JsApp {
		// check node apps
		w.checkNode()
	}

}

func (w *Worker) checkNode() {
	w.installSpinner = spinner.New("Checking for Node installation...")
	w.installSpinner.Start()

	// check for node install
	if check := utils.CheckRequiredInstalled(nodeApps[0]["command"], nodeApps[0]["command-args"]); check {
		w.installSpinner.Success("Node is installed")
	} else {
		w.installSpinner.Error("Please install NODE in order to use this CLI!")

		// exit if Node is not installed
		os.Exit(0)
	}

	w.installSpinner = spinner.New("Checking for installed package manager...")
	w.installSpinner.Start()

	// installed package manager
	pkger := ""

	// check for yarn if installed
	if check := utils.CheckRequiredInstalled(nodeApps[2]["command"], nodeApps[2]["command-args"]); check {
		pkger = nodeApps[2]["name"]
	} else {
		// check for npm if installed
		if check := utils.CheckRequiredInstalled(nodeApps[1]["command"], nodeApps[1]["command-args"]); check {
			pkger = nodeApps[1]["name"]
		}
	}

	// show error and stop the app
	// if yarn / npm is not installed
	if pkger == "" {
		w.installSpinner.Error("No `package manager` installed. Please install NPM or Yarn and try again.")

		// exit the app
		os.Exit(0)
	}

	// show success on verify
	w.installSpinner.Successf("Using `%s` for installing...", pkger)

	// set the nodejs pkg manager
	w.jsPkger = strings.ToLower(pkger)
}
