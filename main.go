package main

import(
	"time"
	"flag"
	"strings"
	"log"
	"fmt"

	"gopkg.in/libgit2/git2go.v22"
	"github.com/SearchSpring/RepoTsar/tsar"
	"github.com/SearchSpring/RepoTsar/config"
)

var configFileName string
var branch string
var repos string
var version bool

const (
	versioninfo = "v0.1.2"
)

func main() {
	// Parse commandline 
	flag.StringVar(&configFileName, "config", "repotsar.yml", "YAML config file")
	flag.StringVar(&branch, "branch", "", "Create branch in repos")
	flag.StringVar(&repos, "repos", "", "Non-spaced Comma separated list of repos (defaults to all)")
	flag.BoolVar(&version, "version",false,"RepoTsar version")
	flag.Parse()

	if version == true {
		fmt.Printf("RepoTsar version %s\n", versioninfo)
		return
	}

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
