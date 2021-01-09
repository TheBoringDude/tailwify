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

// functions to add minimal tailwind config
func (w *Worker) addTailwindMinimalConfig() {
	// configure
	// this will create minimal
	// tailwind.config.js and postcss.config.js
	cmd := exec.Command("npx", w.appConfig.tailwindConfigInstall...)
	cmd.Dir = w.projectDir

	// run
	if err := cmd.Run(); err != nil {
		w.installSpinner.Error("There was a problem trying to add minimal configurations...")
		log.Fatal(err)
	}
}

// just a spinner container
func (w *Worker) fileModifier() {
	w.installSpinner = spinner.New("Configuring TailwindCSS")
	w.installSpinner.Start()

	// modify files
	w.modify()

	// show success message
	w.installSpinner.Success("Succesfully configured TailwindCSS!")
}

// just a simple re-writer handler
func (w *Worker) writer(filename, content string) {
	if err := ioutil.WriteFile(filename, []byte(content), 0755); err != nil {
		log.Fatal(err)
	}
}

// main function that modifies other files
// => replace strings
// => remove unnecessary files
// => creates tailwind config files
// => writes new `tailwind.css`
// => creates additional files
func (w *Worker) modify() {
	// create config files
	w.addTailwindMinimalConfig()

	// LOOP INTO THE FILES TO BE MODIFIED
	for _, i := range w.appConfig.modify {
		filename := path.Join(w.projectDir, i.filename)
		fileBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			w.installSpinner.Errorf("Error trying to modify `%s`. Please modify it on your own...\n", i.filename)
		}

		// replace each set strings tp be replaced
		var content string
		for _, r := range i.replaceContent {
			content = strings.Replace(string(fileBytes), r.textString, r.replaceString, 1)
		}

		// write files
		w.writer(filename, content)
	}

	// write new `tailwind.css`
	w.writer(path.Join(w.projectDir, w.appConfig.tailwindPath), tailwindCSS)

	// LOOP INTO THE FILES TO BE REMOVED
	for _, v := range w.appConfig.remove {
		for _, file := range v.files {
			// remove each
			if err := os.Remove(path.Join(w.projectDir, v.folder, file)); err != nil {
				w.installSpinner.Errorf("Error trying to remove `%s`. You can try to remove it on your own...\n", file)
			}
		}
	}

	// LOOP INTO THE FILES TO BE CREATED IF THERE IS
	if len(w.appConfig.otherFiles) > 0 {
		for _, o := range w.appConfig.otherFiles {
			w.writer(o.filename, o.content)
		}
	}
}
