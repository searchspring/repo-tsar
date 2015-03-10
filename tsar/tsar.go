package tsar

import (
	"log"
	"errors"
	"fmt"

	"gopkg.in/libgit2/git2go.v22"
	"github.com/SearchSpring/RepoTsar/gitutils"
	"github.com/SearchSpring/RepoTsar/fileutils"
	"github.com/SearchSpring/RepoTsar/config"
)

type semaphore chan error

type RepoTsar struct{
	Config config.Config
	Branch string
	ReposList []string
	Signature *git.Signature
}

func (r *RepoTsar) Run() error {
	reposlist := r.ReposList 
	// if the reposlist is empty append all repos from config to reposlist
	c := r.Config.Repos
	if reposlist[0] == "" {
		// delete item from array
		reposlist = append(reposlist[:0], reposlist[0+1:]...)
		for k := range c {
			reposlist = append(reposlist,k)
		}
	}

	// Semaphore for concurrency
	thrnum := len(reposlist)
	sem := make(semaphore, thrnum)
	for _, key := range reposlist {
		go func(k string) {

			_,ok := c[k]
			if ! ok {
				err := errors.New(fmt.Sprintf("Repo %#v is not defined in config\n", k))
				sem <- err
				return
			}
			log.Printf("[%s, url: %s, path: %s, branch: %s]", k, c[k].URL, c[k].Path, c[k].Branch)
	
			// Createpath 
			path,err := fileutils.CreatePath(c[k].Path)
			if err != nil {
				sem <-err
				return
			}
			
			// Clone Repo
			cloneinfo := &gitutils.CloneInfo{
				Reponame: k,
				Path: path,
				URL: c[k].URL,
				Branch: c[k].Branch,
			}
			repo, err := cloneinfo.CloneRepo()
			if err != nil {
				sem <-err
				return
			}
			
			// Git Pull
			pullinfo := gitutils.PullInfo{
				Reponame: k,
				Repo: repo,
				Branch: c[k].Branch,
			}
			err = pullinfo.GitPull()
			if err != nil {
				sem <-err
				return
			}
	
			// If branch option, branch and checkout selected repos
			if r.Branch != "" {
				branchinfo := &gitutils.BranchInfo{
					Reponame: k,
					Branchname: r.Branch,
					Msg: "RepoTsar Branching",
					Repo: *repo,
					Signature: r.Signature,
				}
				err = branchinfo.GitBranch()
				if err != nil {
					sem <-err
					return
				}		
			}
			sem <-nil
			return
		}(key)
	}

	// Wait for threads to finish
	for i := 0; i < thrnum; i++ { 
		err := <-sem
		if err != nil {
			return err
		}
	}
	return nil
}