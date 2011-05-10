package git

import (
	"exec"
	"encoding/line"
	"os"
	"strings"
	"testing"
)

var (
	repo    *Repo
	revwalk *RevWalk
	path    string
	head    string
	author  string
	ref     = &Reference{}
)

func init() {
	path = "./tmp"
	author = "schizoid29@gmail.com"
}

// Repo
func TestInitBare(t *testing.T) {
	repo = &Repo{}
	if err := repo.Init(path, BARE); err != nil {
		t.Fatal("Error:", err)
	}
}

func TestOpenBare(t *testing.T) {
	defer os.RemoveAll(path)
	defer repo.Free()
	err := repo.Open(path)
	if err != nil {
		t.Fatal("Error:", err)
	}
}

func TestInitNotBare(t *testing.T) {
	repo = &Repo{}
	if err := repo.Init(path, NOTBARE); err != nil {
		t.Fatal("Error:", err)
	}
}

func TestOpenNotBare(t *testing.T) {
	err := repo.Open(path + "/.git")
	if err != nil {
		t.Fatal("Error:", err)
	}
}

//FIXME: only fork calls till we have proper boiler plate
func TestCommit(t *testing.T) {
	var (
		cmd *exec.Cmd
	)

	tmpfile := "README"

	f, err := os.Create(path+"/"+tmpfile)
	_, err = f.WriteString("foo\n")
	f.Close()
	if err != nil {
		println(err.String())
		os.Exit(1)
	}
	cmd, err = run("git add " + tmpfile)
	cmd.Close()
	cmd, err = run("git commit -m test")
	cmd.Close()
	cmd, err = run("git log --pretty=%H")
	r := line.NewReader(cmd.Stdout, 256)
	h, _, _ := r.ReadLine()
	head = (string(h))
	cmd.Close()

	if err != nil {
		t.Fatal("Error:", err)
	}
}

//FIXME: change this to a bench and not use forks
func TestManyCommits(t *testing.T) {
	for i := 0; i < 29; i++ {
		TestCommit(t)
	}
}

// Commit

// Ref
func TestRefLookup(t *testing.T) {
	var err os.Error
	refpath := "refs/heads/master"
	_, err = os.Stat(".git/" + refpath)
	if err != nil {
		t.Fatalf("Error:", err)
	}
	err = ref.Lookup(repo, "refs/heads/master")
	if err != nil {
		t.Fatalf("Error:", err)
	}
	if ref.GetOid().String() != head {
		t.Fatalf("Error:", os.NewError("sha does not match head"))
	}
}

// RevWalk
func TestNewRevWalk(t *testing.T) {
	var err os.Error
	revwalk, err = NewRevWalk(repo)
	if err != nil {
		t.Fatal("Error:", err)
	}
}

func TestRevWalkNext(t *testing.T) {
	o, _ := NewOid(head)
	revwalk.Push(o)
	if err := revwalk.Next(o); err != nil {
		t.Fatal("Error:", err)
	}
	if o.String() != head {
		t.Errorf("oid string should match %v is %v", head, o.String())
	}
}

// Oid
func TestNewOid(t *testing.T) {
	if _, err := NewOid(head); err != nil {
		t.Error(err)
	}
}

// Important: this must be called after all of the Test functions
func TestFinal(t *testing.T) {
	if revwalk != nil {
		revwalk.Free()
	}
	if repo != nil {
		repo.Free()
	}
}

// private helper functions
func run(s string) (cmd *exec.Cmd, err os.Error) {
	wd := "./tmp/"
	args := strings.Split(s, " ", -1)
	bin, err := exec.LookPath(args[0])

	cmd, err = exec.Run(bin, args, os.Environ(), wd, exec.DevNull, exec.Pipe, exec.PassThrough)

	return
}
