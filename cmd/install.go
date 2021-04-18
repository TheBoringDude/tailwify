package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/leaanthony/spinner"
	"github.com/ootiq/tailwify/v1/config"
	"github.com/spf13/cobra"
)

// use postcss7
var postcss7 bool

// project dir folder
var dirFolder string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install TailwindCSS with it's required packages.",
	Long: `Install only TailwindCSS, PostCSS and Autoprefixer without
configuring the folder structures / files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// spinner for cool ui :)
		spinner := spinner.New()
		spinner.Start("Installing TailwindCSS and other required packages...")

		install := &exec.Cmd{}

		// check if compat-build is defined in flags
		if postcss7 {
			install = exec.Command("yarn", config.TailwindCompatBuild...)
		} else {
			install = exec.Command("yarn", config.TailwindCSSLatest...)
		}

		// install tailwind
		if err := install.Run(); err != nil {
			spinner.Error("Failed installing TailwindCSS")
			log.Fatal(err)
		}

		spinner.Success("TailwindCSS has been successfully installed!")

		// print success message
		fmt.Printf("\n\nSuccessfully installed TailwindCSS at %s", dirFolder)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")
	installCmd.Flags().BoolVarP(&postcss7, "compat-build", "c", false, "Install the compatibility build (PostCSS7)")
	installCmd.Flags().StringVarP(&dirFolder, "dir", "d", "", "Project directory to install TailwindCSS")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
