package phlow

import (
	"fmt"
	"strings"

	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
)

//WrapUp ...
func WrapUp() {

	//Add all files to index
	if options.GlobalFlagHard {
		fmt.Fprintln(os.Stdout, "Adding files to index")
		if err := githandler.Add(); err != nil {
			fmt.Println("Project files could not be added: " + err.Error())
			return
		}
	}

	//Retrieve branch info - current branch
	info, _ := githandler.Branch()
	commitMessage := "close #" + strings.Replace(info.Current, "-", " ", -1)

	if _, err := githandler.Commit(commitMessage); err != nil {
		fmt.Println("Nothing to commit!")
		return
	}
	fmt.Fprintln(os.Stdout, commitMessage)
	fmt.Println("Remember to squash your commits")
}
