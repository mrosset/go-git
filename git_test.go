package git

import "os"
import "testing"

const (
	oid    = "0b1f85e2daba25df3b2070b554c6464c780d0f9b"
	author = "str1ngs"
)

var (
	repo    *Repo
	revwalk *RevWalk
	path    string
	calls   int = 0
)

func init() {
	path, _ = os.Getwd()
	path = "/tmp"
	repo = &Repo{}
}

func TestInit(t *testing.T) {
	if err := repo.Init(path, BARE); err != nil {
		t.Fatal("Error:", err)
	}
}

func TestOpen(t *testing.T) {
	err := repo.Open(path)
	if err != nil {
		t.Fatal("Error:", err)
	}
}

func TestNewOid(t *testing.T) {
	if _, err := NewOid(oid); err != nil {
		t.Error(err)
	}
}

/*
func TestLookup(t *testing.T) {
	c := &Commit{}
	o, _ := NewOid(oid)
	repo.Lookup(c, o, GIT_OBJ_ANY)
	if c.Author() != author {
		t.Fatal("ERROR:",os.NewError("Lookup failed"))
	}
}
*/

func TestNewRevWalk(t *testing.T) {
	var err os.Error
	revwalk, err = NewRevWalk(repo)
	if err != nil {
		t.Fatal("ERROR:", err)
	}
}

//Important: this must be called after all of the Test functions
func TestFinal(t *testing.T) {
	revwalk.Free()
	repo.Free()
}

func TestTest(t *testing.T) {
	test()
}
