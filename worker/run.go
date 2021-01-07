package worker

import (
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
)

func (w *Worker) runNextJs() {
	// start install
	installSpinner = spinner.New("Installing NextJS")
	installSpinner.Start()

	// run create-next-app
	cmd := exec.Command("npx", "create-next-app", w.ProjectName)
	err := cmd.Run()
	if err != nil {
		installSpinner.Error("There was a problem while trying to install NextJS")
		log.Fatal(err)
	}

	// this will run on success
	installSpinner.Success("Succesfully installed NextJS")

	// install tailwind
	w.installTailwindNextJs()

	// configure and modify files
	w.modifyNextJs()
}
