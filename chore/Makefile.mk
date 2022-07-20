.PHONY      : build proto data

# use gnu make 4.0 or above
project != basename $$PWD
hostname != hostname
packages != go list -mod=vendor -f "{{.ImportPath}}" ./internal/... | sed 's/git.elewise.com\/elma365\/${project}/./'
sources := $(shell find . -name "*.go" -not -path "./vendor/*" | uniq)
docker != test -n "$$DOCKER_TAG" && echo $$DOCKER_TAG || jq '.auths["dreg.elewise.com:5005"].auth' -r < ~/.docker/config.json | base64 -d | cut -d: -f 1
deployment = ${project}
test_coverage = 75
container_name = service

RED = "\033[0;31m"
YEL = "\033[0;33m"
NC = "\033[0m"

tools       :
	# generate
	go get -u github.com/mailru/easyjson/...
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/alvaroloes/enumer
	# gobindata
	go get -u github.com/go-bindata/go-bindata
	# lint
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint


# Docker compose
up          :
	docker-compose up -d

down        :
	docker-compose down

gen-clean :
	grep -rl --exclude-dir vendor --exclude-dir protobuf '// .*generated' . | xargs -- rm

gen			:
	go generate ./internal/...


# go get -u github.com/davecheney/godoc2md
doc         :
	@for PACKAGE in ${packages}; do \
		godoc2md $$PACKAGE > $$PACKAGE/README.md; \
	done

todo        :
	! grep -rn TODO . | grep -v '\(vendor\|git\|idea\|build\)' | grep -v '#[0-9]'

# Lint and tests
lint        :
	golangci-lint run --timeout 5m0s

test        :
	CGO_ENABLED=1 go test -race -cover -count=1 -coverprofile=.coverprofile "./internal/..."

cover       :
	@sed -i '/\(pb\|easyjson\|string\)\.go/d' .coverprofile
	@go tool cover -func=.coverprofile | tail -n 1 | awk '{print "Total coverage:", $$3;}'
	@test `go tool cover -func=.coverprofile | tail -n 1 | awk '{print $$3;}' | sed 's/\..*//'` -ge ${test_coverage}

# Run ALREADY builded service on localhost with STAND dependencies
run         :
	env ELMA365_DEBUG=true ./build/${project} 2>&1 | tee run.log

run-debug   : build
	env ELMA365_DEBUG=true ./build/service 2>&1 | tee run.log | jq 'select((.logger | test("debug")) or (.level != "debug"))'

run/init    :
	env ELMA365_DEBUG=true ./build/init 2>&1 | tee run.log

build       : build/service

build/service : $(sources)
	mkdir -p ./build
	GO111MODULE=on env CGO_ENABLED=0 go build -o build/${project} -mod=vendor ./cmd/service

init       : build/init
	docker build -t dreg.elewise.com:5005/elma365/${project}/init:${branch} -f Dockerfile.init .
	docker push dreg.elewise.com:5005/elma365/${project}/init:${branch}

build/init : $(sources)
	mkdir -p ./build
	GO111MODULE=on env CGO_ENABLED=0 go build -o build/init -mod=vendor ./cmd/init

tc          :
	mkdir -p ./build
	env CGO_ENABLED=0 go build -mod=vendor -o build/load-from-tc ./cmd/load-from-tc
	build/load-from-tc ${project}

docker      : build Dockerfile
	@test "${docker}" = "develop" && echo \\n${RED}Запрещено делать локальный push в ветку ${docker}${NC}\\n && exit 1 || :
	@echo \\n${YEL}Внимание! В качестве тэга будет использован ${docker}. Вы можете изменить тэг с помощью переменной окружения DOCKER_TAG${NC}\\n
	docker build -t dreg.elewise.com:5005/elma365/${project}/${container_name}:${docker} .

push        : docker
	docker push dreg.elewise.com:5005/elma365/${project}/${container_name}:${docker}

# Go modules commands

# можно вместо go get @package делать make get p=@package
get :
	GO111MODULE=on go get ${p}
	make vendor

vendor : go.sum
	GO111MODULE=on go mod vendor

tidy :
	GO111MODULE=on go mod tidy

telepresense :
	telepresence --swap-deployment ${deployment} --run bash
