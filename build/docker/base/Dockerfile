################################################################################
# Docker file for building base image for 99 Group Core Asset Manager
# in Production & Develop
#
# VERSION:
# 0.0.0 - initialization
# 0.0.1 - fix folder initialization
################################################################################
FROM golang:1.19-alpine

# Enable Go Modules.
ENV GO111MODULE=on\
    CGO_ENABLED=1 \
    CGO_CFLAGS_ALLOW=-Xpreprocessor \
    GOPROXY=https://proxy.golang.org \
    GOOS=linux \
    GOARCH=amd64


RUN apk --no-cache add gcc g++ make ca-certificates curl git openssh
RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest && \
    go install github.com/rubenv/sql-migrate/...@latest && \
    mkdir -p /go/src/core-banking

# RUN go env -w GOPRIVATE=github.com/urbanindo/*

# ARG NETRC
# RUN echo ${NETRC} > ~/.netrc && chmod 600 ~/.netrc