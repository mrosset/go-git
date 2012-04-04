go-git
=================
Go bindings to [libgit2](http://libgit2.github.com/)

Status
=================
Currently in active development.

Documentation
-----------------
[gopkgdoc go-git](http://gopkgdoc.appspot.com/pkg/github.com/str1ngs/go-git)


Requirements
-----------------
libgit2 0.12.0
pkg-config
libssl-dev

If your distro does not provide a libgit2 package you can build from go-git
repo.

	git clone --recursive git://github.com/str1ngs/go-git.git
	cd go-git/
	make
	sudo make install

Installation
------------
	export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
    go get github.com/str1ngs/go-git

Using
------------
if /usr/local/lib is not in your ldconfig search path you need to.
	export LD_LIBRARY_PATH=/usr/local/lib
	
