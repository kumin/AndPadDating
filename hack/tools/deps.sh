#!/bin/bash

function installIfNotExist() {
	if ! command -v $1 2 /dev/null &>1; then
		echo "installing $1 ....."
		go install $2
	fi
}

function main() {
	installIfNotExist ${GOPATH}/bin/wire github.com/google/wire/cmd/wire@latest
	installIfNotExist ${GOPATH}/bin/mockery github.com/vektra/mockery/v2@latest
	installIfNotExist ${GOPATH}/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate@latest
}

main $@
