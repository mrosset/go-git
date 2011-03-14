package git

import (
	"exec"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	repo    *Repo
	revwalk *RevWalk
	path    string
	head    string
	ref     = &Reference{}
)

func init() {
	path = "./tmp"
}

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
		err os.Error
		cmd *exec.Cmd
	)

	tmpfile := "README"

	err = ioutil.WriteFile("./tmp/"+tmpfile, []byte("test\n"), 0644)

	cmd, err = run("git add " + tmpfile)
	cmd.Close()
	cmd, err = run("git commit -m test")
	cmd.Close()
	cmd, err = run("git log --pretty=%H")
	buf, _ := ioutil.ReadAll(cmd.Stdout)
	head = strings.Trim(string(buf), "\n")
	println(head)
	cmd.Close()

	if err != nil {
		t.Fatal("Error:", err)
	}
}

func run(s string) (cmd *exec.Cmd, err os.Error) {
	wd := "./tmp/"
	args := strings.Split(s, " ", -1)
	bin, err := exec.LookPath(args[0])

	cmd, err = exec.Run(bin, args, os.Environ(), wd, exec.DevNull, exec.Pipe, exec.PassThrough)

	return
}

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

/*
func TestRefGetOid(t *testing.T) {
	oid := ref.GetOid()
	oid.String()
}
*/
//Important: this must be called after all of the Test functions
func TestFinal(t *testing.T) {
	if revwalk != nil {
		revwalk.Free()
	}
	if repo != nil {
		repo.Free()
	}
}

/*  Not implimented yet 

func TestLookup(t *testing.T) {
	c := &Commit{}
	o, _ := NewOid(oid)
	repo.Lookup(c, o, GIT_OBJ_ANY)
	if c.Author() != author {
		t.Fatal("ERROR:",os.NewError("Lookup failed"))
	}
}

func TestNewOid(t *testing.T) {
	if _, err := NewOid(oid); err != nil {
		t.Error(err)
	}
}

func TestNewRevWalk(t *testing.T) {
	var err os.Error
	revwalk, err = NewRevWalk(repo)
	if err != nil {
		t.Fatal("ERROR:", err)
	}
}
*/


func TestTest(t *testing.T) {
	test()
}
