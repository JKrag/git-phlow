package gitwrapper

import "github.com/praqma/git-phlow/subprocess"

const (
	all string = "--all"
)

//GitFetch
//interface for git fetch commands
type Fetch interface {
	Fetch(origin bool) (string, error)
}

type fetch struct {
	gitFetchCommand string
}


//NewFetch
//Constructor for fetch struct
func NewFetch() *fetch {
	return &fetch{gitFetchCommand:"fetch"}

}

//Fetch
//Doing a normal git fetch
func (f *fetch) Fetch(fromOrigin bool) (string, error) {

	var message string
	var err error

	if fromOrigin {
		message, err = subprocess.SimpleExec(GitCommand, f.gitFetchCommand, all)
	} else {
		message, err = subprocess.SimpleExec(GitCommand, f.gitFetchCommand)
	}

	if err != nil {
		return "", err
	}

	return message, nil
}
