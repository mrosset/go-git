#include <git2.h>
#include <stdio.h>

#define BARE 1
#define NOTBARE 0


int test_repo(git_repository **repo, const char *path, unsigned type);

int main() 
{
    git_repository *repo;
    const char *tmpdir = "./tmp\0";
   
    printf("beginning test\n");

    printf("bare\t\t");
    test_repo(&repo, tmpdir, BARE);

    printf("not bare\t");
    test_repo(&repo, tmpdir, NOTBARE);

    printf("done.\n");

    git_repository_free(repo);
    return 0;
}


int test_repo(git_repository **repo, const char *path, unsigned type) 
{
    int error = GIT_SUCCESS;
    error = git_repository_init(&repo,path,type);

    if (error != 0) {
        printf("ERROR: init %s code %i\n",path,error); 
        return error;
    };
   
    error = git_repository_open(&repo, path); 

    if (error != 0) {
        printf("ERROR: opening %s code %i\n",path,error); 
        return error;
    };
    printf(" pass\n");
    return error; 
}
