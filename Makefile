default: help

# Help Menu for makefile
.PHONY: help
help:
	@echo 'Available commands:'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST) 

.PHONY: build 
build:		## create complete build for both frontend and backend
	cd web;	npm run build
	go build -v -x

.PHONY: cgo-set 
cgo-set: 	## Set CGO to enabled
	go env -w CGO_ENABLED=1

.PHONY: clean 
clean: 		## remove distribution of both frontend and backend
	rm scheduler.exe
	rm -rf web/dist

.PHONY: dev
dev: 		## run development server for frontend ONLY
	cd web;	npm run dev

.PHONY: dist
dist: 		## Create dist/ folder for publish
	rm -rf dist
	cd web;	npm run build
	go build -v -x
	mkdir dist
	mv *.exe dist/
	mkdir dist/web
	mkdir dist/web/dist
	mv web/dist/* dist/web/dist

.PHONY: install
install: 	## install all node_modules of frontend
	cd web; npm install

.PHONY: reinstall
reinstall: 	## delete and then install all node_modules of frontend
	cd web; rm -rf node_modules && npm install

.PHONY: run
run:  		## run golang server with frontend NOT built
	clear
	go run .

.PHONY: run-build
run-build: 	## run golang server with frontend built
	cd web;	npm run build
	clear
	go run .

.PHONY: test
test: 		## perform all tests of golang
	go test ./... -count=1
