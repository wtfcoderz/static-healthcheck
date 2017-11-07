FROM        golang:1.9-alpine as build
WORKDIR     /go/src
ENV         CGO_ENABLED=0
ENV         GO_PATH=/go/src
COPY        . /go/src
RUN         go build -a --installsuffix cgo --ldflags=-s -o healthcheck

FROM        scratch
COPY        --from=build /go/src/healthcheck /
CMD         ["/healthcheck"]
