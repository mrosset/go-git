#include <git2.h>
#include <stdio.h>

#define BARE 1
#define NOTBARE 0


int test_repo(const char *path, unsigned type);

int main() 
{
    const char *tmpdir = "./tmp\0";
   
    printf("beginning test\n");
    printf("bare");
    if (test_repo(tmpdir, BARE) != 0) {
        return 1;
    }
    printf("not bare");
    if (test_repo(tmpdir, NOTBARE) != 0) {
        return 1;
    }
    printf("done.\n");
    return 0;
}


int test_repo(const char *path, unsigned type) 
{
    git_repository *repo;

    int error = GIT_SUCCESS;
    error = git_repository_init(&repo,path,type);

    printf(" init");
    if (error != 0) {
        printf("ERROR: init %s code %i\n",path,error); 
        return error;
    };
   
    error = git_repository_open(&repo, path); 
    printf(" open");
    if (error != 0) {
        printf("ERROR: opening %s code %i\n",path,error); 
        return error;
    };
    printf(" pass\n");
    git_repository_free(repo);
    return error; 
}
