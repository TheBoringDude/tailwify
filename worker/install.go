package worker

import (
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
)

func (w *Worker) installTailwindNextJs() {
	installSpinner = spinner.New("Installing TailwindCSS and other required libraries...")
	installSpinner.Start()

	pkgFunc := ""
	if w.jsPkger == "yarn" {
		pkgFunc = "add"
	} else if w.jsPkger == "npm" {
		pkgFunc = "install"
	}

	// install tailwind
	cmd := exec.Command(w.jsPkger, pkgFunc, "tailwindcss@latest", "postcss@latest", "autoprefixer@latest")
	cmd.Dir = w.projectDir

	if err := cmd.Run(); err != nil {
		installSpinner.Error("Error installing TailwindCSS")
		log.Fatal(err)
	}

	installSpinner.Success("TalwindCSS has been succesfully installed!")
}
