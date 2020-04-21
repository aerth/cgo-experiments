CFLAGS ?= -I../libfoo/
LDFLAGS ?= -L../libfoo -lfoo
foo_test: *.c
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)
clean:
	rm -vf foo_test *.a *.o
