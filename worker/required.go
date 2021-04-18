package worker

import (
	"fmt"
	"os"
	"strings"

	"github.com/leaanthony/spinner"
	"github.com/ootiq/tailwify/v1/utils"
)

func (w *Worker) checkApps() {
	upPkger := strings.ToUpper(w.jsPkger) // uppercase the text

	w.installSpinner = spinner.New(fmt.Sprintf("Checking for `%s` package manager", upPkger))
	w.installSpinner.Start()
	if check := utils.CheckRequiredInstalled(w.jsPkger, "-v"); !check {
		w.installSpinner.Errorf("`%s` is not installed. Please install it and try again.", upPkger)
		os.Exit(1)
	}

	// show success on verify
	w.installSpinner.Successf("Using `%s` for installing...", upPkger)
}
