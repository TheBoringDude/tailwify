package worker

// main installer configurator
type mainConfigApp struct {
	id                    string
	name                  string
	apptype               string
	installer             []appInstaller
	afterCreateInstall    bool
	requiredPackages      []string
	modify                []appModifier
	remove                []appFileRemover
	tailwindPath          string
	tailwindConfigInstall []string
	otherFiles            []additionalFiles
}

// for node package management
type appInstaller struct {
	pkgManager        string
	pkgManagerCommand []string
	pkgInstaller      string
	pkgInstArgs       []string
}

// after the `create-`
type appAfterInstall struct {
	commandArgs []string
}

// additional files to be added
type additionalFiles struct {
	filename string
	content  string
}

// files to be modified
type appModifier struct {
	filename       string
	replaceContent []modifyReplace
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
		afterCreateInstall: false,
		requiredPackages:   []string{"tailwindcss@latest", "postcss@latest", "autoprefixer@latest"},
		modify: []appModifier{
			{
				filename: "tailwind.config.js",
				replaceContent: []modifyReplace{
					{
						textString:    "purge: [],",
						replaceString: "purge: ['./pages/**/*.js', './components/**/*.js'],",
					},
				},
			},
			{
				filename: "pages/_app.js",
				replaceContent: []modifyReplace{
					{
						textString:    "../styles/globals.css",
						replaceString: "../styles/tailwind.css",
					},
				},
			},
		},
		remove: []appFileRemover{
			{
				folder: "styles",
				files:  []string{"Home.module.css", "globals.css"},
			},
		},
		tailwindPath:          "styles/tailwind.css",
		tailwindConfigInstall: []string{"tailwindcss", "init", "-p"},
		otherFiles:            []additionalFiles{},
	}

	return app
}

// FOR GATSBY.JS
func (w *Worker) newGatsbyJs() *mainConfigApp {
	app := &mainConfigApp{
		id:      "gatsby",
		name:    "GatsbyJS",
		apptype: "js",
		installer: []appInstaller{
			{
				pkgManager:        "npm",
				pkgManagerCommand: []string{"install"},
				pkgInstaller:      "gatsby",
				pkgInstArgs:       []string{"new"},
			},
		},
		afterCreateInstall: false,
		requiredPackages:   []string{"gatsby-plugin-postcss", "tailwindcss@latest", "postcss@latest", "autoprefixer@latest"},
		modify: []appModifier{
			{
				filename: "tailwind.config.js",
				replaceContent: []modifyReplace{
					{
						textString:    "purge: [],",
						replaceString: "purge: ['./src/**/*.{js,jsx,ts,tsx}'],",
					},
				},
			},
			{
				filename: "gatsby-config.js",
				replaceContent: []modifyReplace{
					{
						textString:    "plugins: [],",
						replaceString: "plugins: ['gatsby-plugin-postcss'],",
					},
				},
			},
		},
		remove:                []appFileRemover{},
		tailwindPath:          "src/styles/tailwind.css",
		tailwindConfigInstall: []string{"tailwindcss", "init", "-p"},
		otherFiles: []additionalFiles{
			{
				filename: "gatsby-browser.js",
				content:  `import './src/styles/global.css';`,
			},
		},
	}

	return app
}
