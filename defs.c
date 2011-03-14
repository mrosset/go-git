#include <git2.h>

enum
{
	$GIT_OBJ_ANY = GIT_OBJ_ANY,
	$GIT_OBJ_BAD = GIT_OBJ_BAD,
	$GIT_OBJ__EXT1 = GIT_OBJ__EXT1,
    $GIT_OBJ_BLOB = GIT_OBJ_BLOB,     
    $GIT_OBJ_COMMIT = GIT_OBJ_COMMIT,
    $GIT_OBJ_OFS_DELTA = GIT_OBJ_OFS_DELTA,
    $GIT_OBJ_REF_DELTA = GIT_OBJ_REF_DELTA,
    $GIT_OBJ_TAG = GIT_OBJ_TAG,      
    $GIT_OBJ_TREE = GIT_OBJ_TREE,     
    $GIT_OBJ__EXT2 = GIT_OBJ__EXT2,    

    $GIT_SUCCESS = GIT_SUCCESS,
    $GIT_ENOTOID = GIT_ENOTOID
};

typedef struct git_signature $Signature;
typedef struct git_time $GitTime;
typedef struct git_index_entry $IndexEntry;
