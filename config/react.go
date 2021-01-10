package config

// FOR REACT (CREATE-REACT-APP)
func newReactApp() *MainConfigApp {
	return &MainConfigApp{
		ID:      "react",
		Name:    "React",
		Apptype: "js",
		Installer: []appInstaller{
			{
				Pkgmanager:        "npm",
				Pkgmanagercommand: []string{"install"},
				Pkginstaller:      "npx",
				Pkginstargs:       []string{"create-react-app"},
			},
			{
				Pkgmanager:        "yarn",
				Pkgmanagercommand: []string{"add"},
				Pkginstaller:      "yarn",
				Pkginstargs:       []string{"create", "react-app"},
			},
		},
		Aftercreateinstall: false,
		Requiredpackages:   []string{"@craco/craco", "tailwindcss@npm:@tailwindcss/postcss7-compat", "@tailwindcss/postcss7-compat", "postcss@^7", "autoprefixer@^9"},
		Modify: []appModifier{
			{
				Filename: "tailwind.config.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "purge: [],",
						Replacestring: "purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],",
					},
				},
			},
			{
				Filename: "src/index.js",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "index.css",
						Replacestring: "tailwind.css",
					},
				},
			},
			{
				Filename: "package.json",
				Replacecontent: []modifyReplace{
					{
						Textstring:    "react-scripts start",
						Replacestring: "craco start",
					},
					{
						Textstring:    "react-scripts build",
						Replacestring: "craco build",
					},
					{
						Textstring:    "react-scripts test",
						Replacestring: "craco test",
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
		Tailwindconfiginstall: []string{"tailwindcss", "init"},
		Otherfiles: []additionalFiles{
			{
				Filename: "craco.config.js",
				Content: `module.exports = {
  style: {
    postcss: {
      plugins: [
        require('tailwindcss'),
        require('autoprefixer'),
      ],
    },
  },
}`,
			},
		},
	}
}
