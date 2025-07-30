package tests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	g "github.io/iamthen0ise/bb/src/git"
)

func setupRepo(t *testing.T) string {
	dir := t.TempDir()
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git init failed: %v, %s", err, string(out))
	}
	// create initial commit so branch commands work
	if err := os.WriteFile(dir+"/README.md", []byte("init"), 0644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}
	cmd = exec.Command("git", "add", "README.md")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git add failed: %v, %s", err, string(out))
	}
	cmd = exec.Command("git", "commit", "-m", "init")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git commit failed: %v, %s", err, string(out))
	}
	return dir
}

func TestCreateNewBranch(t *testing.T) {
	repo := setupRepo(t)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	defer os.Chdir(cwd)
	if err := os.Chdir(repo); err != nil {
		t.Fatalf("chdir failed: %v", err)
	}

	if err := g.CreateNewBranch("testbranch", false); err != nil {
		t.Fatalf("create branch failed: %v", err)
	}

	cmd := exec.Command("git", "branch", "--list", "testbranch")
	cmd.Dir = repo
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("list branch failed: %v, %s", err, string(out))
	}
	if strings.TrimSpace(string(out)) != "testbranch" {
		t.Fatalf("branch not created, got: %s", string(out))
	}
}
