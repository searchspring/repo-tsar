package gitutils

import (
	"io/ioutil"
	"testing"
	"time"

	"gopkg.in/libgit2/git2go.v22"
)

func createTestRepo(t *testing.T) *git.Repository {
	// figure out where we can create the test repo
	path, err := ioutil.TempDir("", "RepoTsar")
	if err != nil {
		t.Error(err)
	}

	repo, err := git.InitRepository(path, false)
	if err != nil {
		t.Error(err)
	}
	tmpfile := "README"
	err = ioutil.WriteFile(path+"/"+tmpfile, []byte("foo\n"), 0644)

	if err != nil {
		t.Error(err)
	}
	return repo
}

func seedTestRepo(t *testing.T, repo *git.Repository) (*git.Oid, *git.Oid) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Error(err)
	}
	sig := &git.Signature{
		Name:  "Rand Om Hacker",
		Email: "random@hacker.com",
		When:  time.Date(2013, 03, 06, 14, 30, 0, 0, loc),
	}

	idx, err := repo.Index()
	if err != nil {
		t.Error(err)
	}
	err = idx.AddByPath("README")
	if err != nil {
		t.Error(err)
	}
	treeId, err := idx.WriteTree()
	if err != nil {
		t.Error(err)
	}

	message := "This is a commit\n"
	tree, err := repo.LookupTree(treeId)
	if err != nil {
		t.Error(err)
	}

	commitId, err := repo.CreateCommit("HEAD", sig, sig, message, tree)
	if err != nil {
		t.Error(err)
	}
	return commitId, treeId
}