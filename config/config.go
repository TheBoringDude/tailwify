package config

// MainConfigApp main config handler
// main installer configurator
type MainConfigApp struct {
	ID                    string
	Name                  string
	Apptype               string
	Installer             map[string]appInstaller
	Aftercreateinstall    bool
	Requiredpackages      []string
	Modify                []appModifier
	Remove                []appFileRemover
	Tailwindpath          string
	Tailwindconfiginstall []string
	Otherfiles            []additionalFiles
}

// for node package management
type appInstaller struct {
	Pkgmanagercommand []string
	Pkginstaller      string
	Pkginstargs       []string
}

// additional files to be added
type additionalFiles struct {
	Filename string
	Content  string
}

// files to be modified
type appModifier struct {
	Filename       string
	Replacecontent []modifyReplace
}
type modifyReplace struct {
	Textstring    string
	Replacestring string
}

// files to be removed
type appFileRemover struct {
	Folder string
	Files  []string
}

// Configurator handles the config
// to be passed to the worker
// depending to the apptype set
func Configurator(appType string) *MainConfigApp {
	return frameworks[appType].config
}
