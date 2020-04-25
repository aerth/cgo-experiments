LDFLAGS=-Llibs/libfoo
all: libs/libfoo/libfoo.a
	go get -v -d ./...
	env GOBIN=$(PWD)/bin go install ./...
test: all
	@for i in bin/*; \
	  do echo __ $$i __ && $$i; \
	  echo ""; \
	  done;
libs/libfoo/libfoo.a:
	make -C libs
clean:
	rm -rvf bin
