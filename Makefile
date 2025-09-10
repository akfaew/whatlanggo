.ifndef _GOLANG_MK_
_GOLANG_MK_=1

.MAIN: build

GOFLAGS?= -mod=mod
GOWORK?= off
.export GOFLAGS GOWORK

BINDIR?= bin

REPO_ROOT!= git rev-parse --show-toplevel 2>/dev/null || echo ${.CURDIR}
.export REPO_ROOT

TEST_ARGS?= -failfast

DATASTORE_PORT_RUN?=8001
DATASTORE_PORT_TEST?=8002

.if !target(build)
.PHONY: build
build: test
	go build .
.endif

.if !target(run)
.PHONY: run
run: build
	./${NAME}
.endif


.if !target(fmt)
.PHONY: fmt
fmt:
	# golines -m 120 -t 8 --shorten-comments --ignore-generated -w .
	gofumpt -l -w .
	git ls-files -- '*.go' | xargs -n 10 -P `nproc` goimports -local github.com/akfaew -w
	templ fmt .
.endif


.if !target(lint)
.PHONY: lint
lint:
	GOWORK=off golangci-lint run --fix --allow-parallel-runners ${LINT_ARGS}
.endif

.if !target(test)
.PHONY: test

.if defined(NO_FMT)
test: lint
.else
test: fmt lint
.endif
	DATASTORE_EMULATOR_HOST=127.0.0.1:${DATASTORE_PORT_TEST} \
	DATASTORE_EMULATOR_HOST_PATH=127.0.0.1:${DATASTORE_PORT_TEST}/datastore \
	DATASTORE_HOST=http://127.0.0.1:${DATASTORE_PORT_TEST} \
	go test ${TEST_ARGS} ./...
.endif


.if !target(test-docker)
.PHONY: test-docker
test-docker: fmt lint
	${BINDIR}/with-datastore.sh env \
		DATASTORE_EMULATOR_HOST=127.0.0.1:${DATASTORE_PORT_TEST} \
		DATASTORE_EMULATOR_HOST_PATH=127.0.0.1:${DATASTORE_PORT_TEST}/datastore \
		DATASTORE_HOST=http://127.0.0.1:${DATASTORE_PORT_TEST} \
		go test ${TEST_ARGS} ./...
.endif


.if !target(vendor)
.PHONY: vendor
vendor:
	rm -rf vendor/
	GOWORK=off go mod vendor
.endif


.if !target(update)
.PHONY: update
update:
	GOWORK=off go get -u -t ./...
	GOWORK=off go mod tidy
	GOWORK=off go mod verify
	${MAKE} vendor
.endif


.if !target(test-cover)
.PHONY: test-cover
test-cover:
	go test ${TEST_ARGS} -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
.endif

.if !target(test-regen)
.PHONY: test-regen
test-regen:
	rm -rf testdata/output
	mkdir -p testdata/output
	go test ${TEST_ARGS} -regen .
.endif

.if !target(clean)
.PHONY: clean
clean:
	rm -rf coverage.out coverage.html callvis.dot callvis.png vendor/ ${NAME}
.endif

.endif	# _GOLANG_MK_
