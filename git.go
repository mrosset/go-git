package git

/*
#include <stdlib.h>
#include <git2.h>
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

const (
	NOTBARE = iota
	BARE
)

/* Repo */
type Repo struct {
	git_repo *C.git_repository
}

func (v *Repo) Open(path string) (err os.Error) {
	ecode := C.git_repository_open(&v.git_repo, C.CString(path))
	if ecode != GIT_SUCCESS {
		err = os.NewError(fmt.Sprintf("failed to open %v CODE %v", path, ecode))
	}
	return
}

func (v *Repo) Lookup(c *Commit, o *Oid, mask int) {
	//C.git_repository_lookup(&c.git_commit, v.git_repo, o.git_oid, (C.git_otype)(mask))
}

func (v *Repo) Free() {
	C.git_repository_free(v.git_repo)
}

func (v *Repo) Init(path string, isbare uint8) (err os.Error) {
	ecode := C.git_repository_init(&v.git_repo, C.CString(path), C.uint(isbare))
	if ecode != GIT_SUCCESS {
		e := fmt.Sprintf("failed to init %v CODE %v", path, ecode)
		return os.NewError(e)
	}
	return
}

// Commit
type Commit struct {
	git_commit *C.git_commit
}

func (c *Commit) Msg() string {
	msg := C.git_commit_message(c.git_commit)
	return C.GoString(msg)
}

func (c *Commit) Author() string {
	p := C.git_commit_author(c.git_commit)
	return C.GoString(p.name)
}
func (c *Commit) Email() string {
	p := C.git_commit_author(c.git_commit)
	return C.GoString(p.email)
}

// Oid
type Oid struct {
	git_oid *C.git_oid
}

func NewOid(s string) (*Oid, os.Error) {
	o := &Oid{new(C.git_oid)}
	if C.git_oid_mkstr(o.git_oid, C.CString(s)) == GIT_ENOTOID {
		return nil, os.NewError("could not create new oid")
	}
	return o, nil
}

func (v *Oid) String() string {
	p := C.git_oid_allocfmt(v.git_oid)
	sha := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return sha
}

// RevWalk
type RevWalk struct {
	git_revwalk *C.git_revwalk
}

func NewRevWalk(repo *Repo) (*RevWalk, os.Error) {
	r := &RevWalk{}
	if C.git_revwalk_new(&r.git_revwalk, repo.git_repo) != 0 {
		return nil, os.NewError("could not create new RevWalk")
	}
	return r, nil
}

func (v *RevWalk) Reset() {
	C.git_revwalk_reset(v.git_revwalk)
}

func (v *RevWalk) Push(c *Commit) {
	C.git_revwalk_push(v.git_revwalk, c.git_commit)
}

func (v *RevWalk) Next(c *Commit) {
	C.git_revwalk_next(&c.git_commit, v.git_revwalk)
}

//TODO: implement this
func (v *RevWalk) Sorting(sm uint) {
}

func (v *RevWalk) Free() {
	C.git_revwalk_free(v.git_revwalk)
}

//Reference

type Reference struct {
	git_reference *C.git_reference
}

func (v *Reference) Lookup(r *Repo, name string) (err os.Error) {
	ecode := C.git_reference_lookup(&v.git_reference, r.git_repo, C.CString(name))
	if ecode != GIT_SUCCESS {
		err = os.NewError(fmt.Sprintf("failed to Lookup %v CODE %v", name, ecode))
	}
	return
}

func (v *Reference) GetOid() *Oid {
	return &Oid{C.git_reference_oid(v.git_reference)}
}

//Index

type Index struct {
	git_index *C.git_index
}

func (v *Index) Read() (err os.Error) {
	if ecode := C.git_index_read(v.git_index); ecode != GIT_SUCCESS {
		return os.NewError("index read failed with code " + string(ecode))
	}
	return
}

//Private
func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}

//Test foo functions

func test() {
}
