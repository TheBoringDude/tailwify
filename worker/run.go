package worker

import (
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
)

// check for the main installer
func (w *Worker) checkNodePackager() (string, []string) {
	// check command
	appInstall := w.appConfig.Installer[w.jsPkger]

	// if the package manager is not `npm` or `yarn`
	// this will error if the code is modified
	if appInstall.Pkginstaller == "" {
		log.Fatal("Error NODE Package Manager: " + w.jsPkger)
	}

	// append the project name to the args
	newArgs := append(appInstall.Pkginstargs, w.ProjectName)

	return appInstall.Pkginstaller, newArgs
}

// MAIN HANDLER FOR EVERYTHING
func (w *Worker) run() {
	// start install
	w.installSpinner = spinner.New("Installing " + w.appConfig.Name)
	w.installSpinner.Start()

	// run the app installer //
	cmdCommand, cmdArg := w.checkNodePackager()
	cmd := exec.Command(cmdCommand, cmdArg...)
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem while trying to install " + w.appConfig.Name)
		log.Fatal(err)
	}

	// after create-
	// if needed
	if w.appConfig.Aftercreateinstall {
		w.afterInstall()
	}

	// show success message
	w.installSpinner.Success("Succesfully installed " + w.appConfig.Name)

	// install tailwind
	w.installTailwind()

	// configure and modify files
	w.fileModifier()
}

// after install function
// after the create- something, if there is
// some of the frameworks, do not automatically install it
// like, .. having `npm install`  after
func (w *Worker) afterInstall() {
	cmd := exec.Command(w.jsPkger, "install")
	cmd.Dir = w.projectDir

	// install it
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem while trying to install " + w.appConfig.Name)
		log.Fatal(err)
	}
}
