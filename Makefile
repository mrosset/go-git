build:
	@mkdir libgit2_build
	@cd libgit2_build; cmake ../libgit2
	@cd libgit2_build; cmake --build .

libgit2-init:
	@git submodule update --init libgit2

install:
	@cd libgit2_build; cmake --build . --target install

clean:
	rm -rf ./libgit2_build
	rm -rf ./tmp
