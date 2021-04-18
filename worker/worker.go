package worker

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/leaanthony/spinner"
	"github.com/ootiq/tailwify/v1/config"
)

// Worker => main handler for all apps
type Worker struct {
	AppType        string
	ProjectName    string
	JsApp          bool
	jsPkger        string
	PhpApp         bool
	BasicApp       bool
	projectDir     string
	wdPath         string
	installSpinner *spinner.Spinner
	appConfig      *config.MainConfigApp
}

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
func (w *Worker) Start(useNpm bool) {
	// print ascii logo
	fmt.Println(ASCIILogo)

	// check if projectName is blank or not
	if w.ProjectName == "" {
		fmt.Println("\n  Please set a ProjectName by using the `-p` flag...")
		os.Exit(0)
	}

	// set first the paths
	w.getPath()

	// identify which package manager to use
	w.jsPkger = "yarn"
	if useNpm || w.AppType == "gatsby" {
		w.jsPkger = "npm"
	}

	// then, check for the required apps
	w.checkApps()

	// SET THE CONFIG APPCONFIG TO BE USED
	w.appConfig = config.Configurator(w.AppType)

	// run the main installer worker
	w.run()

	// show success message
	fmt.Printf("\n  Your project `%s` has been succesfully configured at '%s' \n=> Feel free to modify it again at your own cost... | TheBoringDude\n", w.ProjectName, w.projectDir)
}
