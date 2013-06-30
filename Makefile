# Comments starting with #: below are remake GNU Makefile comments. See
# https://github.com/rocky/remake/wiki/Rake-tasks-for-gnu-make

.PHONY: all builder interp check test check-quick test-interp

all: tortoise

tortoise: interp builder gub tortoise.go
	go build tortoise.go

#: Build the SSA Builder
builder:
	go install

#: Build the interpeter
interp: builder
	(cd interp && go install)

#: Build the debugger
gub: interp
	(cd gub && go install)

#: Same as "check"
test: check

#: Run all tests (quick and interpreter)
check: check-quick test-interp

#: Run quick tests
check-quick:
	go test -i && go test

#: Longer interpreter tests
test-interp:
	(cd interp && go test -i && go test)
