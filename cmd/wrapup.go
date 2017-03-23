package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "add changes and make auto-commit",
	Long: fmt.Sprintf(`
%s adds all your work from the working area to the index, and makes an automatic close # commit,
which closes an issue on github eg. commit message => "close #42 fetch meaning of life"
`, ui.Format.Bold("wrapup")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.WrapUp()
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	wrapupCmd.Flags().BoolVar(&options.GlobalFlagHard, "hard", false, "Add all before commit")
}
