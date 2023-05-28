NAME := game
OWNER := github.com/jacobsalmela

.GIT_COMMIT=$(shell git rev-parse --short HEAD)
.GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
.GIT_COMMIT_AND_BRANCH=$(.GIT_COMMIT)-$(subst /,-,$(.GIT_BRANCH))
.GIT_VERSION=$(shell git describe --tags 2>/dev/null || echo "$(.GIT_COMMIT)")

.PHONY: \
	bin

all: bin

clean:
	go clean -i ./...
	rm -rvf \
	  bin \
	  $(BUILD_DIR)

test: bin
	shellspec --format tap --no-warning-as-failure

fmt:
	go fmt ./...

tidy:
	go mod tidy

generate:
	go generate ./...

bin:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/${NAME} -ldflags "\
	-X ${OWNER}/${NAME}/cli.version=${.GIT_VERSION} \
	-X ${OWNER}/${NAME}/cli.buildDate=${.BUILDTIME} \
	-X ${OWNER}/${NAME}/cli.sha1ver=${.GIT_COMMIT_AND_BRANCH}"

run:
	go run main.go

rund:
	go run main.go -D
