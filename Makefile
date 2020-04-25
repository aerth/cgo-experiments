all: libs/libfoo/libfoo.a cpp/libcppfoo.a
	cp libs/libfoo/libfoo.a .
	cp cpp/libcppfoo.a .
	go get -v -d ./...
	env GOBIN=$(PWD)/bin go install ./...
	@echo
	@echo Built all examples
test: all
	@echo running tests
	@echo
	@for i in bin/*; \
	  do echo __ $$i __ && $$i; \
	  echo ""; \
	  done;
libs/libfoo/libfoo.a:
	make -C libs
cpp/libcppfoo.a:
	make -C cpp libcppfoo.a
clean:
	rm -rvf bin
