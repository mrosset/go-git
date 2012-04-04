go-git
=================
Go bindings to [libgit2](http://libgit2.github.com/)

Status
=================
Currently in active development.

Documentation
-----------------
http://gopkgdoc.appspot.com/pkg/github.com/str1ngs/go-git


Requirements
-----------------
libgit2 v0.12.0

If your distro does not provide a libgit2 package you can build from go-git
repo.

	make libgit2-build
	sudo make libgit2-install

Installation
------------
	export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
    go get github.com/str1ngs/go-git

Using
------------
if /usr/local/lib is not in your ldconfig search path you need to.
	export LD_LIBRARY_PATH=/usr/local/lib
	
