package gitutils

import (
	"fmt"
	
	"gopkg.in/libgit2/git2go.v22"
)

type PullInfo struct {
	Reponame string
	Repo *git.Repository
	Branch string
}

// Simulate Git pull
func (p *PullInfo ) GitPull() ( error ) {

	// fetch
	remotes,err := p.Repo.ListRemotes()
	if err != nil {
		return err
	} else {
		origin, err := p.Repo.LookupRemote(remotes[0])
		remoteCallbacks := &git.RemoteCallbacks{
			CredentialsCallback:  credentialsCallback,
			CertificateCheckCallback: certificateCheckCallback,
		}
		origin.SetCallbacks(remoteCallbacks)
		if err != nil {
			return err
		} else {
			refspec := make([]string, 0)
			err = origin.Fetch(refspec, "")
			if err != nil {
				return err
			} 
		}
	}

	// merge
	ref := fmt.Sprintf("refs/heads/%s", p.Branch)
	master, err := p.Repo.LookupReference(ref)
	if err != nil {
		return err
	}
	mergeHead, err := p.Repo.AnnotatedCommitFromRef(master)
	if err != nil {
		return err
	}
	mergeHeads := make([]*git.AnnotatedCommit,1)
	mergeHeads[0] = mergeHead
	p.Repo.Merge(mergeHeads,&git.MergeOptions{},&git.CheckoutOpts{})
	return nil
}