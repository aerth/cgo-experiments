test_library_program: libfoo/libfoo.a *.go
	go build -o $@
libfoo/libfoo.a: libfoo/*.c
	make -C libfoo
clean:
	rm -vf test_library_program
	make -C libfoo clean
	make -C foo_test clean
