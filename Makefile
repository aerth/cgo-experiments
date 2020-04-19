all:
	go get -v -d ./...
	env GOBIN=$(PWD)/bin go install ./...
test: all
	@for i in bin/*; \
	  do echo __ $$i __ && $$i; \
	  echo ""; \
	  done;
clean:
	rm -rvf bin
