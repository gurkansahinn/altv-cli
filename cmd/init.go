package cmd

import (
	"altv-cli/view"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Download alt:V binary files",
	Long:  `Downloads alt:V binary files quickly`,
	Run:   execute,
}

func execute(cmd *cobra.Command, args []string) {
	view.BranchSelectorView()
}

func init() {
	RootCmd.AddCommand(initCommand)
}
