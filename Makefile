GO=go
PACKAGES=./...

dependency:
	$(GO) mod tidy

lint:
	#TODO: fix it
	echo golangci-lint run  -v \
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

imports:
	 find -name "*.go" -not -path "./vendor/*" -exec goimports -local github.com/libmonsoon-dev/go-lib -w {} \;

pre-commit: .git/hooks/pre-commit lint build test
	echo "Done"

.git/hooks/pre-commit:
	echo "make pre-commit\n" | install /dev/stdin .git/hooks/pre-commit
