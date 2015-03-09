package fileutils

import(
	"testing"
	"os"
	"io/ioutil"

)

func TestCreatePath(t *testing.T) {
	testpath, err := ioutil.TempDir("", "RepoTsar")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)
	
	// Test path already exists
	path,err := CreatePath(testpath)
	if err != nil {
		t.Error(err)
	}
	if path != testpath {
		t.Errorf("Expected %#v, got %#v", testpath, path)
	}
	// Test path doesn't exist
	path,err = CreatePath(testpath+"/Testdir")
	if err != nil {
		t.Error(err)
	}
	if path != testpath+"/Testdir" {
		t.Errorf("Expected %#v, got %#v", testpath+"/Testdir", path)
	}

	// Test expanding ~
	homedir := os.Getenv("HOME")
	os.Setenv("HOME",testpath)
	defer os.Setenv("HOME",homedir)
	path,err = CreatePath("~/Testdir2")
	if err != nil {
		t.Error(err)
	}
	if path != testpath+"/Testdir2" {
		t.Errorf("Expected %#v, got %#v", testpath+"/Testdir2", path)
	}
}