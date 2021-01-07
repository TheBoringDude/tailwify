package worker

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/leaanthony/spinner"
)

// Worker => main handler for all apps
type Worker struct {
	AppType     string
	ProjectName string
	JsApp       bool
	jsPkger     string
	PhpApp      bool
	BasicApp    bool
	projectDir  string
	wdPath      string
}

var installSpinner = spinner.New()

// getPath will get the current working directory
func (w *Worker) getPath() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// set the working parent dir
	w.wdPath = cwd
	w.projectDir = path.Join(w.wdPath, w.ProjectName)
}

// Start starts generating the template
func (w *Worker) Start() {
	// set first the paths
	w.getPath()

	// then, check for the required apps
	w.checkApps()

	// run specific generators

	// for nextjs
	if w.AppType == "next" {
		w.runNextJs()
	}

	// show success message
	fmt.Printf("\n  Your project `%s` has been succesfully configured at '%s' \n=> Feel free to modify it again at your own cost... | TheBoringDude\n", w.ProjectName, w.projectDir)
}
