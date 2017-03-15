package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/spf13/cobra"
	"github.com/praqma/git-phlow/ui"
)

// purgeCmd represents the purge command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean removes all delivered branches",
	Long: fmt.Sprintf(`
%s removes all branches prefixed with 'delivered/'.
It deletes safely by running 'git branch -d'. By default, both local and remote branches are deleted.
`, ui.Bold("phlow clean")),
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)
		remote := githandler.ConfigBranchRemote(defaultBranch)

		phlow.Clean(remote)
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)

	//Run clean on local branches
	cleanCmd.Flags().BoolVarP(&options.GlobalFlagLocal, "local", "l", false, "remove local delivered branches only")

	//Run clean forcefully
	cleanCmd.Flags().BoolVarP(&options.GlobalFlagForce, "force", "f", false, "force remove delivered branches")

}
