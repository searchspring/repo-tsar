package gitutils

import (
	"gopkg.in/libgit2/git2go.v22"
)

type BranchInfo struct {
	Reponame string
	Branchname string
	Msg string
	Repo git.Repository
	Signature *git.Signature
}

func (b *BranchInfo) createBranch() (error){
	head,err := b.Repo.Head()
	if err != nil {
		return err
	}
	headCommit, err := b.Repo.LookupCommit(head.Target())
	if err != nil {
		return err
	}	
	_,err = b.Repo.CreateBranch(b.Branchname, headCommit, false )
	
	if err != nil {
		return err
	}
	return nil
}

func (b *BranchInfo) GitBranch() (error) {
	
	// If branch doesn't exist, create it
	_,err := b.Repo.LookupReference("refs/heads/"+b.Branchname)
	if err != nil{
		err = b.createBranch()
		if err != nil {
			return err
		}
	}

	// Checkout Branch
	head,err := b.Repo.Head()
	if err != nil {
		return err
	}
	headCommit, err := b.Repo.LookupCommit(head.Target())
	if err != nil {
		return err
	}	
	branchTree,err := b.Repo.LookupTree(headCommit.TreeId())
	if err != nil {
		return err
	}
	err = b.Repo.CheckoutTree(branchTree, &git.CheckoutOpts{})
	if err != nil {
		return err
	}
	err = b.Repo.SetHead("refs/heads/"+b.Branchname)
	if err != nil {
		return err
	}
	return nil
}