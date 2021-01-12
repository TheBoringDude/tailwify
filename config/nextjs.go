package config

// FOR NEXT.JS
func newNextJs() *MainConfigApp {
	return &MainConfigApp{
		ID:      "next",
		Name:    "NextJS",
		Apptype: "js",
		Installer: []appInstaller{
			{
				Pkgmanager:        "npm",
				Pkgmanagercommand: []string{"install"},
				Pkginstaller:      "npx",
				Pkginstargs:       []string{"create-next-app"},
			},
			{
				Pkgmanager:        "yarn",
				Pkgmanagercommand: []string{"add"},
				Pkginstaller:      "yarn",
				Pkginstargs:       []string{"create", "next-app"},
			},
		},
		Aftercreateinstall: false,
		Requiredpackages:   []string{"tailwindcss@latest", "postcss@latest", "autoprefixer@latest"},
		Modify: []appModifier{
			{
				Filename: "tailwind.config.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "purge: [],",
						Replacestring: "purge: ['./pages/**/*.js', './components/**/*.js'],",
					},
				},
			},
			{
				Filename: "pages/_app.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "../styles/globals.css",
						Replacestring: "../styles/tailwind.css",
					},
				},
			},
		},
		Remove: []appFileRemover{
			{
				Folder: "styles",
				Files:  []string{"globals.css"},
			},
		},
		Tailwindpath:          "styles/tailwind.css",
		Tailwindconfiginstall: []string{"tailwindcss", "init", "-p"},
		Otherfiles:            []additionalFiles{},
	}
}
