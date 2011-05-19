// godefs -g git defs.c

// MACHINE GENERATED - DO NOT EDIT.

package git

// Constants
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
	GIT_REF_SYMBOLIC  = 0x2
	GIT_SUCCESS       = 0
	GIT_ENOTOID       = -0x2
)

// Types

type GitTime struct {
	Time         int64
	Offset       int32
	Pad_godefs_0 [4]byte
}
