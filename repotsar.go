package main

import (
	"time"
	"log"
	"strings"
	"flag"
	"gopkg.in/libgit2/git2go.v22"
	"github.com/SearchSpring/RepoTsar/gitutils"
	"github.com/SearchSpring/RepoTsar/fileutils"
	"github.com/SearchSpring/RepoTsar/config"
)

var configFileName string
var branch string
var reposCSV string
type empty struct {}
type semaphore chan empty

func main() {
	// Parse commandline 
	flag.StringVar(&configFileName, "config", "repotsar.yml", "YAML config file")
	flag.StringVar(&branch, "branch", "", "Create branch in repos")
	flag.StringVar(&reposCSV, "repos", "", "Non-spaced Comma separated list of repos (defaults to all)")
	flag.Parse()

	config,err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	c := config.Repos

	// Git Signature
	signature := &git.Signature{
		Name: config.Signature.Name,
		Email: config.Signature.Email,
		When: time.Now(),
	}

	reposlist := strings.Split(reposCSV, ",") 
	// if the reposlist is empty append all repos from config to reposlist
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
				log.Fatalf("Repo %#v is not defined in config\n", k)
			}
			log.Printf("[%s, url: %s, path: %s, branch: %s]", k, c[k].URL, c[k].Path, c[k].Branch)
	
			// Createpath 
			path,err := fileutils.CreatePath(c[k].Path)
			if err != nil {
				log.Fatal(err)
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
				log.Fatal(err)
			}
			
			// Git Pull
			pullinfo := gitutils.PullInfo{
				Reponame: k,
				Repo: repo,
				Branch: c[k].Branch,
			}
			err = pullinfo.GitPull()
			if err != nil {
				log.Fatal(err)
			}
	
			// If branch option, branch and checkout selected repos
			if branch != "" {
				branchinfo := &gitutils.BranchInfo{
					Reponame: k,
					Branchname: branch,
					Msg: "RepoTsar Branching",
					Repo: *repo,
					Signature: signature,
				}
				err = branchinfo.GitBranch()
				if err != nil {
					log.Fatal(err)
				}		
			}
			e := empty{}
			sem <- e
		}(key)
	}

	// Wait for threads to finish
	for i := 0; i < thrnum; i++ {<-sem}
}