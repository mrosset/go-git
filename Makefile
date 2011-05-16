test:
	@make -C pkg/git $@

pkg:
	@echo "building pkgs"
	@make -C pkg/git

cmd: pkg
	@echo "building commands"
	@make -C cmd/gogit

clean:
	@echo "cleaning"
	@make -C pkg/git $@
	@make -C cmd/gogit $@

install: pkg
	make -C pkg/git install

libgit2-init:
	@git submodule update --init libgit2

libgit2-build: libgit2-init
	cd libgit2; ./waf configure
	cd libgit2; ./waf build

libgit2-install:
	cd libgit2; ./waf install

html: clean
	@cat index.head.html > index.html
	@godoc -html ./pkg/git >> index.html
