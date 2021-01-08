package worker

// main installer configurator
type mainConfigApp struct {
	id               string
	name             string
	apptype          string
	installer        []appInstaller
	requiredPackages []string
	modify           []appModifier
	remove           []appFileRemover
}

// for node package management
type appInstaller struct {
	pkgManager        string
	pkgManagerCommand []string
	pkgInstaller      string
	pkgInstArgs       []string
}

// files to be modified
type appModifier struct {
	filename       string
	replaceContent modifyReplace
}
type modifyReplace struct {
	textString    string
	replaceString string
}

// files to be removed
type appFileRemover struct {
	folder string
	files  []string
}

// FOR NEXT.JS
func (w *Worker) newNextJs() *mainConfigApp {
	app := &mainConfigApp{
		id:      "next",
		name:    "NextJS",
		apptype: "js",
		installer: []appInstaller{
			{
				pkgManager:        "npm",
				pkgManagerCommand: []string{"install"},
				pkgInstaller:      "npx",
				pkgInstArgs:       []string{"create-next-app"},
			},
			{
				pkgManager:        "yarn",
				pkgManagerCommand: []string{"add"},
				pkgInstaller:      "yarn",
				pkgInstArgs:       []string{"create", "next-app"},
			},
		},
		requiredPackages: []string{"tailwindcss@latest", "postcss@latest", "autoprefixer@latest"},
		modify: []appModifier{
			{
				filename: "tailwind.config.js",
				replaceContent: modifyReplace{
					textString:    "purge: [],",
					replaceString: "purge: ['./pages/**/*.js', './components/**/*.js'],",
				},
			},
			{
				filename: "pages/_app.js",
				replaceContent: modifyReplace{
					textString:    "../styles/globals.css",
					replaceString: "../styles/tailwind.css",
				},
			},
		},
		remove: []appFileRemover{
			{
				folder: "styles",
				files:  []string{"Home.module.css", "globals.css"},
			},
		},
	}

	return app
}
