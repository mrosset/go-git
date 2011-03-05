#include <git2.h>
#include <stdio.h>
#include <string.h>

#define BARE 1
#define NOTBARE 0

int main() {
    git_repository *repo;
    int error;
    const char *path = "/home/strings/git/go-git/tmp";

    error = git_repository_init(&repo,path,NOTBARE);

    if (error != 0) {
        printf("ERROR: init %s code %i\n",path,BARE); 
        return(1);
    };
    
    error = git_repository_open(&repo, path); 
    if (error != 0) {
        printf("ERROR: opening %s code %i\n",path,error); 
        return(1);
    };

    /* do stuff with the repository */
    git_repository_free(repo);
    return(0);
}
