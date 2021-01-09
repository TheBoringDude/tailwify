package config

// FOR VUE3 (CREATE-VITE-APP)
func newViteApp() *MainConfigApp {
	return &MainConfigApp{
		ID:      "vite-vue3",
		Name:    "Vue3 w/ Vite",
		Apptype: "js",
		Installer: []appInstaller{
			{
				Pkgmanager:        "npm",
				Pkgmanagercommand: []string{"install"},
				Pkginstaller:      "npx",
				Pkginstargs:       []string{"create-vite-app"},
			},
			{
				Pkgmanager:        "yarn",
				Pkgmanagercommand: []string{"add"},
				Pkginstaller:      "yarn",
				Pkginstargs:       []string{"create", "vite-app"},
			},
		},
		Aftercreateinstall: true,
		Requiredpackages:   []string{"-D", "tailwindcss@npm:@tailwindcss/postcss7-compat", "@tailwindcss/postcss7-compat", "postcss@^7", "autoprefixer@^9"},
		Modify: []appModifier{
			{
				Filename: "tailwind.config.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "purge: [],",
						Replacestring: "purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],",
					},
				},
			},
			{
				Filename: "src/main.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "./index.css",
						Replacestring: "./tailwind.css",
					},
				},
			},
		},
		Remove: []appFileRemover{
			{
				Folder: "src",
				Files:  []string{"index.css"},
			},
		},
		Tailwindpath:          "src/tailwind.css",
		Tailwindconfiginstall: []string{"tailwindcss", "init", "-p"},
		Otherfiles:            []additionalFiles{},
	}
}
