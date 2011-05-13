package git

/*
#cgo LDFLAGS: -lgit2 -lcrypto -ldl -lz 
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

var (
	println = fmt.Println
	printf  = fmt.Println
)

type Test *C.git_repository

/* Repo */
type Repo struct {
	git_repo *C.git_repository
}

func (v *Repo) Open(path string) (err os.Error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ecode := C.git_repository_open(&v.git_repo, cpath)
	if ecode != GIT_SUCCESS {
		return LastError()
	}
	return
}

func (v *Repo) Free() {
	C.git_repository_free(v.git_repo)
}

func (v *Repo) Init(path string, isbare uint8) (err os.Error) {
	ecode := C.git_repository_init(&v.git_repo, C.CString(path), C.uint(isbare))
	if ecode < GIT_SUCCESS {
		return LastError()
	}
	return
}
// Tree

type Tree struct {
	git_tree *C.git_tree
}

func TreeFromIndex(repo *Repo, index *Index) (*Oid, os.Error) {
	oid := NewOid()
	ecode := C.git_tree_create_fromindex(oid.git_oid, index.git_index)
	if ecode < GIT_SUCCESS {
		return nil, LastError()
	}
	return oid, nil
}

// Commit
type Commit struct {
	git_commit *C.git_commit
}

func CommitCreate(repo *Repo, tree, parent *Oid, author, commiter *Signature, message string) os.Error {
	m := C.CString(message)
	oid := NewOid()
	update_ref := C.CString("HEAD")
	ecode := C.git_commit_create(
		oid.git_oid,
		repo.git_repo,
		update_ref,
		author.git_signature,
		commiter.git_signature,
		m,
		tree.git_oid,
		1,
		&parent.git_oid)
	if ecode < GIT_SUCCESS {
		return LastError()
	}
	return nil
}

func (c *Commit) Lookup(r *Repo, o *Oid) (err os.Error) {
	ecode := C.git_commit_lookup(&c.git_commit, r.git_repo, o.git_oid)
	if ecode < GIT_SUCCESS {
		return LastError()
	}
	return err
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

func NewOid() *Oid {
	o := &Oid{new(C.git_oid)}
	return o
}

func NewOidString(s string) (*Oid, os.Error) {
	o := &Oid{new(C.git_oid)}
	if C.git_oid_mkstr(o.git_oid, C.CString(s)) < GIT_SUCCESS {
		return nil, LastError()
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
	rev := new(RevWalk)
	if C.git_revwalk_new(&rev.git_revwalk, repo.git_repo) < GIT_SUCCESS {
		return nil, LastError()
	}
	return rev, nil
}

func (v *RevWalk) Reset() {
	C.git_revwalk_reset(v.git_revwalk)
}

func (v *RevWalk) Push(o *Oid) {
	C.git_revwalk_push(v.git_revwalk, o.git_oid)
}

func (v *RevWalk) Next(o *Oid) (err os.Error) {
	if C.git_revwalk_next(o.git_oid, v.git_revwalk) < GIT_SUCCESS {
		return LastError()
	}
	return err
}

func GetHeadString(repo *Repo) (string, os.Error) {
	ref := new(Reference)
	err := ref.Lookup(repo, "refs/heads/master")
	if err != nil {
		return "", err
	}
	head := ref.GetOid()
	return head.String(), nil
}

func GetHead(repo *Repo) (*Oid, os.Error) {
	ref := new(Reference)
	err := ref.Lookup(repo, "refs/heads/master")
	if err != nil {
		return nil, err
	}
	head := ref.GetOid()
	return head, nil
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
	if ecode < GIT_SUCCESS {
		return LastError()
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

func (v *Index) Open(repo *Repo) (err os.Error) {
	if ecode := C.git_index_open_inrepo(&v.git_index, repo.git_repo); ecode != GIT_SUCCESS {
		return LastError()
	}
	return
}

func (v *Index) Add(file string) (err os.Error) {
	s := C.CString(file)
	defer C.free(unsafe.Pointer(s))
	if C.git_index_add(v.git_index, s, 0) < GIT_SUCCESS {
		return LastError()
	}
	return
}

func (v *Index) Read() (err os.Error) {
	if ecode := C.git_index_read(v.git_index); ecode != GIT_SUCCESS {
		return LastError()
	}
	return
}

func (v *Index) Write() (err os.Error) {
	if ecode := C.git_index_write(v.git_index); ecode != GIT_SUCCESS {
		return LastError()
	}
	return
}

func (v *Index) Free() {
	C.git_index_free(v.git_index)
}

type Signature struct {
	git_signature *C.git_signature
}

// Signature
func NewSignature(name, email string) *Signature {
	n := C.CString(name)
	e := C.CString(email)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(e))
	s := &Signature{C.git_signature_now(n, e)}
	return s
}

//errors.h:GIT_EXTERN(const char *) git_lasterror(void);

// Helper functions
func LastError() os.Error {
	lasterror := C.GoString(C.git_lasterror())
	return os.NewError(lasterror)
}

//Private
func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}
