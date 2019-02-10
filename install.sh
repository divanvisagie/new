#!/bin/bash
cp new /usr/local/bin/
mkdir -p /usr/local/man/man1
install -g 0 -o 0 -m 0644 new.1 /usr/local/man/man1
gzip /usr/local/man/man1/new.1
