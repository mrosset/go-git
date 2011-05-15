test: modules clean
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

install:
	make -C pkg/git install

modules: libgit2
	@git sub init libgit2
	@git sub update libgit2

html: pkg/git
	godoc -html ./pkg/git > index.html
	godoc -html ./pkg/git > README.md
