package main

import (
	"flag"
	"fmt"
	git "github.com/str1ngs/go-git/pkg/git"
	"os"
	"path"
)

var (
	test    = flag.Bool("t", false, "test")
	printf  = fmt.Printf
	println = fmt.Println
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	if *test {
		doTest()
	}
}

func doTest() {
	var (
		tmpdir  = "./tmp"
		tmpfile = path.Join(tmpdir, "README")
		repo    = &git.Repo{}
	)
	defer repo.Free()
	err := repo.Open(tmpdir + "/.git")
	if err != nil {
		println(err)
		return
	}
	println("opening", tmpdir)
	f, err := os.Create(tmpfile)
	if err != nil {
		println(err)
		return
	}
	println("creating", tmpfile)
	f.WriteString("foo\n")
	println("writing", tmpfile)
	f.Close()
	index := new(git.Index)
	defer index.Free()
	err = index.Open(repo)
	if err != nil {
		println(err)
		return
	}
	err = index.Add("README")
	if err != nil {
		println(err)
		return
	}
	err = index.Write()
	if err != nil {
		println(err)
		return
	}
	println("added", tmpfile, "to index")

	tree, err := git.TreeFromIndex(repo, index)
	if err != nil {
		println(err)
		return
	}
	parent, err := git.GetHead(repo)
	if err != nil {
		println(err)
		return
	}
	s := git.NewSignature("Foo Bar", "foo@bar.com")
	err = git.CommitCreate(repo, tree, parent, s, s, "test commit")
	if err != nil {
		println(err)
		return
	}
	head, err := git.GetHeadString(repo)
	if err != nil {
		println(err)
		return
	}
	println("commited", head)
}
