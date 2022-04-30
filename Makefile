BINDIR:=$(abspath bin)

install-go-tools:
	@mkdir -p ${BINDIR}
	@cat tools.go | awk -F'"' '/_/ {print $$2}' | GOBIN=$(abspath ${BINDIR}) xargs -tI {} go install {}
