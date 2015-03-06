package gitutils

import (
	"gopkg.in/libgit2/git2go.v22"
)


// Callbacks for git2go
func credentialsCallback(url string, username_from_url string, allowed_types git.CredType) (git.ErrorCode, *git.Cred) {
	code, cred := git.NewCredSshKeyFromAgent(username_from_url)
	return git.ErrorCode(code), &cred
}

func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}