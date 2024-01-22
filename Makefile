default: help

# Help Menu for makefile
.PHONY: help
help:
	@echo 'Available commands: dev, clean, reinstall, build, run, test'

# start the development server
.PHONY: dev
dev:
	cd web;	npm run start

# clean all build file
.PHONY: clean
clean:
	rm AnimeImageAnalyser-v2.exe
	rm -rf web/dist
	
# build both react & golang executable
.PHONY: build
build:
	cd web;	npm run build
	go build -v -x

# Run the golang program
.PHONY: run
run:
	cd web;	npm run build
	clear
	go run .

# Run all tests for golang
.PHONY: test
test:
	go test ./... -count=1

## Reinstall all dependencies
.PHONY: reinstall
reinstall:
	cd web; rm -rf node_modules && npm install

# Install all dependencies
.PHONY: install
install:
	cd web; npm install