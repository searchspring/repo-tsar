package gitutils

import (
	"fmt"
	
	"gopkg.in/libgit2/git2go.v22"
)

type CloneInfo struct {
	Reponame string
	Path string
	URL string
	Branch string
}

// Clone a repo
func (c *CloneInfo) CloneRepo() (*git.Repository, error) {
	repo, err := git.OpenRepository(c.Path)
	if err != nil {
		cloneOptions := &git.CloneOptions{
			Bare:             false,
			CheckoutBranch:   c.Branch,
			CheckoutOpts: &git.CheckoutOpts{
				Strategy: git.CheckoutSafe,
			},
		}
		repo, err := git.Clone(c.URL,c.Path,cloneOptions)
		if err != nil {
			ret := fmt.Errorf(" %s : %s\n", err, c.URL)
			return nil, ret
		}
		return repo, nil
	}
	return repo, nil
}
