test: gogit
	make -C pkg/git test

gogit:
	make -C cmd/gogit

clean:
	make -C pkg/git clean
	make -C cmd/gogit clean
