package worker

import (
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
)

func (w *Worker) tailwindNodeArgs() []string {
	var nodePkgerCommand []string

	// just loop in the package managers
	for _, p := range w.appConfig.installer {
		if p.pkgManager == w.jsPkger {
			nodePkgerCommand = p.pkgManagerCommand

			break
		}
	}

	newArgs := append(nodePkgerCommand, w.appConfig.requiredPackages...)

	return newArgs
}

func (w *Worker) installTailwindNextJs() {
	w.installSpinner = spinner.New("Installing TailwindCSS and other required libraries...")
	w.installSpinner.Start()

	// set install arguments
	installArgs := w.tailwindNodeArgs()

	// install tailwind
	cmd := exec.Command(w.jsPkger, installArgs...)
	cmd.Dir = w.projectDir

	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("Error installing TailwindCSS")
		log.Fatal(err)
	}

	w.installSpinner.Success("TalwindCSS has been succesfully installed!")
}
