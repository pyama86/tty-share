VERSION  := $(shell git tag | tail -n1 | sed 's/v//g')
TEST ?= $(shell go list ./... | grep -v -e vendor)
test:
	go test -v $(TEST)

.PHONY: release_major
## release_major: release nke (major)
release_major:
	git semv major --bump

.PHONY: release_minor
## release_minor: release nke (minor)
release_minor:
	git semv minor --bump

.PHONY: release_patch
## release_patch: release nke (patch)
release_patch:
	git semv patch --bump

frontend:
	cd server/frontend && npm install && npm run build
release: frontend
	goreleaser --rm-dist --skip-validate
