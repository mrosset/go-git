build:
	@cd libgit2; ./waf configure
	@cd libgit2; ./waf build

libgit2-init:
	@git submodule update --init libgit2

install:
	@cd libgit2; ./waf install

clean:
	@cd libgit2; ./waf clean


