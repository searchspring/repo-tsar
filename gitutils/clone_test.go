package gitutils

import(
	"testing"
	"os"
	"io/ioutil"

)

func TestCloneInfoStruct(t *testing.T){
	path, err := ioutil.TempDir("", "RepoTsar")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(path)

	for i := 0; i < 2; i++ {
		cloneinfo := &CloneInfo {
			Reponame: "TestRepo",
			Path: path,
			URL: "ssh://git@github.com/libgit2/git2go.git",
			Branch: "master",
		}
		_,err = cloneinfo.CloneRepo()
		if err != nil {
			t.Error(err)
		}
	}



}