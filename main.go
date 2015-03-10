package main

import(
	"time"
	"flag"
	"strings"
	"log"

	"gopkg.in/libgit2/git2go.v22"
	"github.com/SearchSpring/RepoTsar/tsar"
	"github.com/SearchSpring/RepoTsar/config"
)

var configFileName string
var branch string
var repos string

func main() {
	// Parse commandline 
	flag.StringVar(&configFileName, "config", "repotsar.yml", "YAML config file")
	flag.StringVar(&branch, "branch", "", "Create branch in repos")
	flag.StringVar(&repos, "repos", "", "Non-spaced Comma separated list of repos (defaults to all)")
	flag.Parse()

	config,err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	reposlist := strings.Split(repos, ",")
	// Git Signature
	signature := &git.Signature{
		Name: config.Signature.Name,
		Email: config.Signature.Email,
		When: time.Now(),
	}
	tsar := &tsar.RepoTsar{
		Config: config,
		Branch: branch,
		ReposList: reposlist,
		Signature: signature,
	}
	err = tsar.Run()
	if err != nil {
		log.Fatal(err)
	} 
}