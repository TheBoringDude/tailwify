package config

// FOR GATSBY.JS
func newGatsbyJs() *MainConfigApp {
	return &MainConfigApp{
		ID:      "gatsby",
		Name:    "GatsbyJS",
		Apptype: "js",
		Installer: []appInstaller{
			{
				Pkgmanager:        "npm",
				Pkgmanagercommand: []string{"install"},
				Pkginstaller:      "gatsby",
				Pkginstargs:       []string{"new"},
			},
		},
		Aftercreateinstall: false,
		Requiredpackages:   []string{"gatsby-plugin-postcss", "tailwindcss@latest", "postcss@latest", "autoprefixer@latest"},
		Modify: []appModifier{
			{
				Filename: "tailwind.config.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "purge: [],",
						Replacestring: "purge: ['./src/**/*.{js,jsx,ts,tsx}'],",
					},
				},
			},
			{
				Filename: "gatsby-config.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "plugins: [],",
						Replacestring: "plugins: ['gatsby-plugin-postcss'],",
					},
				},
			},
		},
		Remove:                []appFileRemover{},
		Tailwindpath:          "src/styles/tailwind.css",
		Tailwindconfiginstall: []string{"tailwindcss", "init", "-p"},
		Otherfiles: []additionalFiles{
			{
				Filename: "gatsby-browser.js",
				Content:  `import './src/styles/global.css';`,
			},
		},
	}
}
