# Hi dear! it is my Makefile
GOPATH := $(or $(GOPATH), $(HOME)/go)
DATABASE_URL = $(MYSQL_ADDRS)
PROJECT_DIR = $(shell pwd)

CMD_DIR = $(PROJECT_DIR)/cmd

MIGRATIONS_DIR = $(PROJECT_DIR)/migrations

WIREGEN = github.com/google/wire/cmd/wire

GOTEST = go test -tags unit -parallel $(shell nproc)

#TODO:
#Use pakage from GOPATH
#Check install package first

.PHONY: migrate
migrate: 
	migrate -source "file:$(MIGRATIONS_DIR)" -database "mysql://$(DATABASE_URL)" up $(step)

di:
	go run -mod=mod $(WIREGEN) gen $(PROJECT_DIR)/...

.PHONY: lint
lint:
	golangci-lint run -v --timeout=1m

mock:
	@rm -rf /mocks
	mockery --dir . --output mocks --keeptree --all --disable-version-string --exported

test-covering: di
	$(GOTEST) -v -race -count=1 -coverprofile=tmp/coverprofile.out.tmp ./... && go tool cover -html=tmp/coverprofile.out.tmp 

test-unit: di mock
	$(GOTEST) -v -race -count=1 ./...

test-integration: di
	go test -tags integration -v -race -p 1 -count=1 ./...

build: di 
	env GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(PROJECT_DIR)/out/ $(CMD_DIR)/server-ctl/

.PHONY: clean
	go clean
	rm $(PROJECT_DIR)/out

