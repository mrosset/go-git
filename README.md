go-git
=================
Go bindings to [libgit2](http://libgit2.github.com/)

Status
=================
Currently in active development.

Documentation
-----------------
http://str1ngs.github.com/go-git


Installation
------------

    sudo apt-get install libssl-dev # for libcypto dependency
    hub clone str1ngs/go-git # if you don't have hub, see defunkt/hub
                             # or use vanilla git.
    pushd go-git
    make module
    pushd libgit2
    ./waf configure
    ./waf build
    sudo ./waf install
    popd && popd
    goinstall github.com/str1ngs/go-git/pkg/git
