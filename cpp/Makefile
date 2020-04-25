foo: libcppfoo.a *.c *.cpp *.go
	go build -o $@
libcppfoo.a: foo.o
	ar rv $@ $^
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $<
clean:
	rm -f *.a *.o foo
