PACKAGE

package git
import "."

CONSTANTS

const (
    GIT_OBJ_ANY       = -0x2
    GIT_OBJ_BAD       = -0x1
    GIT_OBJ__EXT1     = 0
    GIT_OBJ_BLOB      = 0x3
    GIT_OBJ_COMMIT    = 0x1
    GIT_OBJ_OFS_DELTA = 0x6
    GIT_OBJ_REF_DELTA = 0x7
    GIT_OBJ_TAG       = 0x4
    GIT_OBJ_TREE      = 0x2
    GIT_OBJ__EXT2     = 0x5
    GIT_SUCCESS       = 0
    GIT_ENOTOID       = -0x2
)
Constants

const (
    NOTBARE = iota
    BARE
)


FUNCTIONS

func CommitCreate(repo *Repo, tree, parent *Oid, author, commiter *Signature, message string) os.Error

func GetHeadString(repo *Repo) (string, os.Error)

func LastError() os.Error
Helper functions


TYPES

type Commit struct {
    // contains filtered or unexported fields
}
Commit

func (c *Commit) Author() string

func (c *Commit) Email() string

func (c *Commit) Lookup(r *Repo, o *Oid) (err os.Error)

func (c *Commit) Msg() string

type GitTime struct {
    Time   int64
    Offset int32
}

type Index struct {
    // contains filtered or unexported fields
}
Index

func (v *Index) Add(file string) (err os.Error)

func (v *Index) Free()

func (v *Index) Open(repo *Repo) (err os.Error)

func (v *Index) Read() (err os.Error)

func (v *Index) Write() (err os.Error)

type IndexEntry struct {
    Ctime          [12]byte /* git_index_time */
    Mtime          [12]byte /* git_index_time */
    Dev            uint32
    Ino            uint32
    Mode           uint32
    Uid            uint32
    Gid            uint32
    File_size      int64
    Oid            [20]byte /* git_oid */
    Flags          uint16
    Flags_extended uint16
    Path           *int8
}

type Oid struct {
    // contains filtered or unexported fields
}
Oid

func GetHead(repo *Repo) (*Oid, os.Error)

func NewOid() *Oid

func NewOidString(s string) (*Oid, os.Error)

func TreeFromIndex(repo *Repo, index *Index) (*Oid, os.Error)

func (v *Oid) String() string

type Reference struct {
    // contains filtered or unexported fields
}
Reference

func (v *Reference) GetOid() *Oid

func (v *Reference) Lookup(r *Repo, name string) (err os.Error)

type Repo struct {
    // contains filtered or unexported fields
}
Repo

func (v *Repo) Free()

func (v *Repo) Init(path string, isbare uint8) (err os.Error)

func (v *Repo) Open(path string) (err os.Error)

type RevWalk struct {
    // contains filtered or unexported fields
}
RevWalk

func NewRevWalk(repo *Repo) (*RevWalk, os.Error)

func (v *RevWalk) Free()

func (v *RevWalk) Next(o *Oid) (err os.Error)

func (v *RevWalk) Push(o *Oid)

func (v *RevWalk) Reset()

func (v *RevWalk) Sorting(sm uint)
TODO: implement this

type Signature struct {
    // contains filtered or unexported fields
}
Signature

func NewSignature(name, email string) *Signature

type Test *C.git_repository

type Tree struct {
    // contains filtered or unexported fields
}
Tree

