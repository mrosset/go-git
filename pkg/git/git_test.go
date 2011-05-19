package git

import (
	"exec"
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	repo    *Repo
	revwalk *RevWalk
	tree    *Tree
	path    string
	author  string
	ref     = new(Reference)
)

func init() {
	path = "./tmp"
	author = "foot@bar.com"
}

// Repo
func TestInitBare(t *testing.T) {
	repo = new(Repo)
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
	repo = new(Repo)
	if err := repo.Init(path, NOTBARE); err != nil {
		t.Fatal("Error:", err)
	}
}

func TestOpenNotBare(t *testing.T) {
	err := repo.Open(path + "/.git")
	checkFatal(t, err)
}

//FIXME: figure out how to seed an intial HEAD
func TestSeed(t *testing.T) {
	tmpfile := "README"
	f, err := os.Create(path + "/" + tmpfile)
	_, err = f.WriteString("foo\n")
	f.Close()
	checkFatal(t, err)
	err = run("git add " + tmpfile)
	checkFatal(t, err)
	err = run("git commit -m test")
	checkFatal(t, err)
}

// Index 
func TestIndexAdd(t *testing.T) {
	index := new(Index)
	defer index.Free()
	err := index.Open(repo)
	check(t, err)
	tmpfile := "README"
	f, err := os.OpenFile(path+"/"+tmpfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	check(t, err)
	_, err = f.WriteString("foo\n")
	f.Close()
	err = index.Add(tmpfile)
	check(t, err)
	err = index.Write()
	check(t, err)
}

// Commit
func TestCommit(t *testing.T) {
	TestIndexAdd(t)
	index := new(Index)
	defer index.Free()
	err := index.Open(repo)
	check(t, err)
	tree, err := TreeFromIndex(repo, index)
	check(t, err)
	head, err := GetHead(repo)
	check(t, err)
	s := NewSignature("Foo Bar", "foo@bar.com")
	defer s.Free()
	err = CommitCreate(repo, tree, head, s, s, "some stuff here")
	check(t, err)
}
func TestManyCommits(t *testing.T) {
	for i := 0; i < 29; i++ {
		TestCommit(t)
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
	head, _ := GetHeadString(repo)
	o, _ := NewOidString(head)
	revwalk.Push(o)
	if err := revwalk.Next(o); err != nil {
		t.Fatal("Error:", err)
	}
	if o.String() != head {
		t.Errorf("oid string should match %v is %v", head, o.String())
	}
}

func TestRevWalk(t *testing.T) {
	r := new(Repo)
	defer r.Free()
	err := r.Open("./tmp/.git")
	check(t, err)
	o := NewOid()
	for {
		if err := revwalk.Next(o); err != nil {
			break
		}
		c := new(Commit)
		c.Lookup(repo, o)
		// Output example
		fmt.Printf("%v %v %v %v\n", o.String(), c.Author(), c.Email(), c.Msg())
	}
}

// Oid
func TestNewOid(t *testing.T) {
	head, err := GetHeadString(repo)
	check(t, err)
	if _, err := NewOidString(head); err != nil {
		t.Error(err)
	}
}

// Singature
func TestSignature(t *testing.T) {
	NewSignature("foo", "bar")
}

func TestTSignature(t *testing.T) {
}

// Tree
func TestTreeFromIndex(t *testing.T) {
	index := new(Index)
	defer index.Free()
	err := index.Open(repo)
	check(t, err)
	_, err = TreeFromIndex(repo, index)
	check(t, err)
}

func TestTreeFromCommit(t *testing.T) {
	head, err := GetHead(repo)
	check(t, err)
	commit := new(Commit)
	err = commit.Lookup(repo, head)
	check(t, err)
	tree, err = TreeFromCommit(repo, commit)
	check(t, err)
}

func TestTreeEntryByName(t *testing.T) {
	expected := "README"
	entry, err := tree.EntryByName(expected)
	if err != nil {
		t.Fatal("Expected to find a file, but was unable to.")
	}
	if entry.Filename() != expected {
		t.Fatal("EntryByName did not return the proper file. Expected %v, got %v",
			expected,
			entry.Filename())
	}
}


func TestInvalidTreeEntryByName(t *testing.T) {
	expected := "README.does-not-exist"
	_, err := tree.EntryByName(expected)
	if err == nil {
		t.Fatal("Was expecting a does not exist error, but did not recieve one.")
	}
}

func TestTreeEntryByIndex(t *testing.T) {
	expected := "README"
	entry, err := tree.EntryByIndex(0)
	if err != nil {
		t.Fatal("Was unable to find the first entry via index.")
	}
	if entry.Filename() != expected {
		t.Fatalf(
			"EntryByIndex did not return the proper file. Expected %v, got %v",
			expected,
			entry.Filename())
	}
}

func TestTreeEntryCount(t *testing.T) {
	if tree.EntryCount() != 1 {
		t.Fatalf("Expected 1 file in the tree, but there were %v", tree.EntryCount())
	}
}

// Important: this must be called after all of the Test functions
func TestFinal(t *testing.T) {
	if tree != nil {
		tree.Free()
	}
	if revwalk != nil {
		revwalk.Free()
	}
	if repo != nil {
		repo.Free()
	}
}

// private helper functions
func run(s string) (err os.Error) {
	wd := path
	args := strings.Split(s, " ", -1)
	bin, err := exec.LookPath(args[0])

	cmd, err := exec.Run(bin, args, os.Environ(), wd, exec.DevNull, exec.Pipe, exec.PassThrough)
	if err != nil {
		return err
	}
	w, err := cmd.Wait(0)
	if err != nil {
		return err
	}
	if !w.Exited() || w.ExitStatus() != 0 {
		return os.NewError("failed to run " + s)
	}
	return
}

func checkFatal(t *testing.T, err os.Error) {
	if err != nil {
		fmt.Printf("Fatal: %T %v\n", t, err)
		os.Exit(0)
	}
}

func check(t *testing.T, err os.Error) {
	if err != nil {
		t.Error(err)
	}
}
