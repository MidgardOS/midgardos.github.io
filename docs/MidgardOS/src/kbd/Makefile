CC := gcc
PREFIX := /usr
BINDIR := $(PREFIX)/bin
MANDIR := $(PREFIX)/share/man
CFLAGS := -Wall -Wextra -O2
LDFLAGS :=
TARGETS := numlockbios fbtest
.PHONY: all clean install

all: $(TARGETS)
clean:
	rm -f $(TARGETS) *.o
install: all
	install -v -d -m 755 -o root -g root $(DESTDIR)$(BINDIR)
	install -v -m 755 -o root -g root $(TARGETS) $(DESTDIR)$(BINDIR)
	install -v -d -m 755 -o root -g root $(DESTDIR)$(MANDIR)/man8
	install -v -m 644 -o root -g root fbtest.8 $(DESTDIR)$(MANDIR)/man8
numlockbios: numlockbios.c
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)
fbtest: fbtest.c
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)
