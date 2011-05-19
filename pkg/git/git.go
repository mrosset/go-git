package git

/*
#cgo CFLAGS: -O2 -pipe -march=native -mtune=native
#cgo LDFLAGS: -lgit2 -lcrypto -ldl -lz 
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

type Test *C.git_repository

// Repo
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

func (t *Tree) Free() {
	C.git_tree_close(t.git_tree)
}

func TreeLookup(repo *Repo, oid *Oid) (*Tree, os.Error) {
	tree := new(Tree)
	ecode := C.git_tree_lookup(&tree.git_tree, repo.git_repo, oid.git_oid)
	if ecode < GIT_SUCCESS {
		return nil, LastError()
	}
	return tree, nil
}

func TreeFromIndex(repo *Repo, index *Index) (*Oid, os.Error) {
	oid := NewOid()
	ecode := C.git_tree_create_fromindex(oid.git_oid, index.git_index)
	if ecode < GIT_SUCCESS {
		return nil, LastError()
	}
	return oid, nil
}

func TreeFromCommit(repo *Repo, commit *Commit) (*Tree, os.Error) {
	tree := new(Tree)
	ecode := C.git_commit_tree(&tree.git_tree, commit.git_commit)
	if ecode < GIT_SUCCESS {
		return nil, LastError()
	}
	return tree, nil
}

func (t *Tree) EntryByName(filename string) (*Entry, os.Error) {
	entry := new(Entry)
	entry.git_tree_entry = C.git_tree_entry_byname(t.git_tree, C.CString(filename))
	if entry.git_tree_entry == nil {
		return nil, os.NewError("Unable to find entry.")
	}
	return entry, nil
}

func (t *Tree) EntryByIndex(index int) (*Entry, os.Error) {
	entry := new(Entry)
	entry.git_tree_entry = C.git_tree_entry_byindex(t.git_tree, C.int(index))
	if entry.git_tree_entry == nil {
		return nil, os.NewError("Unable to find entry.")
	}
	return entry, nil
}

func (t *Tree) EntryCount() int {
	num := C.git_tree_entrycount(t.git_tree)
	return int(num)
}

// Entry
type Entry struct {
	git_tree_entry *C.git_tree_entry
}

func (e *Entry) Oid() *Oid {
	return &Oid{C.git_tree_entry_id(e.git_tree_entry)}
}

func (e *Entry) Filename() string {
	filename := C.git_tree_entry_name(e.git_tree_entry)
	return C.GoString(filename)
}

// Commit
type Commit struct {
	git_commit *C.git_commit
}

//TODO: do not use hardcoded update_ref
func CommitCreate(repo *Repo, tree, parent *Oid, author, commiter *Signature, message string) os.Error {
	oid := NewOid()
	m := C.CString(message)
	defer C.free(unsafe.Pointer(m))
	update_ref := C.CString("HEAD")
	defer C.free(unsafe.Pointer(update_ref))
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
	return &Oid{new(C.git_oid)}
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
	sha1 := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return sha1
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

//TODO: end of walk produces a LastError figure out how to reset it and return a os.EOF instead
// possibly return a Oid? Would make for less boiler plate.
func (v *RevWalk) Next(o *Oid) (err os.Error) {
	if C.git_revwalk_next(o.git_oid, v.git_revwalk) < GIT_SUCCESS {
		return LastError()
	}
	return err
}

//TODO: do not assume we are working on refs/heads/master
func GetHead(repo *Repo) (*Oid, os.Error) {
	ref := new(Reference)
	err := ref.Lookup(repo, "refs/heads/master")
	if err != nil {
		return nil, err
	}
	head, err := ref.GetOid()
	if err != nil {
		return nil, err
	}
	return head, nil
}

func GetHeadString(repo *Repo) (string, os.Error) {
	head, err := GetHead(repo)
	if err != nil {
		return "", err
	}
	return head.String(), err
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
	if v.git_reference == nil {
		//return os.NewError("Reference Lookup: Failed to look up " + name)
	}
	return
}

func (v *Reference) SetTarget(target string) (err os.Error) {
	ctarget := C.CString(target)
	defer C.free(unsafe.Pointer(ctarget))
	ecode := C.git_reference_set_target(v.git_reference, ctarget)
	if ecode < GIT_SUCCESS {
		return LastError()
	}
	return nil
}

func (v *Reference) Type() {
	if C.git_reference_type(v.git_reference) == GIT_REF_SYMBOLIC {
		println("THIS IS A SYMBOLIC REF")
	}
}

func (v *Reference) GetOid() (*Oid, os.Error) {
	oid := C.git_reference_oid(v.git_reference)
	if oid == nil {
		//return nil, os.NewError("GetOid Failed: unable to get Oid for reference")
	}
	return &Oid{oid}, nil
}

//Index
type Index struct {
	git_index *C.git_index
}

func (v *Index) Open(repo *Repo) (err os.Error) {
	if ecode := C.git_index_open_inrepo(&v.git_index, repo.git_repo); ecode != GIT_SUCCESS {
		estring := fmt.Sprintf("failed to open index error code %v", ecode)
		return os.NewError(estring)
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

func (v *Index) EntryCount() int {
	return int(C.git_index_entrycount(v.git_index))
}

func (v *Index) Free() {
	C.git_index_free(v.git_index)
}

//TODO: its possible we can use godef to generate this struct
// Signature
type Signature struct {
	git_signature *C.git_signature
}

func (s Signature) Free() {
	C.git_signature_free(s.git_signature)
}

func NewSignature(name, email string) *Signature {
	n := C.CString(name)
	e := C.CString(email)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(e))
	s := &Signature{C.git_signature_now(n, e)}
	return s
}

// Helper functions
func LastError() os.Error {
	lasterror := C.GoString(C.git_lasterror())
	return os.NewError(lasterror)
}

//Private
func printT(i interface{}) {
	fmt.Printf("%T = %v\n", i, i)
}

func abort() {
	fmt.Println("GOT ****************************** HERE")
	os.Exit(0)
}
