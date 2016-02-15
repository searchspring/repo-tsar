package gitutils

import (
	"testing"
	"os"
	"io/ioutil"

)

func TestPullInfoStruct(t *testing.T){
	path, err := ioutil.TempDir("", "RepoTsar")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(path)

	cloneinfo := &CloneInfo {
		Reponame: "TestRepo",
		Path: path,
		URL: "https://github.com/libgit2/git2go.git",
		Branch: "master",
	}
	repo,err := cloneinfo.CloneRepo()
	if err != nil {
		t.Error(err)
	}

	pullinfo := PullInfo{
		Reponame: "Test Repo",
		Repo: repo,
		Branch: "master",
	}

	err = pullinfo.GitPull()
	if err != nil {
		t.Error(err)
	}
}