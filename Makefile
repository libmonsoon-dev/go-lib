GO=go
PACKAGES=./...

dependency:
	$(GO) mod tidy; $(GO) mod vendor

lint:
	golangci-lint run  -v \
 		--disable gosimple \
 		--disable staticcheck \
 		--disable unused \
 		$(PACKAGES)

build:
	$(GO) build -v $(PACKAGES)

test:
	$(GO) test -v -race $(PACKAGES)

bench:
	$(GO) test -bench=. $(PACKAGES)

pre-commit: .git/hooks/pre-commit dependency lint build test
	echo "Done"

.git/hooks/pre-commit:
	echo "make pre-commit\n" | install /dev/stdin .git/hooks/pre-commit
