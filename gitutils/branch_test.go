package gitutils

import(
	"testing"
	"os"
	"time"

	"gopkg.in/libgit2/git2go.v22"
)


func TestBranchInfoStruct(t *testing.T){
	repo := createTestRepo(t)
	seedTestRepo(t,repo)
	defer os.RemoveAll(repo.Workdir())
	defer repo.Free()

	signature := &git.Signature{
		Name: "Test Name",
		Email: "test@email.com",
		When: time.Now(),
	}

	branchinfo := BranchInfo{
		Reponame: "test",
		Branchname: "testbranch",
		Msg: "RepoTsar Branching",
		Repo: *repo,
		Signature: signature,
	}
	
	for i := 0; i < 2; i++ {
		err := branchinfo.GitBranch()
		if err != nil {
			t.Error(err)
		}
	}

}