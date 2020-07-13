package cmd

import (
	"os"

	"github.com/eankeen/globe/config"
	"github.com/eankeen/globe/sync"
	"github.com/eankeen/globe/validate"
	"github.com/spf13/cobra"
)

var initsCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Globe's configuration files",
	Long:  `Initiates configuration files to be used by Globe`,
	Run: func(cmd *cobra.Command, args []string) {
		// get data
		storeDir := cmd.Flag("store-dir").Value.String()
		projectDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		project := config.GetData(projectDir, storeDir)

		// validate values
		validate.Validate(validate.ValidationValues{
			StoreDir: storeDir,
			Project:  project,
		})

		// use values
		sync.ProcessFiles(project, project.InitFiles.Files)
	},
}

func init() {
	RootCmd.AddCommand(initsCmd)
}
