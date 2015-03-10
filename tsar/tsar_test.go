package tsar

import (
	"testing"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/libgit2/git2go.v22"
	"github.com/SearchSpring/RepoTsar/config"
)

func TestRepoTsarStruct(t *testing.T){
	testpath, err := ioutil.TempDir("", "RepoTsar")
	if err != nil {
		t.Error(err)
	}

	defer os.RemoveAll(testpath)
	config := config.Config{
		Repos: map[string]config.Repo{
			"testrepo1": config.Repo{
				URL: "ssh://git@github.com/libgit2/git2go.git",
				Path: testpath+"/Test1",
				Branch: "master",
			},
			"testrepo2": config.Repo{
				URL: "ssh://git@github.com/libgit2/git2go.git",
				Path: testpath+"Test2",
				Branch: "master",
			},

		},
		Signature: config.Signature{
			Name: "Testy Testerson",
			Email: "test@test.com",
		},
	}

	// one repo
	repotsar := RepoTsar{
		Config: config,
		Branch: "test",
		ReposList: []string{"testrepo1"},
		Signature: &git.Signature{
			Name: config.Signature.Name,
			Email: config.Signature.Email,
			When: time.Now(),
		},
	}
	err = repotsar.Run()
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(testpath)

	// no reposlist
	repotsar = RepoTsar{
		Config: config,
		Branch: "test",
		ReposList: []string{""},
		Signature: &git.Signature{
			Name: config.Signature.Name,
			Email: config.Signature.Email,
			When: time.Now(),
		},
	}
	err = repotsar.Run()
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(testpath)

	// no branch
	// one repo
	repotsar = RepoTsar{
		Config: config,
		Branch: "",
		ReposList: []string{"testrepo1"},
		Signature: &git.Signature{
			Name: config.Signature.Name,
			Email: config.Signature.Email,
			When: time.Now(),
		},
	}
	err = repotsar.Run()
	if err != nil {
		t.Error(err)
	}

	// repo already exists
	repotsar = RepoTsar{
		Config: config,
		Branch: "",
		ReposList: []string{"testrepo1"},
		Signature: &git.Signature{
			Name: config.Signature.Name,
			Email: config.Signature.Email,
			When: time.Now(),
		},
	}
	err = repotsar.Run()
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(testpath)

	// repo already exists
	repotsar = RepoTsar{
		Config: config,
		Branch: "",
		ReposList: []string{"bogusrepo"},
		Signature: &git.Signature{
			Name: config.Signature.Name,
			Email: config.Signature.Email,
			When: time.Now(),
		},
	}
	err = repotsar.Run()
	if err == nil {
		t.Error("Repo doesn't exist, should have generated an error and didn't")
	}
	os.RemoveAll(testpath)


}