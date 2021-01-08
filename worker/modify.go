package worker

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/leaanthony/spinner"
)

// this will be used for the styling,
// since this is easier to be modified
const tailwindCSS = `@tailwind base;
@tailwind components;
@tailwind utilities;`

func (w *Worker) addTailwindMinimalConfig() {
	// configure
	// this will create minimal
	// tailwind.config.js and postcss.config.js
	cmd := exec.Command("npx", "tailwindcss", "init", "-p")
	cmd.Dir = w.projectDir

	// run
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem trying to add minimal configurations...")
		log.Fatal(err)
	}
}

// main handler for modifying files ...
func (w *Worker) fileModifier() {
	w.installSpinner = spinner.New("Configuring TailwindCSS")
	w.installSpinner.Start()

	if w.AppType == "next" {
		w.modifyNextJs()
	}

	// show success message
	w.installSpinner.Success("Succesfully configured TailwindCSS!")
}

// just a simple re-writer handler
func (w *Worker) writer(filename, content string) {
	if err := ioutil.WriteFile(filename, []byte(content), 0755); err != nil {
		log.Fatal(err)
	}
}

// FOR NEXT.JS
func (w *Worker) modifyNextJs() {
	// create config files
	w.addTailwindMinimalConfig()

	// LOOP INTO THE FILES TO BE MODIFIED
	for _, i := range w.appConfig.modify {
		filename := path.Join(w.projectDir, i.filename)
		fileBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			w.installSpinner.Errorf("Error trying to modify `%s`. Please modify it on your own...\n", i.filename)
		}

		content := strings.Replace(string(fileBytes), i.replaceContent.textString, i.replaceContent.replaceString, 1)

		// write files
		w.writer(filename, content)
	}

	// write new `tailwind.css`
	w.writer(path.Join(w.projectDir, "styles/tailwind.css"), tailwindCSS)

	// LOOP INTO THE FILES TO BE REMOVED
	for _, v := range w.appConfig.remove {
		for _, file := range v.files {
			// remove each
			if err := os.Remove(path.Join(w.projectDir, v.folder, file)); err != nil {
				w.installSpinner.Errorf("Error trying to remove `%s`. You can try to remove it on your own...\n", file)
			}
		}
	}
}
