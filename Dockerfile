FROM docker.io/golang:1.21.1-bookworm AS compiler

WORKDIR /src

COPY . /src/build

RUN cd /src/build && go mod download && make build 

FROM alpine:3.14.2

COPY --from=compiler src/build/out/* /bin/

CMD ["./bin/server-ctl"]
