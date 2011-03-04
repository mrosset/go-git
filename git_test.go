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
	fatal   bool = false
	calls   int  = 0
)

func init() {
	path, _ = os.Getwd()
	path += "/tmp"
	repo = &Repo{}
}

func TestInit(t *testing.T) {
	if fatal {
		return
	}
	if err := repo.Init(path, NOTBARE); err != nil {
		fatal = true
		t.Error(err)
	}
}

func TestOpen(t *testing.T) {
	if fatal {
		return
	}
	err := repo.Open(path + "/.git")
	if err != nil {
		fatal = true
		println("We can not test with out a working Repo")
		println(err.String())
	}
}

func TestNewOid(t *testing.T) {
	if fatal {
		return
	}
	if _, err := NewOid(oid); err != nil {
		fatal = true
		t.Error(err)
	}
}

func TestLookup(t *testing.T) {
	if fatal {
		return
	}
	c := &Commit{}
	o, _ := NewOid(oid)
	repo.Lookup(c, o, GIT_OBJ_ANY)
	if c.Author() != author {
		fatal = true
		t.Error(os.NewError("Lookup failed"))
	}
}

func TestNewRevWalk(t *testing.T) {
	var err os.Error
	if fatal {
		return
	}
	revwalk, err = NewRevWalk(repo)
	if err != nil {
		fatal = true
		t.Error(err)
	}
}

//Important: this must be called after all of the Test functions
func TestFinal(t *testing.T) {
	if fatal {
		return
	}
	revwalk.Free()
	repo.Free()
}

func TestTest(t *testing.T) {

	test()
}
