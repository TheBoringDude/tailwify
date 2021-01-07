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

func (w *Worker) modifyNextJs() {
	installSpinner = spinner.New("Configuring TailwindCSS")
	installSpinner.Start()

	// configure
	// this will create minimal
	// tailwind.config.js and postcss.config.js
	cmd := exec.Command("npx", "tailwindcss", "init", "-p")
	cmd.Dir = w.projectDir

	// run
	if err := cmd.Run(); err != nil {
		installSpinner.Error("There was a problem trying to add minimal configurations...")
		log.Fatal(err)
	}

	// modify the purge settings of tailwind.config.js
	filenameTwConfig := path.Join(w.projectDir, "tailwind.config.js")
	twConfigByte, err := ioutil.ReadFile(filenameTwConfig)
	if err != nil {
		installSpinner.Error("Error trying to modify `tailwind.config.js` purge settings.")
		log.Fatal(err)
	}
	twConfig := strings.Replace(string(twConfigByte), "purge: [],", "purge: ['./pages/**/*.js', './components/**/*.js'],", 1)

	// remove unnecessary files
	toremoveFiles := []string{"Home.module.css", "globals.css"}
	for _, i := range toremoveFiles {
		// remove each
		if err := os.Remove(path.Join(w.projectDir, "styles", i)); err != nil {
			installSpinner.Error("Error trying to remove `%s` \n", i)
			log.Fatal(err)
		}
	}

	// change the css with `tailwind.css`
	filenameTailwind := path.Join(w.projectDir, "styles/tailwind.css")

	// replace on the appJs
	filenameAppJs := path.Join(w.projectDir, "pages/_app.js")
	appJsBytes, err := ioutil.ReadFile(filenameAppJs)
	if err != nil {
		installSpinner.Error("Error trying to modify `pages/_app.js`")
		log.Fatal(err)
	}
	appJs := strings.Replace(string(appJsBytes), "../styles/globals.css", "../styles/tailwind.css", 1)

	// write files
	w.writer(installSpinner, []map[string]string{
		{
			"filename": filenameTwConfig,
			"content":  twConfig,
		},
		{
			"filename": filenameTailwind,
			"content":  tailwindCSS,
		},
		{
			"filename": filenameAppJs,
			"content":  appJs,
		},
	})

	// show success message
	installSpinner.Success("Succesfully configured TailwindCSS!")
}

// just a simple re-writer
func (w *Worker) writer(spinner *spinner.Spinner, contents []map[string]string) {
	for _, i := range contents {
		if err := ioutil.WriteFile(i["filename"], []byte(i["content"]), 0755); err != nil {
			spinner.Errorf("Error trying to modify %s \n", i["filename"])
			log.Fatal(err)
		}
	}
}
