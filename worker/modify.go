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
	cmd := exec.Command("npx", w.appConfig.Tailwindconfiginstall...)
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
	for _, i := range w.appConfig.Modify {
		filename := path.Join(w.projectDir, i.Filename)
		fileBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			w.installSpinner.Errorf("Error trying to modify `%s`. Please modify it on your own...\n", i.Filename)
		}

		// replace each set strings tp be replaced
		var content string
		for _, r := range i.Replacecontent {
			content = strings.Replace(string(fileBytes), r.Textstring, r.Replacestring, 1)
		}

		// write files
		w.writer(filename, content)
	}

	// write new `tailwind.css`
	w.tailwindWriter()

	// LOOP INTO THE FILES TO BE REMOVED
	for _, v := range w.appConfig.Remove {
		for _, file := range v.Files {
			// remove each
			if err := os.Remove(path.Join(w.projectDir, v.Folder, file)); err != nil {
				w.installSpinner.Errorf("Error trying to remove `%s`. You can try to remove it on your own...\n", file)
			}
		}
	}

	// LOOP INTO THE FILES TO BE CREATED IF THERE IS
	if len(w.appConfig.Otherfiles) > 0 {
		for _, o := range w.appConfig.Otherfiles {
			w.writer(path.Join(w.projectDir, o.Filename), o.Content)
		}
	}
}

// adds or creates `tailwind.css` to the given path
func (w *Worker) tailwindWriter() {
	splts := strings.Split(w.appConfig.Tailwindpath, "/")
	folders := splts[:len(splts)-1]

	if err := os.MkdirAll(path.Join(w.projectDir, strings.Join(folders[:], "/")), 0755); err != nil {
		log.Fatalf("\nThere was a problem trying to create `%s`. Please create it your own, .. ", w.appConfig.Tailwindpath)
	}

	twContent := tailwindCSS

	// for vite-vue3
	if w.AppType == "vite-vue3" {
		twContent = "/*! @import */\n" + tailwindCSS
	}

	w.writer(path.Join(w.projectDir, w.appConfig.Tailwindpath), twContent)
}
